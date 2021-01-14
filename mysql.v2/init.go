package mysql

import "code.aliyun.com/new_backend/scodi_nqc/model.v2"

func init() {
	{
		uSql := &ApiAccessHistory{}
		model.ModelInstance().CheckTable(uSql, nil)
	}
}
