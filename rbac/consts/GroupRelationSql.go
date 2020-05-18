package consts

const (
	// 查询用户组下的人员列表
	QueryUserDetailByGroup = `
		select b.*
		from rbac_server_usergrouprelation a
		left join rbac_server_user b 
		on a.account = b.id
		where a.groupid = ?
	`
	// 查询用户组分配的角色列表
	QueryRoleDetailByGroup = `
		select b.*
		from rbac_server_grouprolerelation a
		left join rbac_server_role b 
		on a.roleid = b.id
		where a.groupid = ?
	`
	// 查询用户组分配的菜单权限
	QueryMenuDetailByGroup = `
		select b.*
		from rbac_server_groupmenurelation a
		left join rbac_server_menu b 
		on a.menuid = b.id
		where a.groupid = ?
	`
	// 查询用户组分配的功能权限
	QueryFunctionDetailByGroup = `
		select b.*
		from rbac_server_groupfunctionrelation a
		left join rbac_server_function b 
		on a.funcid = b.id
		where a.groupid = ?
	`
)
