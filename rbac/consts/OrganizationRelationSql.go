package consts

const (
	QueryUserDetailByOrganization = `
		select b.* 
		from rbac_server_userorganizationrelation a 
		left join rbac_server_user b 
		on a.account = b.id 
		where a.orgid = %s 
	`
)
