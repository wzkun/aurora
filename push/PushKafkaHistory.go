package push

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
	"github.com/wzkun/aurora/consts"
	"github.com/wzkun/aurora/errstring"
)

// PushKafkaHistory
type PushKafkaHistory struct {
	ProjectId  string `json:"projectId" comment:"项目id" optional:"false"`
	MappingKey string `json:"mappingKey" comment:"MappingKey" optional:"false"`
	DataClass  string `json:"dataClass" comment:"DataClass" optional:"false"`
	OperatorId string `json:"operatorId" comment:"操作人id" optional:"false"`
	Item       string `json:"item" comment:"Item" optional:"false"`
	Operation  int    `json:"operation" comment:"Operation" optional:"false"`
}

// PushMessageToKafka
func PushMessageToKafka(projectId, mappingKey, dataClass, url, operator string, item string, operation int) error {
	req := &fasthttp.Request{} //相当于获取一个对象
	req.SetRequestURI(url)     //设置请求的url

	args := &PushKafkaHistory{}
	args.ProjectId = projectId
	args.MappingKey = mappingKey
	args.DataClass = dataClass
	args.Item = item
	args.OperatorId = operator
	args.Operation = int(operation)

	bargs, _ := json.Marshal(args)
	req.SetBody(bargs)
	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")

	resp := &fasthttp.Response{} //相应结果的对象
	client := &fasthttp.Client{}
	err := client.Do(req, resp)
	if err != nil {
		errstring.MakeErrorDebug(consts.PushMessageToKafka, 1, err)
		return errstring.MakeResponseError(errstring.PushMessageToKafkaError, consts.PushMessageToKafkaDetail, consts.PushMessageToKafka)
	}

	return nil
}
