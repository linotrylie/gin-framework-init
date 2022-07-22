package commonres

// 管理后台&&商户管理后台错误码
var (
	ErrorParamValidate = NewErrorResp(20000000, "参数验证未通过")
	ErrorParam         = NewErrorResp(20010000, "参数解析失败")
	ErrorLogin         = NewErrorResp(20010001, "登录失败")
	ErrorCreateUser    = NewErrorResp(20010002, "新增管理员失败")
	ErrorSetUserStatus = NewErrorResp(20010003, "设置管理员状态失败")
	ErrorDeleteUser    = NewErrorResp(20010004, "删除用户失败")
	ErrorChangePWD     = NewErrorResp(20010005, "修改密码失败")
	ErrorCreateRole    = NewErrorResp(20010006, "新增角色失败")
	ErrorGetUserList   = NewErrorResp(20010007, "获取用户列表失败")
	ErrorGetRoleList   = NewErrorResp(20010008, "获取角色列表失败")
	ErrorUpdateRole    = NewErrorResp(20010009, "编辑角色失败")
	ErrorUpdateUser    = NewErrorResp(20010010, "编辑管理员失败")
	ErrorCreateNode    = NewErrorResp(20010011, "创建节点失败")
	ErrorGetNode       = NewErrorResp(20010012, "节点树获取失败")
	ErrorGetUserRole   = NewErrorResp(20010013, "用户角色获取失败")
	ErrorNotAuth       = NewErrorResp(20010014, "无权限")
	ErrorGetNodeDetail = NewErrorResp(20010017, "节点详情获取失败")
)
