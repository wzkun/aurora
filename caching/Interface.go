package caching

import (
	"code.aliyun.com/new_backend/scodi_nqc/elastic"
	"github.com/wzkun/aurora/utils/decode"
)

type Item interface {
	Idx() string
	ElasticIndex() string
	MarshalToJson() ([]byte, error)
	TableName() string
}

// ModelMarshalToElastic 序列化为ES对象
func ModelMarshalToElastic(o Item) (elastic.ElasticItem, error) {
	idx := o.Idx()
	parrent := o.ElasticIndex()

	body, err := o.MarshalToJson()
	if err != nil {
		return nil, err
	}
	return elastic.NewElasticItem2(idx, parrent, body)
}

// MultiMarshalToElastic 序列化为ES对象
func MultiMarshalToElastic(os ...Item) ([]elastic.ElasticItem, error) {
	rds := make([]elastic.ElasticItem, 0)

	for _, o := range os {
		idx := o.Idx()
		parrent := o.ElasticIndex()

		body, err := o.MarshalToJson()
		if err != nil {
			return nil, err
		}

		rd, err := elastic.NewElasticItem2(idx, parrent, body)
		if err != nil {
			return nil, err
		}

		rds = append(rds, rd)
	}

	return rds, nil
}

type ModelItem interface {
	MarshalToJson() ([]byte, error)
}

// MarshalItemToKafkaJSON 序列化为KAFKA JSON
func MarshalItemToKafkaJson(o ModelItem) ([]byte, error) {
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
func MarshalMultiItemToKafkaJson(os []ModelItem) ([]byte, error) {
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
