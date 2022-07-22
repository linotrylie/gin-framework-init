package auth

import (
	"equity/core/model/authmodel/request"
	"equity/core/model/commonreq"
	"equity/core/model/commonres"
	"equity/global"
	"github.com/gin-gonic/gin"
)

// RoleApi 角色相关Api
type RoleApi struct{}

// @Tags rbac/role
// @Summary 新增角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param request body request.RoleCreateParam true "新增角色参数"
// @Success 200 {string} string ""{"code":0,"data":null,"msg":"ok"}""
// @Router /admin/role/create [post]
func (l *RoleApi) CreateRole(c *gin.Context) {
	param := request.RoleCreateParam{}
	err := c.ShouldBind(&param)
	if err != nil {
		global.Logger.ErrorF("/role/create,(c,&param) errs: %v", err)
		commonres.FailWithCodeMessage(commonres.ErrorParam, err.Error(), c)
		return
	}

	err = param.Validate()
	if err != nil {
		global.Logger.ErrorF("param.Validate() errs: %v", err)
		commonres.FailWithCodeMessage(commonres.ErrorParamValidate, err.Error(), c)
		return
	}

	err = rbacService.CreateRole(c, &param)
	if err != nil {
		global.Logger.ErrorF(" rbacService.CreateRole err: %v", err)
		commonres.FailWithCodeMessage(commonres.ErrorCreateRole, err.Error(), c)
		return
	}
	commonres.Ok(c)
}

// @Tags rbac/role
// @Summary 分页获取角色列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string ""{"code":0,"data":null,"msg":"ok"}""
// @Router /admin/role/list [post]
func (t *UserApi) RoleList(c *gin.Context) {
	param := request.SearchRoleListParam{}
	err := c.ShouldBind(&param)
	if err != nil {
		global.Logger.ErrorF("/role/list [post] (c,&param) errs: %v", err)
		commonres.FailWithCodeMessage(commonres.ErrorParam, err.Error(), c)
		return
	}

	if err = param.Validate(); err != nil {
		global.Logger.ErrorF("/role/list,param.Validate() errs: %v", err)
		commonres.FailWithCodeMessage(commonres.ErrorParamValidate, err.Error(), c)
		return
	}

	err, list, total := rbacService.RoleList(c, &param)
	if err != nil {
		global.Logger.ErrorF("rbacService.RoleList err: %v", err)
		commonres.FailWithCodeMessage(commonres.ErrorGetRoleList, err.Error(), c)
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

// @Tags rbac/role
// @Summary 获取角色列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string ""{"code":0,"data":null,"msg":"ok"}""
// @Router /admin/role/all [post]
func (t *UserApi) RoleAll(c *gin.Context) {
	err, list, _ := rbacService.RoleAll(c)
	if err != nil {
		global.Logger.ErrorF("rbacService.RoleList err: %v", err)
		commonres.FailWithCodeMessage(commonres.ErrorGetRoleList, err.Error(), c)
		return
	}
	commonres.OkWithDetailed(list, commonres.GetDataSuccess, c)
}

// @Tags rbac/role
// @Summary 获取角色列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string ""{"code":0,"data":null,"msg":"ok"}""
// @Router /admin/role/delete [delete]
func (t *UserApi) RoleDelete(c *gin.Context) {
	param := commonreq.PrimaryIdParam{}
	err := c.ShouldBind(&param)
	if err != nil {
		global.Logger.ErrorF("/role/delete [post] (c,&param) errs: %v", err)
		commonres.FailWithCodeMessage(commonres.ErrorParam, err.Error(), c)
		return
	}

	if err = param.Validate(); err != nil {
		global.Logger.ErrorF("/role/delete,param.Validate() errs: %v", err)
		commonres.FailWithCodeMessage(commonres.ErrorParamValidate, err.Error(), c)
		return
	}
	err = rbacService.RoleDelete(c, &param)
	if err != nil {
		global.Logger.ErrorF("rbacService.RoleDelete err: %v", err)
		commonres.FailWithCodeMessage(commonres.ErrorDeleteUser, err.Error(), c)
		return
	}
	commonres.Ok(c)
}

// @Tags rbac/role
// @Summary 编辑角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string ""{"code":0,"data":null,"msg":"ok"}""
// @Router /admin/role/update [post]
func (t *UserApi) RoleUpdate(c *gin.Context) {
	param := request.RoleUpdateParam{}
	err := c.ShouldBind(&param)
	if err != nil {
		global.Logger.ErrorF("/role/update [post] (c,&param) errs: %v", err)
		commonres.FailWithCodeMessage(commonres.ErrorParam, err.Error(), c)
		return
	}

	if err = param.Validate(); err != nil {
		global.Logger.ErrorF("/role/update,param.Validate() errs: %v", err)
		commonres.FailWithCodeMessage(commonres.ErrorParamValidate, err.Error(), c)
		return
	}
	err = rbacService.RoleUpdate(c, &param)
	if err != nil {
		global.Logger.ErrorF("rbacService.RoleUpdate err: %v", err)
		commonres.FailWithCodeMessage(commonres.ErrorUpdateRole, err.Error(), c)
		return
	}
	commonres.Ok(c)
}

// @Tags rbac/role
// @Summary 获取角色详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string ""{"code":0,"data":null,"msg":"ok"}""
// @Router /admin/role/detail [post]
func (t *UserApi) RoleDetail(c *gin.Context) {
	param := commonreq.PrimaryIdParam{}
	err := c.ShouldBind(&param)
	if err != nil {
		global.Logger.ErrorF("/role/node [post] (c,&param) errs: %v", err)
		commonres.FailWithCodeMessage(commonres.ErrorParam, err.Error(), c)
		return
	}

	if err = param.Validate(); err != nil {
		global.Logger.ErrorF("/role/node,param.Validate() errs: %v", err)
		commonres.FailWithCodeMessage(commonres.ErrorParamValidate, err.Error(), c)
		return
	}
	nodeIds, err := rbacService.RoleDetail(c, &param)
	if err != nil {
		global.Logger.ErrorF("rbacService.RoleUpdate err: %v", err)
		commonres.FailWithCodeMessage(commonres.ErrorUpdateRole, err.Error(), c)
		return
	}
	commonres.OkWithData(nodeIds, c)
}
