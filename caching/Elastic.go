package caching

import "code.aliyun.com/bim_backend/zoogoer/elastic"

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
