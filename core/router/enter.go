package router

import (
	"equity/core/router/adminrouter/auth"
	"equity/core/router/adminrouter/system"
)

type AllRouterGroup struct {
	// 管理后台api
	Auth   auth.RouterGroup         // 权限管理
	System system.AdminSystemRouter // 系统

}

var AllRouterGroupApp = new(AllRouterGroup)
