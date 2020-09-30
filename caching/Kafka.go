package caching

import "github.com/wzkun/aurora/utils/decode"

// MarshalItemToKafkaJSON 序列化为KAFKA JSON
func MarshalItemToKafkaJson(o Item) ([]byte, error) {
	results := make([]interface{}, 0)
	data := make(map[string]interface{})

	result, _ := o.MarshalToJson()

	var v interface{}
	decode.JSON.Unmarshal(result, &v)
	results = append(results, v)

	data["tunnelWyData"] = results
	return decode.JSON.Marshal(data)
}

// MarshalMultiItemToKafkaJson 序列化为KAFKA JSON
func MarshalMultiItemToKafkaJson(os []Item) ([]byte, error) {
	results := make([]interface{}, 0)
	data := make(map[string]interface{})
	for _, o := range os {
		result, _ := o.MarshalToJson()
		var v interface{}
		decode.JSON.Unmarshal(result, &v)
		results = append(results, v)
	}

	data["tunnelWyData"] = results
	return decode.JSON.Marshal(data)
}
