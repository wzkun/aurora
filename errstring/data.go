package errstring

import (
	"strings"

	"code.aliyun.com/bim_backend/zoogoer/gun/errors"
	"github.com/sirupsen/logrus"
	"github.com/wzkun/aurora/consts"
)

const (
	ItemNotExist            = "record_not_exist.domain.app_error"
	ItemAlreadyExist        = " record_already_exist.domain.app_error"
	MarshalToElasticError   = "marshal_to_elastic_error.domain.app_error"
	PushMessageToKafkaError = "push_message_to_kafka_error.domain.app_error"
)

// MakeItemNotExistDetail
func MakeItemNotExistDetail(id, apiName string) error {
	detail := "数据不存在【" + id + "】"
	return errors.NewClientErr(nil, ItemNotExist, detail, apiName, nil)
}

// MakeErrorDebug
func MakeErrorDebug(apiName string, sort int, err error) {
	apiName = "--**== " + apiName + "==**-- " + string(sort) + " :  "
	logrus.Debugln(apiName, err)
}

// MakeResponseError
func MakeResponseError(code, detail, apiName string) error {
	return errors.NewClientErr(nil, code, detail, apiName, nil)
}

// MakeResponseError2
func MakeResponseError2(code, detail, apiName string, err error) error {
	if strings.Contains(err.Error(), "Duplicated") {
		newdetail := consts.ItemAlreadyDetail + ": " + err.Error()
		return errors.NewClientErr(nil, ItemAlreadyExist, newdetail, apiName, nil)
	}

	return MakeResponseError(code, detail, apiName)
}
