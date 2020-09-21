package errstring

import (
	"code.aliyun.com/bim_backend/zoogoer/gun/errors"
	"github.com/sirupsen/logrus"
)

const (
	ItemNotExist = "record_not_exist.domain.app_error"
)

// MakeItemNotExistDetail
func MakeItemNotExistDetail(id, apiName string) error {
	detail := "数据不存在【" + id + "】"
	return errors.NewClientErr(nil, ItemNotExist, detail, apiName, nil)
}

// MakeErrorDebug
func MakeErrorDebug(apiName string, sort int, err error) {
	apiName = "--**== " + apiName + "--**== " + string(sort) + " :  "
	logrus.Debugln(apiName, err)
}

// MakeResponseError
func MakeResponseError(code, detail, apiName string) error {
	return errors.NewClientErr(nil, code, detail, apiName, nil)
}
