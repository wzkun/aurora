package consts

const (
	QueryUserDetailByMenu = `
		select b.*
		from rbac_server_usermenurelation a
		left join rbac_server_user b 
		on a.account = b.id
		where a.menuid = ?
	`

	QueryGroupDetailByMenu = `
		select b.*
		from rbac_server_groupmenurelation a
		left join rbac_server_group b 
		on a.groupid = b.id
		where a.menuid = ?
	`

	QueryRoleDetailByMenu = `
		select b.*
		from rbac_server_rolemenurelation a
		left join rbac_server_role b 
		on a.roleid = b.id
		where a.menuid = ?
	`

	QueryFunctionDetailByMenu = `
		select b.*
		from rbac_server_menufunctionrelation a
		left join rbac_server_function b 
		on a.funcid = b.id
		where a.menuid = ?
	`
)
