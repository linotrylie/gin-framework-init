package auth

import (
	"equity/core/model/authmodel/request"
	"equity/core/model/commonreq"
	"equity/core/model/commonres"
	"equity/global"
	"github.com/gin-gonic/gin"
)

type UserApi struct{}

// @Tags rbac/user
// @Summary 用户登录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"code":0,"data":{"expire":1651654641,"token":"xxxx"},"msg":"ok"}"
// @Router /admin/user/login [post]
func (u *UserApi) Login(c *gin.Context) {
	param := request.UserLoginParam{}
	err := c.ShouldBind(&param)
	if err != nil {
		global.Logger.ErrorF("/user/login [post] (c,&param) errs: %v", err)
		commonres.FailWithCodeMessage(commonres.ErrorParam, err.Error(), c)
		return
	}

	if err = param.Validate(); err != nil {
		global.Logger.ErrorF("/user/login,param.Validate() errs: %v", err)
		commonres.FailWithCodeMessage(commonres.ErrorParamValidate, err.Error(), c)
		return
	}
	token, exp, userName, err := rbacService.Login(c, &param)
	if err != nil {
		global.Logger.ErrorF(" rbacService.Login err: %v", err)
		commonres.FailWithCodeMessage(commonres.ErrorLogin, err.Error(), c)
		return
	}
	commonres.OkWithData(gin.H{"token": token, "expire": exp, "userName": userName}, c)
}

// @Tags rbac/user
// @Summary 当前用户所拥有的角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"code":0,"data":{"expire":1651654641,"token":"xxxx"},"msg":"ok"}"
// @Router /admin/user/detail [post]
func (u *UserApi) UserDetail(c *gin.Context) {
	param := commonreq.PrimaryIdParam{}
	err := c.ShouldBind(&param)
	if err != nil {
		global.Logger.ErrorF("/user/detail [post] (c,&param) errs: %v", err)
		commonres.FailWithCodeMessage(commonres.ErrorParam, err.Error(), c)
		return
	}

	if err = param.Validate(); err != nil {
		global.Logger.ErrorF("/user/detail,param.Validate() errs: %v", err)
		commonres.FailWithCodeMessage(commonres.ErrorParamValidate, err.Error(), c)
		return
	}
	detail, err := rbacService.UserDetail(c, &param)
	if err != nil {
		global.Logger.ErrorF(" rbacService.detail err: %v", err)
		commonres.FailWithCodeMessage(commonres.ErrorGetUserRole, err.Error(), c)
		return
	}
	commonres.OkWithData(detail, c)
}

// @Tags rbac/user
// @Summary 新增用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string ""{"code":0,"data":null,"msg":"ok"}""
// @Router /admin/user/create [post]
func (u *UserApi) CreateUser(c *gin.Context) {
	param := request.CreateUserParam{}
	err := c.ShouldBind(&param)
	if err != nil {
		global.Logger.ErrorF("/user/create [post] (c,&param) errs: %v", err)
		commonres.FailWithCodeMessage(commonres.ErrorParam, err.Error(), c)
		return
	}

	if err = param.Validate(); err != nil {
		global.Logger.ErrorF("/user/create,param.Validate() errs: %v", err)
		commonres.FailWithCodeMessage(commonres.ErrorParamValidate, err.Error(), c)
		return
	}
	err = rbacService.CreateUser(c, &param)
	if err != nil {
		global.Logger.ErrorF(" rbacService.CreateUser err: %v", err)
		commonres.FailWithCodeMessage(commonres.ErrorCreateUser, err.Error(), c)
		return
	}
	commonres.Ok(c)
}

// @Tags rbac/user
// @Summary 编辑用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string ""{"code":0,"data":null,"msg":"ok"}""
// @Router /admin/user/update [post]
func (u *UserApi) UpdateUser(c *gin.Context) {
	param := request.UpdateUserParam{}
	err := c.ShouldBind(&param)
	if err != nil {
		global.Logger.ErrorF("/user/update [post] (c,&param) errs: %v", err)
		commonres.FailWithCodeMessage(commonres.ErrorParam, err.Error(), c)
		return
	}

	if err = param.Validate(); err != nil {
		global.Logger.ErrorF("/user/update,param.Validate() errs: %v", err)
		commonres.FailWithCodeMessage(commonres.ErrorParamValidate, err.Error(), c)
		return
	}
	err = rbacService.UpdateUser(c, &param)
	if err != nil {
		global.Logger.ErrorF(" rbacService.UpdateUser err: %v", err)
		commonres.FailWithCodeMessage(commonres.ErrorUpdateUser, err.Error(), c)
		return
	}
	commonres.Ok(c)
}

// @Tags rbac/user
// @Summary 设置用户状态
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string ""{"code":0,"data":null,"msg":"ok"}""
// @Router /admin/user/status [post]
func (u *UserApi) SetUserStatus(c *gin.Context) {
	param := request.SetUserStatusParam{}
	err := c.ShouldBind(&param)
	if err != nil {
		global.Logger.ErrorF("/user/status [post] (c,&param) errs: %v", err)
		commonres.FailWithCodeMessage(commonres.ErrorParam, err.Error(), c)
		return
	}

	if err = param.Validate(); err != nil {
		global.Logger.ErrorF("/user/status,param.Validate() errs: %v", err)
		commonres.FailWithCodeMessage(commonres.ErrorParamValidate, err.Error(), c)
		return
	}
	err = rbacService.SetUserStatus(c, &param)
	if err != nil {
		global.Logger.ErrorF("rbacService.SetUserStatus err: %v", err)
		commonres.FailWithCodeMessage(commonres.ErrorSetUserStatus, err.Error(), c)
		return
	}
	commonres.Ok(c)
}

// @Tags rbac/user
// @Summary 用户删除
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string ""{"code":0,"data":null,"msg":"ok"}""
// @Router /admin/user/delete [delete]
func (u *UserApi) DeleteUser(c *gin.Context) {
	param := commonreq.PrimaryIdParam{}
	err := c.ShouldBind(&param)
	if err != nil {
		global.Logger.ErrorF("/user/delete [post] (c,&param) errs: %v", err)
		commonres.FailWithCodeMessage(commonres.ErrorParam, err.Error(), c)
		return
	}

	if err = param.Validate(); err != nil {
		global.Logger.ErrorF("/user/delete,param.Validate() errs: %v", err)
		commonres.FailWithCodeMessage(commonres.ErrorParamValidate, err.Error(), c)
		return
	}
	err = rbacService.Delete(c, &param)
	if err != nil {
		global.Logger.ErrorF("rbacService.Delete err: %v", err)
		commonres.FailWithCodeMessage(commonres.ErrorDeleteUser, err.Error(), c)
		return
	}
	commonres.Ok(c)
}

// @Tags rbac/user
// @Summary 修改密码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string ""{"code":0,"data":null,"msg":"ok"}""
// @Router /admin/user/change/password [post]
func (u *UserApi) ChangeUserPassword(c *gin.Context) {
	param := request.ChangeUserPWDParam{}
	err := c.ShouldBind(&param)
	if err != nil {
		global.Logger.ErrorF("/user/change/password [post] (c,&param) errs: %v", err)
		commonres.FailWithCodeMessage(commonres.ErrorParam, err.Error(), c)
		return
	}

	if err = param.Validate(); err != nil {
		global.Logger.ErrorF("/user/change/password,param.Validate() errs: %v", err)
		commonres.FailWithCodeMessage(commonres.ErrorParamValidate, err.Error(), c)
		return
	}
	err = rbacService.ChangePassword(c, &param)
	if err != nil {
		global.Logger.ErrorF("rbacService.ChangePassword err: %v", err)
		commonres.FailWithCodeMessage(commonres.ErrorChangePWD, err.Error(), c)
		return
	}
	commonres.Ok(c)
}

// @Tags rbac/user
// @Summary 获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string ""{"code":0,"data":null,"msg":"ok"}""
// @Router /v1/user/list [post]
func (u *UserApi) UserList(c *gin.Context) {
	param := request.SearchUserListParam{}
	err := c.ShouldBind(&param)
	if err != nil {
		global.Logger.ErrorF("/user/list [post] (c,&param) errs: %v", err)
		commonres.FailWithCodeMessage(commonres.ErrorParam, err.Error(), c)
		return
	}

	if err = param.Validate(); err != nil {
		global.Logger.ErrorF("/user/list,param.Validate() errs: %v", err)
		commonres.FailWithCodeMessage(commonres.ErrorParamValidate, err.Error(), c)
		return
	}
	err, list, total := rbacService.UserList(c, &param)
	if err != nil {
		global.Logger.ErrorF("rbacService.UserList err: %v", err)
		commonres.FailWithCodeMessage(commonres.ErrorGetUserList, err.Error(), c)
		return
	}
	pageData := commonres.PageResult{
		List: list,
		Pager: commonres.Pager{
			Total:    total,
			PageNo:   param.PageNo,
			PageSize: param.PageSize,
		},
	}
	commonres.OkWithDetailed(pageData, commonres.GetDataSuccess, c)
}

// @Tags rbac/user
// @Summary 获取用户拥有的节点列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string ""{"code":0,"data":null,"msg":"ok"}""
// @Router /v1/user/node [post]
func (u *UserApi) UserNode(c *gin.Context) {
	tree, err := rbacService.UserNodeTree(c)
	if err != nil {
		global.Logger.ErrorF("rbacService.UserNode err: %v", err)
		commonres.FailWithCodeMessage(commonres.ErrorGetUserList, err.Error(), c)
		return
	}
	commonres.OkWithDetailed(tree, commonres.GetDataSuccess, c)
}
