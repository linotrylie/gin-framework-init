package auth

import (
	"equity/core/api/admin"
	"github.com/gin-gonic/gin"
)

type NotValidateAuthRouter struct{}

func (v *NotValidateAuthRouter) InitNotValidateAuthRouter(Router *gin.RouterGroup) {
	BaseRouter := Router.Group("user")
	UserApiGroup := admin.ApiGroupApp.RbacApiGroup
	{
		BaseRouter.POST("login", UserApiGroup.Login)
	}
}
