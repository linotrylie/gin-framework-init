package admin

import (
	"equity/core/api/admin/auth"
	"equity/core/api/admin/system"
)

// AdminApiGroup api 分组
type AdminApiGroup struct {
	SystemApiGroup system.SysApi
	RbacApiGroup   auth.ApiGroup
}

var ApiGroupApp = new(AdminApiGroup)
