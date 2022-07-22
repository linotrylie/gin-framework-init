package auth

import (
	"equity/core/model/authmodel/request"
	"equity/core/model/commonreq"
	"equity/core/model/commonres"
	"equity/global"
	"equity/utils"
	"github.com/gin-gonic/gin"
)

type NodeApi struct{}

// @Tags rbac/node
// @Summary 新增节点
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string ""{"code":0,"data":null,"msg":"ok"}""
// @Router /admin/node/create [post]
func (n *NodeApi) CreateNode(c *gin.Context) {
	param := request.AuthNodeCreateParam{}
	err := c.ShouldBind(&param)
	if err != nil {
		global.Logger.ErrorF("/node/create,c.ShouldBind(param)errs: %v", err)
		commonres.FailWithCodeMessage(commonres.ErrorParam, err.Error(), c)
		return
	}
	err = param.Validate()
	if err != nil {
		global.Logger.ErrorF("/node/create,param.Validate() errs: %v", err)
		commonres.FailWithCodeMessage(commonres.ErrorParamValidate, err.Error(), c)
		return
	}
	err = rbacService.CreateNode(c, &param)
	if err != nil {
		global.Logger.ErrorF(" rbacService.CreateNode err: %v", err)
		msg := utils.WapErrWithTrace(err, "")
		commonres.FailWithCodeMessage(commonres.ErrorCreateNode, msg, c)
		return
	}
	commonres.Ok(c)
}

// @Tags rbac/node
// @Summary 所有节点列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string ""{"code":0,"data":null,"msg":"ok"}""
// @Router /admin/node/tree [post]
func (n *NodeApi) NodeTree(c *gin.Context) {
	tree, err := rbacService.NodeTree(c)
	if err != nil {
		global.Logger.ErrorF(" rbacService.NodeTree err: %v", err)
		commonres.FailWithCodeMessage(commonres.ErrorGetNode, err.Error(), c)
		return
	}
	commonres.OkWithDetailed(tree, "获取节点树成功", c)
}

// @Summary 编辑节点
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string ""{"code":0,"data":null,"msg":"ok"}""
// @Router /admin/node/update [post]
func (n *NodeApi) UpdateNode(c *gin.Context) {
	param := request.AuthNodeUpdateParam{}
	err := c.ShouldBind(&param)
	if err != nil {
		global.Logger.ErrorF("/node/update,c.ShouldBind(param)errs: %v", err)
		commonres.FailWithCodeMessage(commonres.ErrorParam, err.Error(), c)
		return
	}
	err = param.Validate()
	if err != nil {
		global.Logger.ErrorF("/node/update,param.Validate() errs: %v", err)
		commonres.FailWithCodeMessage(commonres.ErrorGetNodeDetail, err.Error(), c)
		return
	}

	err = rbacService.UpdateNode(c, &param)
	if err != nil {
		global.Logger.ErrorF(" rbacService.UpdateNode err: %v", err)
		commonres.FailWithCodeMessage(commonres.ErrorGetNode, err.Error(), c)
		return
	}
	commonres.Ok(c)
}

// @Summary 编辑节点
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string ""{"code":0,"data":null,"msg":"ok"}""
// @Router /admin/node/detail [post]
func (n *NodeApi) NodeDetail(c *gin.Context) {
	param := commonreq.PrimaryIdParam{}
	err := c.ShouldBind(&param)
	if err != nil {
		global.Logger.ErrorF("/node/detail,c.ShouldBind(param)errs: %v", err)
		commonres.FailWithCodeMessage(commonres.ErrorParam, err.Error(), c)
		return
	}
	err = param.Validate()
	if err != nil {
		global.Logger.ErrorF("/node/detail,param.Validate() errs: %v", err)
		commonres.FailWithCodeMessage(commonres.ErrorGetNodeDetail, err.Error(), c)
		return
	}

	detail, err := rbacService.NodeDetail(c, &param)
	if err != nil {
		global.Logger.ErrorF(" rbacService.NodeDetail err: %v", err)
		commonres.FailWithCodeMessage(commonres.ErrorGetNode, err.Error(), c)
		return
	}
	commonres.OkWithData(detail, c)
}
