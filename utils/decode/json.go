package decode

import (
	"encoding/json"
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

var (
	// JSON decoder or encoder
	JSON = jsoniter.ConfigCompatibleWithStandardLibrary
)

// ToJSONStr 对象转换成字符串
// 对象字段必须大写,否则结果为空
func ToJSONStr(data interface{}) (string, error) {
	result, err := json.Marshal(data)
	return fmt.Sprintf("%s", result), err
}

// ToPageJSON 转换成json字符串
func ToPageJSON(datas interface{}, count, pageIndex, pageSize int) (string, error) {
	data, err := json.Marshal(datas)
	result := fmt.Sprintf("{\"rows\":%s,\"pageSize\":%d,\"total\":%d,\"page\":%d}", data, pageSize, count, pageIndex)
	return result, err
}

// Str2Struct 字符串转对象
func Str2Struct(source string, destination interface{}) error {
	err := json.Unmarshal([]byte(source), destination)
	return err
}

// Str2Map 字符转Map
func Str2Map(source string) (map[string]interface{}, error) {
	res := make(map[string]interface{})
	err := json.Unmarshal([]byte(source), &res)
	return res, err
}
