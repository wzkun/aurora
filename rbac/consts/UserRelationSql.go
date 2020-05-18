package consts

const (
	QueryRoleBriefByUser = `
		select a.roleid as id, b.name as name 
		from rbac_server_userrolerelation a 
		left join rbac_server_role b 
		on a.roleid = b.id 
		where a.account = %s 
	`

	QueryRoleDetailByUser = `
		select b.* 
		from rbac_server_userrolerelation a 
		left join rbac_server_role b 
		on a.roleid = b.id 
		where a.account = %s 
	`

	QueryOrganizationBriefByUser = `
		select a.orgid as id, b.name as name 
		from rbac_server_userorganizationrelation a 
		left join rbac_server_organization b 
		on a.orgid = b.id 
		where a.account = %s 
	`

	QueryOrganizationDetailByUser = `
		select b.* 
		from rbac_server_userorganizationrelation a 
		left join rbac_server_organization b 
		on a.orgid = b.id 
		where a.account = %s 
	`

	QueryMenuBriefByUser = `
		select a.menuid as id, b.name as name 
		from rbac_server_usermenurelation a 
		left join rbac_server_menu b 
		on a.menuid = b.id 
		where a.account = %s 
	`

	QueryMenuDetailByUser = `
		select b.* 
		from rbac_server_usermenurelation a 
		left join rbac_server_menu b 
		on a.menuid = b.id 
		where a.account = %s 
	`

	QueryGroupBriefByUser = `
		select a.groupid as id, b.name as name 
		from rbac_server_usergrouprelation a 
		left join rbac_server_group b 
		on a.groupid = b.id 
		where a.account = %s 
	`

	QueryGroupDetailByUser = `
		select b.* 
		from rbac_server_usergrouprelation a 
		left join rbac_server_group b 
		on a.groupid = b.id 
		where a.account = %s 
	`

	QueryFunctionBriefByUser = `
		select a.funcid as id, b.name as name 
		from rbac_server_userfunctionrelation a 
		left join rbac_server_function b 
		on a.funcid = b.id 
		where a.account = %s 
	`

	QueryFunctionDetailByUser = `
		select b.* 
		from rbac_server_userfunctionrelation a 
		left join rbac_server_function b 
		on a.funcid = b.id 
		where a.account = %s 
	`
)
