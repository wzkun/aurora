package errstring

import (
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/wzkun/aurora/consts"
)

const (
	ItemNotExist            = "record_not_exist.domain.app_error"
	ItemAlreadyExist        = "record_already_exist.domain.app_error"
	MarshalToElasticError   = "marshal_to_elastic_error.domain.app_error"
	PushMessageToKafkaError = "push_message_to_kafka_error.domain.app_error"

	CreateElasticSearchDataError = "create_elastic_search_data_error.domain.app_error"
	DeleteElasticSearchDataError = "delete_elastic_search_data_error.domain.app_error"
	UpdateElasticSearchDataError = "update_elastic_search_data_error.domain.app_error"

	CreateMysqlDataError = "create_mysql_data_error.domain.app_error"
	DeleteMysqlDataError = "delete_mysql_data_error.domain.app_error"
	UpdateMysqlDataError = "update_mysql_data_error.domain.app_error"
	QueryMysqlDataError  = "query_mysql_data_error.domain.app_error"
)

// MakeItemNotExistDetail
func MakeItemNotExistDetail(id, apiName string) error {
	detail := "数据不存在【" + id + "】"
	return NewClientErr(nil, ItemNotExist, detail, apiName, nil)
}

// MakeErrorDebug
func MakeErrorDebug(apiName string, sort int, err error) {
	apiName = "--**== " + apiName + "==**-- " + string(sort) + " :  "
	logrus.Debugln(apiName, err)
}

// MakeResponseError
func MakeResponseError(code, detail, apiName string) error {
	return NewClientErr(nil, code, detail, apiName, nil)
}

// MakeResponseError2
func MakeResponseError2(code, detail, apiName string, err error) error {
	if strings.Contains(err.Error(), "Duplicate") {
		newdetail := consts.ItemAlreadyExistDetail + ": " + err.Error()
		return NewClientErr(nil, ItemAlreadyExist, newdetail, apiName, nil)
	}
	if err == gorm.ErrRecordNotFound {
		return NewClientErr(nil, ItemNotExist, consts.ItemNotExistDetail, apiName, nil)
	}

	return MakeResponseError(code, detail, apiName)
}
