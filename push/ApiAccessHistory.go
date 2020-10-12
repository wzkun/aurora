package push

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
	"github.com/wzkun/aurora/consts"
	"github.com/wzkun/aurora/errstring"
)

// ApiAccessHistory
type ApiAccessHistory struct {
	Account     string `json:"account" comment:"账号"`
	ServerName  string `json:"serverName" comment:"微服务名称" optional:"false"`
	ModuleName  string `json:"moduleName" comment:"模块名称" optional:"false"`
	ServiceName string `json:"serviceName" comment:"服务名称" optional:"false"`
	ApiName     string `json:"apiName" comment:"接口名称" optional:"false"`
	Request     string `json:"request" comment:"请求参数" optional:"false"`
	Response    string `json:"response" comment:"相应参数"`
	Errorstring string `json:"errorstring" comment:"错误信息"`
}

// RecordApiAccessHistory
func RecordApiAccessHistory(account, serverName, moduleName, serviceName, apiName, request, response, errorstring, url string) error {
	req := &fasthttp.Request{} //相当于获取一个对象
	req.SetRequestURI(url)     //设置请求的url

	args := &ApiAccessHistory{}
	args.Account = account
	args.ServerName = serverName
	args.ModuleName = moduleName
	args.ServiceName = serviceName
	args.ApiName = apiName
	args.Request = request
	args.Response = response
	args.Errorstring = errorstring

	bargs, _ := json.Marshal(args)
	req.SetBody(bargs)
	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")

	resp := &fasthttp.Response{} //相应结果的对象
	client := &fasthttp.Client{}
	err := client.Do(req, resp)
	if err != nil {
		errstring.MakeErrorDebug(consts.RecordApiAccessHistory, 1, err)
		return errstring.MakeResponseError(errstring.RecordApiAccessHistoryError, consts.RecordApiAccessHistoryDetail, consts.RecordApiAccessHistory)
	}

	return nil
}
