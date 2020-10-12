package errstring

import (
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/wzkun/aurora/consts"
)

const (
	ItemNotExist     = "record_not_exist.domain.app_error"
	ItemAlreadyExist = "record_already_exist.domain.app_error"
	RecordNotFound   = "record not found"
	RecordDuplicate  = "Duplicate"

	MarshalToElasticError       = "marshal_to_elastic_error.domain.app_error"
	PushMessageToKafkaError     = "push_message_to_kafka_error.domain.app_error"
	RecordApiAccessHistoryError = "record_api_access_history_error.domain.app_error"

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
func MakeErrorDebug(apiName string, sort int64, err error) {
	apiName = "--**== " + apiName + "==**-- " + strconv.FormatInt(sort, 10) + " :  "
	logrus.Debugln(apiName, err)
}

// MakeResponseError
func MakeResponseError(code, detail, apiName string) error {
	return NewClientErr(nil, code, detail, apiName, nil)
}

// MakeResponseError3
func MakeResponseError3(err error, code, detail, apiName string) error {
	return NewClientErr(err, code, detail, apiName, nil)
}

// MakeResponseError2
func MakeResponseError2(code, detail, apiName string, err error) error {
	if strings.Contains(err.Error(), RecordDuplicate) {
		newdetail := consts.ItemAlreadyExistDetail + ": " + err.Error()
		return MakeResponseError(ItemAlreadyExist, newdetail, apiName)
	}

	if strings.Contains(err.Error(), RecordNotFound) {
		return MakeResponseError(ItemNotExist, consts.ItemNotExistDetail, apiName)
	}

	return MakeResponseError(code, detail, apiName)
}
