package utils

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/sirupsen/logrus"
	"github.com/wzkun/aurora/errstring"
	"github.com/wzkun/aurora/mysql"
	"github.com/wzkun/aurora/utils/decode"
)

// TraceRPCRequest Function
func TraceRPCRequest(servcie, method string, req interface{}) {
	logrus.WithFields(logrus.Fields{
		"Service": servcie,
		"Methold": method,
		"Step":    "PreRequest",
		"Request": req,
	}).Info()
}

// TraceRPCResponse Function
func TraceRPCResponse(operator, moduleName, servcie, method string, err error, req, resp proto.Message) {
	request, _ := decode.PROTO.MarshalToJSON(req)
	response, _ := decode.PROTO.MarshalToJSON(resp)

	go func() {
		errorstring := ""
		if err != nil {
			errorstring = err.Error()
		}
		RecordApiAccessHistory(operator, moduleName, servcie, method, string(request), string(response), errorstring)
	}()

	logrus.WithFields(logrus.Fields{
		"Service":  servcie,
		"Methold":  method,
		"Step":     "PostRequest",
		"operator": operator,
		"Response": &response,
		"Error":    err,
	}).Info()
}

// TraceHttpResponse Function
func TraceHttpResponse(operator, moduleName, servcie, method string, err error, req, resp interface{}) {
	request, _ := json.Marshal(req)
	response, _ := json.Marshal(resp)

	go func() {
		errorstring := ""
		if err != nil {
			errorstring = err.Error()
		}
		RecordApiAccessHistory(operator, moduleName, servcie, method, string(request), string(response), errorstring)
	}()

	logrus.WithFields(logrus.Fields{
		"Service":  servcie,
		"Methold":  method,
		"Step":     "PostRequest",
		"operator": operator,
		"Response": resp,
		"Error":    err,
	}).Info()
}

// RecordApiAccessHistory function
func RecordApiAccessHistory(account, moduleName, serviceName, apiName, request, response, errorstring string) {
	rd := &mysql.ApiAccessHistory{}

	rd.Id = NewUUIdV4()
	rd.Account = account
	rd.ServerName = ""
	rd.ModuleName = moduleName
	rd.ServiceName = serviceName
	rd.ApiName = apiName
	rd.CreateTime = time.Now().String()
	rd.Request = request

	if strings.Contains(apiName, "CreateMulti") || strings.Contains(apiName, "MergeMulti") {
		rd.Response = ""
	} else if strings.Contains(apiName, "Query") || strings.Contains(apiName, "Search") {
		rd.Response = ""
	} else {
		rd.Response = response
	}

	rd.Error = errorstring

	err := rd.Insert()
	if err != nil {
		errstring.MakeErrorDebug("RecordApiAccessHistory", 1, err)
	}
}
