package service

import (
	"equity/core/service/admin/auth"
)

type ServicesGroup struct {
	// 管理后台服务分组
	RBACServiceGroup auth.ServiceGroup
}

var AllServiceGroupApp = new(ServicesGroup)
