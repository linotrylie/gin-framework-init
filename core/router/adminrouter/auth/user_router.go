package auth

import (
	"equity/core/api/admin"
	"github.com/gin-gonic/gin"
)

type RbacRouter struct{}

func (v *RbacRouter) InitRbacRouter(Router *gin.RouterGroup) {
	// 用户
	userRouter := Router.Group("user")
	// 角色
	roleRouter := Router.Group("role")
	// 节点
	nodeRouter := Router.Group("node")
	rbacApiGroup := admin.ApiGroupApp.RbacApiGroup
	{
		// user
		userRouter.POST("create", rbacApiGroup.CreateUser)
		userRouter.POST("update", rbacApiGroup.UpdateUser)
		userRouter.POST("status", rbacApiGroup.SetUserStatus)
		userRouter.DELETE("delete", rbacApiGroup.DeleteUser)
		userRouter.POST("change/password", rbacApiGroup.ChangeUserPassword)
		userRouter.POST("list", rbacApiGroup.UserList)
		userRouter.POST("detail", rbacApiGroup.UserDetail)
		userRouter.POST("delete", rbacApiGroup.DeleteUser)
		userRouter.POST("node", rbacApiGroup.UserNode)

		// role
		roleRouter.POST("create", rbacApiGroup.CreateRole)
		roleRouter.POST("list", rbacApiGroup.RoleList)
		roleRouter.DELETE("delete", rbacApiGroup.RoleDelete)
		roleRouter.POST("update", rbacApiGroup.RoleUpdate)
		roleRouter.POST("detail", rbacApiGroup.RoleDetail)
		roleRouter.POST("all", rbacApiGroup.RoleAll)

		// node
		nodeRouter.POST("create", rbacApiGroup.CreateNode)
		nodeRouter.POST("tree", rbacApiGroup.NodeTree)
		nodeRouter.POST("update", rbacApiGroup.UpdateNode)
		nodeRouter.POST("detail", rbacApiGroup.NodeDetail)
	}
}
