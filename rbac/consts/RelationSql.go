package consts

const (
	QueryRoleBrief = `
		select a.roleid as id, b.name as name 
		from rbac_server_userrolerelation a 
		left join rbac_server_role b 
		on a.roleid = b.id 
		where a.account = %s 
	`

	QueryRoleDetail = `
		select b.* 
		from rbac_server_userrolerelation a 
		left join rbac_server_role b 
		on a.roleid = b.id 
		where a.account = %s 
	`

	QueryOrganizationBrief = `
		select a.orgid as id, b.name ad name 
		from rbac_server_userorganizationrelation a 
		left join rbac_server_organization b 
		on a.orgid = b.id 
		where a.account = %s 
	`

	QueryOrganizationDetail = `
		select b.* 
		from rbac_server_userorganizationrelation a 
		left join rbac_server_organization b 
		on a.orgid = b.id 
		where a.account = %s 
	`

	QueryMenuBrief = `
		select a.menuid as id, b.name ad name 
		from rbac_server_usermenurelation a 
		left join rbac_server_menu b 
		on a.menuid = b.id 
		where a.account = %s 
	`

	QueryMenuDetail = `
		select b.* 
		from rbac_server_usermenurelation a 
		left join rbac_server_menu b 
		on a.menuid = b.id 
		where a.account = %s 
	`

	QueryGroupBrief = `
		select a.groupid as id, b.name ad name 
		from rbac_server_usergrouprelation a 
		left join rbac_server_group b 
		on a.groupid = b.id 
		where a.account = %s 
	`

	QueryGroupDetail = `
		select b.* 
		from rbac_server_usergrouprelation a 
		left join rbac_server_group b 
		on a.groupid = b.id 
		where a.account = %s 
	`

	QueryFunctionBrief = `
		select a.funcid as id, b.name ad name 
		from rbac_server_userfunctionrelation a 
		left join rbac_server_function b 
		on a.funcid = b.id 
		where a.account = %s 
	`
	
	QueryFunctionDetail = `
	select b.* 
	from rbac_server_userfunctionrelation a 
	left join rbac_server_function b 
	on a.funcid = b.id 
	where a.account = %s 
`

)
