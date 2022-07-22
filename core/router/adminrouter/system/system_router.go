package system

import (
	"equity/core/api/admin"
	"github.com/gin-gonic/gin"
)

type AdminSystemRouter struct{}

func (s *AdminSystemRouter) InitHealthRouter(Router *gin.RouterGroup) {
	BaseRouter := Router.Group("system")
	systemApi := admin.ApiGroupApp.SystemApiGroup
	{
		BaseRouter.GET("health", systemApi.Health)
	}
}
