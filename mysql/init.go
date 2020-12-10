package mysql

import "code.aliyun.com/new_backend/scodi_nqc/model"

func init() {
	{
		uSql := &ApiAccessHistory{}
		model.ModelInstance().CheckTable(uSql, nil)
	}
}
