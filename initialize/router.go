package initialize

import (
	"equity/core/consts"
	"equity/core/middleware"
	"equity/core/model/commonres"
	"equity/core/router"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// Routers 初始化总路由
func Routers() *gin.Engine {
	var Router = gin.Default()
	Router.Use(gin.Logger())
	Router.Use(gin.Recovery())
	Router.Use(middleware.Recovery())
	Router.Use(middleware.AccessLog())
	Router.Use(middleware.Tracing())
	Router.Use(middleware.Cors())
	Router.GET("/", func(c *gin.Context) {})
	Router.GET("/favicon.ico", func(c *gin.Context) {})
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	Router.NoRoute(func(c *gin.Context) {
		//msg := fmt.Sprintf("404 not found,runMode=%s,method=%s,url=%s",
		//	gin.Mode(),
		//	c.Request.Method,
		//	c.Request.Host+c.Request.RequestURI,
		//)
		//go message.SendMsgToFeiShu(msg)
		commonres.FailWithMessage("404 not found", c)
	})

	// 管理后台路由组
	adminAuthRouter := router.AllRouterGroupApp.Auth
	adminSystemRouter := router.AllRouterGroupApp.System

	// 不需要鉴权的路由组
	adminPublicGroup := Router.Group(consts.AdminApiPrefix)
	{
		// 接受供应商的异步通知,不需要建权
		adminAuthRouter.InitNotValidateAuthRouter(adminPublicGroup)
		adminSystemRouter.InitHealthRouter(adminPublicGroup)
	}

	// 需要鉴权的路由组
	adminPrivateGroup := Router.Group(consts.AdminApiPrefix)
	adminPrivateGroup.Use(middleware.JWT()).Use(middleware.RBAC())
	{
		// 权限
		adminAuthRouter.InitRbacRouter(adminPrivateGroup)
	}
	return Router
}
