packages consts

const(
	QueryUserDetailByGroup=`
		select b.*
		from rbac_server_usergrouprelation a
		left join rbac_server_user b 
		on a.account = b.id
		where a.groupid = %s
	`
	
	QueryRoleDetailByGroup=`
		select b.*
		from rbac_server_grouprolerelation a
		left join rbac_server_role b 
		on a.roleid = b.id
		where a.groupid = %s
	`

	QueryMenuDetailByGroup=`
		select b.*
		from rbac_server_groupmenurelation a
		left join rbac_server_menu b 
		on a.menuid = b.id
		where a.groupid = %s
	`
	
	QueryFunctionDetailByGroup=`
		select b.*
		from rbac_server_groupfunctionrelation a
		left join rbac_server_function b 
		on a.funcid = b.id
		where a.groupid = %s
	`

)