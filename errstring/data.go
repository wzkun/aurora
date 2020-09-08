package errstring

import (
	"code.aliyun.com/bim_backend/zoogoer/gun/errors"
)

const (
	ItemNotExist = "record_not_exist.domain.app_error"
)

// MakeItemNotExistDetail
func MakeItemNotExistDetail(id, apiName string) error {
	detail := "数据不存在【" + id + "】"
	err := errors.NewClientErr(nil, ItemNotExist, detail, apiName, nil)

	return err
}
