package auth

import "equity/core/service"

type ApiGroup struct {
	UserApi
	RoleApi
	NodeApi
}

var (
	rbacService = service.AllServiceGroupApp.RBACServiceGroup
)
