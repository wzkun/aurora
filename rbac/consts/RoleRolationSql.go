package consts

const (
	QueryUserDetailByRole = `
		select b.*
		from rbac_server_userrolerelation a
		left join rbac_server_user b 
		on a.account = b.id
		where a.roleid = ?
	`

	QueryGroupDetailByRole = `
		select b.*
		from rbac_server_grouprolerelation a
		left join rbac_server_group b 
		on a.groupid = b.id
		where a.roleid = ?
	`

	QueryMenuDetailByRole = `
		select b.*
		from rbac_server_rolemenurelation a
		left join rbac_server_menu b 
		on a.menuid = b.id
		where a.roleid = ?
	`

	QueryFunctionDetailByRole = `
		select b.*
		from rbac_server_rolefunctionrelation a
		left join rbac_server_function b 
		on a.funcid = b.id
		where a.roleid = ?
	`
)
