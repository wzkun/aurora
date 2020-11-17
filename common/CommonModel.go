package common

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"code.aliyun.com/new_backend/scodi_nqc/decode"
	"code.aliyun.com/new_backend/scodi_nqc/model"
	"code.aliyun.com/new_backend/scodi_nqc/redis"
	"code.aliyun.com/new_backend/scodi_nqc/utils"

	"github.com/jinzhu/gorm"
	"github.com/wzkun/aurora/caching"
)

// CommonModel 数据模型操作对象
type CommonModel struct {
	Imp      *model.ModelImp
	RedisImp redis.Redis
}

// NewCommonModel  新建数据模型操作对象
func NewCommonModel() *CommonModel {
	imp := model.NewModelImp()
	if imp == nil {
		return nil
	}
	model2 := &CommonModel{Imp: imp, RedisImp: redis.GetInstance()}
	return model2
}

// Insert 插入单条数据
func (model2 *CommonModel) Insert(value interface{}) (err error) {
	return model2.GormDB().Create(value).Error
}

// Inserts 批量插入操作
func (model2 *CommonModel) Inserts(values ...interface{}) (err error) {
	return model2.InsertsWithTag("insert into ", values...)
}

// InsertsIgnore 批量插入操作(忽略重复)
func (model2 *CommonModel) InsertsIgnore(values ...interface{}) (err error) {
	return model2.InsertsWithTag("insert ignore into ", values...)
}

// ReplaceMulti 批量插入操作(覆盖重复)
func (model2 *CommonModel) ReplaceMulti(values ...interface{}) (err error) {
	return model2.InsertsWithTag("replace into ", values...)
}

// ValueIsTime
func ValueIsTime(field reflect.Value) bool {
	tempType := field.Type()
	return tempType.AssignableTo(reflect.TypeOf((*time.Time)(nil)))
}

// ValueIsStruct
func ValueIsStruct(field reflect.Value) bool {
	tempType := field.Type()
	for tempType.Kind() == reflect.Ptr {
		tempType = tempType.Elem()
	}

	return tempType.Kind() == reflect.Struct
}

// InsertsWithTag 批量插入操作
func (model2 *CommonModel) InsertsWithTag(tag string, values ...interface{}) (err error) {
	fmt.Println("==tag===", tag)
	if len(values) <= 0 {
		return nil
	}
	tx := model2.GormDB().Begin()

	pre := 200
	num := len(values) / pre
	if len(values)%pre != 0 {
		num++
	}
	sqlString := bytes.Buffer{}
	sqlString.WriteString(tag + " ")
	sqlString.WriteString(values[0].(caching.Item).TableName())
	sqlString.WriteString(" (")
	rValue := reflect.Indirect(reflect.ValueOf(values[0]))
	for i := 0; i < rValue.NumField(); i++ {

		if ValueIsStruct(rValue.Field(i)) {
			if !ValueIsTime(rValue.Field(i)) {
				continue
			}
		}
		jsonName := rValue.Type().Field(i).Tag.Get("json")
		if len(jsonName) == 0 {
			jsonName = utils.Marshal(rValue.Type().Field(i).Name)
		}
		if (rValue.Field(i).Kind() == reflect.Int ||
			rValue.Field(i).Kind() == reflect.Int8 || rValue.Field(i).Kind() == reflect.Int32 ||
			rValue.Field(i).Kind() == reflect.Int64 || rValue.Field(i).Kind() == reflect.Int16) &&
			(rValue.Field(i).Interface() == nil || rValue.Field(i).Interface() == 0) {
			continue
		}
		if jsonName == "-" {
			continue
		}
		sqlString.WriteString("`")
		sqlString.WriteString(jsonName)
		sqlString.WriteString("`")
		if i != rValue.NumField()-1 {
			sqlString.WriteString(",")
		}
	}
	if strings.HasSuffix(sqlString.String(), ",") {
		sqlString.Truncate(sqlString.Len() - 1)
	}
	sqlString.WriteString(") values(")
	fmt.Println("sqlString.String()", sqlString.String())
	for i := 0; i < num; i++ {

		var items []interface{}
		if (i+1)*pre > len(values) {
			items = values[i*pre:]
		} else {
			items = values[i*pre : (i+1)*pre]
		}
		sqlString2 := bytes.Buffer{}
		sqlString2.WriteString(sqlString.String())
		var vList []interface{}
		for index, v := range items {
			rValue2 := reflect.Indirect(reflect.ValueOf(v))
			for i := 0; i < rValue2.NumField(); i++ {
				if ValueIsStruct(rValue2.Field(i)) {
					if !ValueIsTime(rValue.Field(i)) {
						continue
					}
				jsonName := rValue2.Type().Field(i).Tag.Get("json")
				if len(jsonName) == 0 {
					jsonName = utils.Marshal(rValue2.Type().Field(i).Name)
				}
				if (rValue.Field(i).Kind() == reflect.Int ||
					rValue.Field(i).Kind() == reflect.Int8 || rValue.Field(i).Kind() == reflect.Int32 ||
					rValue.Field(i).Kind() == reflect.Int64 || rValue.Field(i).Kind() == reflect.Int16) && rValue.Field(i).Interface() == 0 {
					continue
				}
				if jsonName == "-" {
					continue
				}
				sqlString2.WriteString("?")
				if i != rValue.NumField()-1 {
					sqlString2.WriteString(",")
				}
				vList = append(vList, rValue2.Field(i).Interface())
			}
			if strings.HasSuffix(sqlString2.String(), ",") {
				sqlString2.Truncate(sqlString2.Len() - 1)
			}
			if index != len(items)-1 {
				sqlString2.WriteString("),(")
			} else {
				sqlString2.WriteString(")")
			}
		}

		//fmt.Println("vList",vList[:30])
		DB := tx.Exec(sqlString2.String(), vList...)
		if DB.Error != nil {
			fmt.Println("db", DB.Error)
			tx.Rollback()
			return DB.Error
		} else {
			fmt.Println("db成功", i, "/", num)
		}
	}

	return tx.Commit().Error
}

// Save 新建数据模型
func (model2 *CommonModel) Save(tableName string, value interface{}) (err error) {
	return model2.Imp.Table(tableName).Save(value).Error
}

// Updates 批量更新操作，默认字段，如bool的false，int的0，string的空字符串都会被忽略
func (model2 *CommonModel) Updates(tableName string, value interface{}) (err error) {
	return model2.Imp.Table(tableName).UpdateMulti(value)
}

// Get function
func (model2 *CommonModel) Get(id string, value interface{}) error {
	return model2.GormDB().Where("id = ?", id).Find(value).Error
}

// Get function
func (model2 *CommonModel) Gets(ids []string, value interface{}) error {
	return model2.GormDB().Where("id in (?)", ids).Find(value).Error
}

// Delete 新建数据模型
func (model2 *CommonModel) Delete(value interface{}) (err error) {
	return model2.Imp.Delete(value).Error
}

// DeleteMulti 批量删除操作
func (model2 *CommonModel) DeleteMulti(value interface{}) (err error) {
	return model2.Imp.DeleteMulti(value)
}

// SaveToRedis 单条Redis保存
func (model2 *CommonModel) SaveToRedis(tableName string, key, value interface{}) (err error) {
	_, err = model2.RedisImp.Set(tableName, key, value)
	return
}

// QueryByRedisByID 单条Redis请求
func (model2 *CommonModel) QueryByRedisByID(tableName, id string, value interface{}) (err error) {
	return model2.RedisImp.GetResult(tableName, id, value)
}

// DeleteRedis 单条Redis删除
func (model2 *CommonModel) DeleteRedis(tableName string, key interface{}) (err error) {
	return model2.RedisImp.Delete(tableName, key)
}

// DeleteRedis 单条Redis删除
func (model2 *CommonModel) GetRedisResult(tableName string, key interface{}, value interface{}) (err error) {
	result, err := model2.RedisImp.Exec("HGET", tableName, key)
	if err != nil {
		return err
	}

	if result == nil || len(result.([]byte)) == 0 {
		return nil
	}

	rValue := reflect.ValueOf(value)
	var value2 interface{}
	if rValue.IsValid() {
		value2 = reflect.New(reflect.Indirect(rValue).Type()).Interface()
	} else {
		return errors.New("value is invalid")
	}

	err = utils.ConvertToInterface(result.([]byte), value2)
	if err != nil {
		return err
	}

	value = value2

	return nil
}

// Query Query
func (model2 *CommonModel) Query(value interface{}) error {
	return model2.GormDB().First(value).Error
}

// QueryAll QueryAll
func (model2 *CommonModel) QueryAll(value interface{}) error {
	return model2.GormDB().Find(value).Error
}

// Limit Limit
func (model2 *CommonModel) Limit(limit ...int) *CommonModel {
	model2.Imp = model2.Imp.Limit(limit...)
	return model2
}

// Offset Offset
func (model2 *CommonModel) Offset(offset int) *CommonModel {
	model2.Imp = model2.Imp.Offset(offset)
	return model2
}

// GormDB 获取gorm操作对象
func (model2 *CommonModel) GormDB() *gorm.DB {
	return model2.Imp.DB
}

// DB 获取数据库操作对象
func (model2 *CommonModel) DB() *sql.DB {
	return model2.GormDB().DB()
}

// New 重置
func (model2 *CommonModel) New() *CommonModel {
	model2.Imp.New()
	return model2
}

// Begin 开启事务
func (model2 *CommonModel) Begin() *CommonModel {
	model2.Imp.DB = model2.Imp.Begin()
	model2.New()
	return model2
}

// Commit 提交事务
func (model2 *CommonModel) Commit() error {
	model2.Imp.DB = model2.Imp.Commit()
	return model2.Imp.DB.Error
}

// Rollback 事务回滚
func (model2 *CommonModel) Rollback() error {
	return model2.Imp.Rollback().Error
}

// MarshalToJson 序列化为JSON
func (model2 *CommonModel) MarshalToJson(value interface{}) ([]byte, error) {
	var result = make(map[string]interface{})
	objValue := reflect.Indirect(reflect.ValueOf(value))
	for i := 0; i < objValue.NumField(); i++ {
		es := objValue.Type().Field(i).Tag.Get("es")
		jsonName := objValue.Type().Field(i).Tag.Get("nJson")
		if len(jsonName) == 0 {
			//jsonName = objValue.Type().Field(i).Name
			continue
		}
		if es == "true" {
			var v interface{}
			decode.JSON.Unmarshal([]byte(objValue.Field(i).String()), &v)
			result[jsonName] = v
		} else {
			result[jsonName] = objValue.Field(i).Interface()
			//reflect.ValueOf(result[objType.Field(i).Name]).Set(objValue.Field(i))
		}
	}
	return decode.JSON.Marshal(result)
}

// WhereProjectID 查询条件设置
func (model2 *CommonModel) WhereProjectID(Predicate model.Predicate, values ...string) *CommonModel {
	values2 := make([]interface{}, len(values))
	for index, v := range values {
		values2[index] = v
	}
	model2.Imp = model2.Imp.Where("`project_id`", Predicate, values2...)
	return model2
}

// WhereType 查询条件设置
func (model2 *CommonModel) WhereType(Predicate model.Predicate, values ...string) *CommonModel {
	values2 := make([]interface{}, len(values))
	for index, v := range values {
		values2[index] = v
	}
	model2.Imp = model2.Imp.Where("`type`", Predicate, values2...)
	return model2
}

// WhereLandId 查询条件设置
func (model2 *CommonModel) WhereLandID(Predicate model.Predicate, values ...string) *CommonModel {
	values2 := make([]interface{}, len(values))
	for index, v := range values {
		values2[index] = v
	}
	model2.Imp = model2.Imp.Where("`land_id`", Predicate, values2...)
	return model2
}

// WhereModelType 查询条件设置
func (model2 *CommonModel) WhereModelType(Predicate model.Predicate, values ...string) *CommonModel {
	values2 := make([]interface{}, len(values))
	for index, v := range values {
		values2[index] = v
	}
	model2.Imp = model2.Imp.Where("`model_type`", Predicate, values2...)
	return model2
}

// WhereVersion 查询条件设置
func (model2 *CommonModel) WhereVersion(Predicate model.Predicate, values ...string) *CommonModel {
	values2 := make([]interface{}, len(values))
	for index, v := range values {
		values2[index] = v
	}
	model2.Imp = model2.Imp.Where("`version`", Predicate, values2...)
	return model2
}

// WhereID 查询条件设置
func (model2 *CommonModel) WhereID(Predicate model.Predicate, values ...string) *CommonModel {
	values2 := make([]interface{}, len(values))
	for index, v := range values {
		values2[index] = v
	}
	model2.Imp = model2.Imp.Where("`id`", Predicate, values2...)
	return model2
}

// WhereName 查询条件设置
func (model2 *CommonModel) WhereName(Predicate model.Predicate, values ...string) *CommonModel {
	values2 := make([]interface{}, len(values))
	for index, v := range values {
		values2[index] = v
	}
	model2.Imp = model2.Imp.Where("`name`", Predicate, values2...)
	return model2
}

// WhereUserID 查询条件设置
func (model2 *CommonModel) WhereUserID(Predicate model.Predicate, values ...string) *CommonModel {
	values2 := make([]interface{}, len(values))
	for index, v := range values {
		values2[index] = v
	}
	model2.Imp = model2.Imp.Where("`user_id`", Predicate, values2...)
	return model2
}

// WhereKind 查询条件设置
func (model2 *CommonModel) WhereKind(Predicate model.Predicate, values ...string) *CommonModel {
	values2 := make([]interface{}, len(values))
	for index, v := range values {
		values2[index] = v
	}
	model2.Imp = model2.Imp.Where("`kind`", Predicate, values2...)
	return model2
}

// func (imp *ModelImp) Where(name string, predicate Predicate, values ...interface{}) *ModelImp {
// 	//imp.DB = imp.DB.Where(fmt.Sprintf("%v %v ?", name, predicate), value)
// 	return imp
// }

// Where 通用查询条件设置
func (model2 *CommonModel) Where(name string, predicate model.Predicate, values ...interface{}) *CommonModel {
	values2 := make([]interface{}, len(values))
	for index, v := range values {
		values2[index] = v
	}

	model2.Imp = model2.Imp.Where(name, predicate, values2...)
	return model2
}

// WhereEqual 通用查询条件设置
func (model2 *CommonModel) WhereEqual(name string, values ...interface{}) *CommonModel {
	values2 := make([]interface{}, len(values))
	for index, v := range values {
		values2[index] = v
	}

	model2.Imp = model2.Imp.Where(name, model.Equal, values2...)
	return model2
}

// WherePublishRecordID 查询条件设置
func (model2 *CommonModel) WherePublishRecordID(Predicate model.Predicate, values ...string) *CommonModel {
	values2 := make([]interface{}, len(values))
	for index, v := range values {
		values2[index] = v
	}
	model2.Imp = model2.Imp.Where("`publish_record_id`", Predicate, values2...)
	return model2
}

// WhereTypeID 查询条件设置
func (model2 *CommonModel) WhereTypeID(Predicate model.Predicate, values ...string) *CommonModel {
	values2 := make([]interface{}, len(values))
	for index, v := range values {
		values2[index] = v
	}
	model2.Imp = model2.Imp.Where("`type_id`", Predicate, values2...)
	return model2
}

// WhereStandardID 查询条件设置
func (model2 *CommonModel) WhereStandardID(Predicate model.Predicate, values ...string) *CommonModel {
	values2 := make([]interface{}, len(values))
	for index, v := range values {
		values2[index] = v
	}
	model2.Imp = model2.Imp.Where("`standard_id`", Predicate, values2...)
	return model2
}

// WhereParentID 查询条件设置
func (model2 *CommonModel) WhereParentID(Predicate model.Predicate, values ...string) *CommonModel {
	values2 := make([]interface{}, len(values))
	for index, v := range values {
		values2[index] = v
	}
	model2.Imp = model2.Imp.Where("`parent_id`", Predicate, values2...)
	return model2
}

// WhereIsFinished 查询条件设置
func (model2 *CommonModel) WhereIsFinished(Predicate model.Predicate, values ...bool) *CommonModel {
	values2 := make([]interface{}, len(values))
	for index, v := range values {
		values2[index] = v
	}
	model2.Imp = model2.Imp.Where("`is_finished`", Predicate, values2...)
	return model2
}

// WhereKind 查询条件设置
func (model2 *CommonModel) WhereTitle(Predicate model.Predicate, value string) *CommonModel {
	model2.Imp = model2.Imp.Where("`title`", Predicate, value)
	return model2
}

// WhereLeftLike 查询条件设置
func (model2 *CommonModel) WhereLeftLike(name, value string) *CommonModel {
	model2.Imp.DB = model2.GormDB().Where(fmt.Sprintf("%v like (?)", name), "%"+value)
	return model2
}

// WhereRightLike 查询条件设置
func (model2 *CommonModel) WhereRightLike(name, value string) *CommonModel {
	model2.Imp.DB = model2.GormDB().Where(fmt.Sprintf("%v like (?)", name), value+"%")
	return model2
}

// WhereLike 查询条件设置
func (model2 *CommonModel) WhereLike(name, value string) *CommonModel {
	model2.Imp.DB = model2.GormDB().Where(fmt.Sprintf("%v like (?)", name), "%"+value+"%")
	return model2
}

// WhereProcInstID 查询条件设置
func (model2 *CommonModel) WhereProcInstID(Predicate model.Predicate, values ...string) *CommonModel {
	values2 := make([]interface{}, len(values))
	for index, v := range values {
		values2[index] = v
	}
	model2.Imp = model2.Imp.Where("`proc_inst_id`", Predicate, values2...)
	return model2
}

// WhereProcdefID 查询条件设置
func (model2 *CommonModel) WhereProcdefID(Predicate model.Predicate, values ...string) *CommonModel {
	values2 := make([]interface{}, len(values))
	for index, v := range values {
		values2[index] = v
	}
	model2.Imp = model2.Imp.Where("`procdef_id`", Predicate, values2...)
	return model2
}

// WhereProcdefName 查询条件设置
func (model2 *CommonModel) WhereProcdefName(Predicate model.Predicate, values ...string) *CommonModel {
	values2 := make([]interface{}, len(values))
	for index, v := range values {
		values2[index] = v
	}
	model2.Imp = model2.Imp.Where("`procdef_name`", Predicate, values2...)
	return model2
}

// WhereStartUserName 查询条件设置
func (model2 *CommonModel) WhereStartUserName(Predicate model.Predicate, values ...string) *CommonModel {
	values2 := make([]interface{}, len(values))
	for index, v := range values {
		values2[index] = v
	}
	model2.Imp = model2.Imp.Where("`start_user_name`", Predicate, values2...)
	return model2
}

// WhereStatus 查询条件设置
func (model2 *CommonModel) WhereStatus(Predicate model.Predicate, values ...string) *CommonModel {
	values2 := make([]interface{}, len(values))
	for index, v := range values {
		values2[index] = v
	}
	model2.Imp = model2.Imp.Where("`status`", Predicate, values2...)
	return model2
}

// WhereDepartmentID 查询条件设置
func (model2 *CommonModel) WhereDepartmentID(Predicate model.Predicate, values ...string) *CommonModel {
	values2 := make([]interface{}, len(values))
	for index, v := range values {
		values2[index] = v
	}
	model2.Imp = model2.Imp.Where("`department_id`", Predicate, values2...)
	return model2
}

// WhereCompanyID 查询条件设置
func (model2 *CommonModel) WhereCompanyID(Predicate model.Predicate, values ...string) *CommonModel {
	values2 := make([]interface{}, len(values))
	for index, v := range values {
		values2[index] = v
	}
	model2.Imp = model2.Imp.Where("`company_id`", Predicate, values2...)
	return model2
}

// WhereCandidate 查询条件设置
func (model2 *CommonModel) WhereCandidate(Predicate model.Predicate, values ...string) *CommonModel {
	values2 := make([]interface{}, len(values))
	for index, v := range values {
		values2[index] = v
	}
	model2.Imp = model2.Imp.Where("`candidate`", Predicate, values2...)
	return model2
}

// WhereSectionID 查询条件设置
func (model2 *CommonModel) WhereSectionID(Predicate model.Predicate, values ...string) *CommonModel {
	values2 := make([]interface{}, len(values))
	for index, v := range values {
		values2[index] = v
	}
	model2.Imp = model2.Imp.Where("`section_id`", Predicate, values2...)
	return model2
}

// WhereTaskID 查询条件设置
func (model2 *CommonModel) WhereTaskID(Predicate model.Predicate, values ...string) *CommonModel {
	values2 := make([]interface{}, len(values))
	for index, v := range values {
		values2[index] = v
	}
	model2.Imp = model2.Imp.Where("`task_id`", Predicate, values2...)
	return model2
}

// WhereExecutionID 查询条件设置
func (model2 *CommonModel) WhereExecutionID(Predicate model.Predicate, values ...string) *CommonModel {
	values2 := make([]interface{}, len(values))
	for index, v := range values {
		values2[index] = v
	}
	model2.Imp = model2.Imp.Where("`execution_id`", Predicate, values2...)
	return model2
}

// WhereFormValueID 查询条件设置
func (model2 *CommonModel) WhereFormValueID(Predicate model.Predicate, values ...string) *CommonModel {
	values2 := make([]interface{}, len(values))
	for index, v := range values {
		values2[index] = v
	}
	model2.Imp = model2.Imp.Where("`form_value_id`", Predicate, values2...)
	return model2
}

// WhereStartUserID 查询条件设置
func (model2 *CommonModel) WhereStartUserID(Predicate model.Predicate, values ...string) *CommonModel {
	values2 := make([]interface{}, len(values))
	for index, v := range values {
		values2[index] = v
	}
	model2.Imp = model2.Imp.Where("`start_user_id`", Predicate, values2...)
	return model2
}

// WhereNodeID 查询条件设置
func (model2 *CommonModel) WhereNodeID(Predicate model.Predicate, values ...string) *CommonModel {
	values2 := make([]interface{}, len(values))
	for index, v := range values {
		values2[index] = v
	}
	model2.Imp = model2.Imp.Where("`node_id`", Predicate, values2...)
	return model2
}

// WhereDay 查询条件设置
func (model2 *CommonModel) WhereDay(Predicate model.Predicate, values ...string) *CommonModel {
	values2 := make([]interface{}, len(values))
	for index, v := range values {
		values2[index] = v
	}
	model2.Imp = model2.Imp.Where("`day`", Predicate, values2...)
	return model2
}

// WhereDistrictCode 查询条件设置
func (model2 *CommonModel) WhereDistrictCode(Predicate model.Predicate, values ...string) *CommonModel {
	values2 := make([]interface{}, len(values))
	for index, v := range values {
		values2[index] = v
	}
	model2.Imp = model2.Imp.Where("`district_code`", Predicate, values2...)
	return model2
}

// WhereRegionID 查询条件设置
func (model2 *CommonModel) WhereRegionID(Predicate model.Predicate, values ...string) *CommonModel {
	values2 := make([]interface{}, len(values))
	for index, v := range values {
		values2[index] = v
	}
	model2.Imp = model2.Imp.Where("`region_id`", Predicate, values2...)
	return model2
}

// WhereEbsID 查询条件设置
func (model2 *CommonModel) WhereEbsID(Predicate model.Predicate, values ...string) *CommonModel {
	values2 := make([]interface{}, len(values))
	for index, v := range values {
		values2[index] = v
	}
	model2.Imp = model2.Imp.Where("`ebs_id`", Predicate, values2...)
	return model2
}

// WherePEbsID 查询条件设置
func (model2 *CommonModel) WherePEbsID(Predicate model.Predicate, values ...string) *CommonModel {
	values2 := make([]interface{}, len(values))
	for index, v := range values {
		values2[index] = v
	}
	model2.Imp = model2.Imp.Where("`p_ebs_id`", Predicate, values2...)
	return model2
}

// WhereEbsUnitTreeID 查询条件设置
func (model2 *CommonModel) WhereEbsUnitTreeID(Predicate model.Predicate, values ...string) *CommonModel {
	values2 := make([]interface{}, len(values))
	for index, v := range values {
		values2[index] = v
	}
	model2.Imp = model2.Imp.Where("`ebs_unit_tree_id`", Predicate, values2...)
	return model2
}

// OrderBy 排序条件设置
func (model2 *CommonModel) OrderBy(name string, desc bool) *CommonModel {
	model2.Imp = model2.Imp.Order(name, desc)
	return model2
}

// OrderByProcInstID 排序条件设置
func (model2 *CommonModel) OrderByProcInstID(desc bool) *CommonModel {
	model2.Imp = model2.Imp.Order("`proc_inst_id`", desc)
	return model2
}

// OrderByProjectID 排序条件设置
func (model2 *CommonModel) OrderByProjectID(desc bool) *CommonModel {
	model2.Imp = model2.Imp.Order("`project_id`", desc)
	return model2
}

// OrderByID 排序条件设置
func (model2 *CommonModel) OrderByID(desc bool) *CommonModel {
	model2.Imp = model2.Imp.Order("`id`", desc)
	return model2
}

// OrderByName 排序条件设置
func (model2 *CommonModel) OrderByName(desc bool) *CommonModel {
	model2.Imp = model2.Imp.Order("`name`", desc)
	return model2
}

// OrderByCreateTime 排序条件设置
func (model2 *CommonModel) OrderByCreateTime(desc bool) *CommonModel {
	model2.Imp = model2.Imp.Order("`create_time`", desc)
	return model2
}

// OrderByCreateTime 排序条件设置
func (model2 *CommonModel) OrderByModifyTime(desc bool) *CommonModel {
	model2.Imp = model2.Imp.Order("`modify_time`", desc)
	return model2
}

// OrderByNodeID 排序条件设置
func (model2 *CommonModel) OrderByNodeID(desc bool) *CommonModel {
	model2.Imp = model2.Imp.Order("`node_id`", desc)
	return model2
}

// OrderByClaimTime 排序条件设置
func (model2 *CommonModel) OrderByClaimTime(desc bool) *CommonModel {
	model2.Imp = model2.Imp.Order("`claim_time`", desc)
	return model2
}

// Count 数据条数
func (model2 *CommonModel) Count(model interface{}) int {
	count, _ := model2.Imp.Count(model)
	return count
}
