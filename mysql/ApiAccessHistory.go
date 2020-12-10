package mysql

import (
	"bytes"
	"time"

	"github.com/wzkun/aurora/common"
	"github.com/wzkun/aurora/utils"
)

// ApiAccessHistory 数据模型
type ApiAccessHistory struct {
	Id          string ` optional:"true" json:"id" nJson:"id" gorm:"comment:'数据编号'"`
	Account     string ` optional:"true" json:"account" nJson:"account" gorm:"comment:'用户账号'"`
	ServerName  string ` optional:"true" json:"server_name" nJson:"serverName" gorm:"comment:'微服务名称'"`
	ModuleName  string ` json:"module_name" nJson:"moduleName" optional:"true" gorm:"comment:'模块名称'"`
	ServiceName string ` optional:"true" json:"service_name" nJson:"serviceName" gorm:"comment:'服务名称'"`
	ApiName     string ` optional:"true" json:"api_name" nJson:"apiName" gorm:"comment:'接口名称'"`
	CreateTime  string ` optional:"true" json:"create_time" nJson:"createTime" gorm:"comment:'操作时间'"`
	Request     string ` nJson:"request" optional:"true" json:"request" gorm:"comment:'接口参数';type:longtext"`
	Response    string ` optional:"true" json:"response" nJson:"response" gorm:"comment:'返回数据';type:longtext"`
	Error       string ` optional:"true" json:"error" nJson:"error" gorm:"comment:'错误信息';type:longtext"`
}

// TableName 数据表名称
func (*ApiAccessHistory) TableName() string {
	return "common_server_apiaccesshistory"
}

// Insert Insert
func (o *ApiAccessHistory) Insert() (err error) {
	return common.NewCommonModel().GormDB().Create(o).Error
}

// Save Save
func (o *ApiAccessHistory) Save() (err error) {
	return common.NewCommonModel().Imp.Table(o.TableName()).Save(o).Error
}

// Delete Delete
func (o *ApiAccessHistory) Delete() (err error) {
	return common.NewCommonModel().GormDB().Delete(o).Error
}

// MarshalToJson 序列化为JSON
func (o *ApiAccessHistory) MarshalToJson() ([]byte, error) {
	return common.NewCommonModel().MarshalToJson(o)
}

// UID 内存id
func (o *ApiAccessHistory) UID() string {
	b := bytes.Buffer{}
	b.WriteString(o.TableName())
	b.WriteString(o.Id)
	return b.String()
}

// Idx
func (o *ApiAccessHistory) Idx() string {
	return o.Id
}

// Size
func (o *ApiAccessHistory) Size() int {
	return 1
}

// ElasticIndex InterSection弹性搜索表名
func (o *ApiAccessHistory) ElasticIndex() string {
	return "common_buckets_log_collections_" + o.TableName()
}

// DBModel
func (o *ApiAccessHistory) DBModel() ElasticModel {
	return nil
}

// QueryAllItem
func (o *ApiAccessHistory) QueryAllItem(imp *ModelImp) ([]ModelItem, error) {
	return nil, nil
}

// QueryItems
func (o *ApiAccessHistory) QueryItems(sqlCommand string) ([]ModelItem, error) {
	return nil, nil
}

// InsertAll
func (o *ApiAccessHistory) InsertAll(*ModelImp, []ModelItem) {
}

//NewModelApiAccessHistory function.
func NewModelApiAccessHistory(account, serverName, moduleName, serviceName, apiName, request, response, errorstring string) *ApiAccessHistory {
	rd := &ApiAccessHistory{}

	rd.Id = utils.NewUUIdV4()
	rd.Account = account
	rd.ServerName = serverName
	rd.ModuleName = moduleName
	rd.ServiceName = serviceName
	rd.ApiName = apiName
	rd.CreateTime = time.Now().String()
	rd.Request = request
	rd.Response = response
	rd.Error = errorstring

	return rd
}
