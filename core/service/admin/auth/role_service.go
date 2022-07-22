package auth

import (
	"equity/core/dao/authdao"
	"equity/core/model/authmodel"
	"equity/core/model/authmodel/request"
	"equity/core/model/authmodel/response"
	"equity/core/model/commonreq"
	"equity/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// RoleService 角色相关服务
type RoleService struct {
	roleDao *authdao.RoleDao
}

// CreateRole 新增角色
// 编辑角色的时候,将节点id和节点的父级存起来
func (r *RoleService) CreateRole(c *gin.Context, req *request.RoleCreateParam) error {
	if utils.ContainsDuplicate(req.NodeId) {
		return errors.New("节点id不能重复")
	}

	if len(req.Name) <= 0 {
		return errors.New("角色名称不能为空")
	}
	role, err := r.roleDao.GetRole(c, 0, req.Name)
	if err != nil {
		return err
	}

	if role != nil && role.Id > 0 {
		return errors.New(fmt.Sprintf("角色名称不能重复,id=%d", role.Id))
	}

	roleModel := authmodel.TAuthRole{
		Name:       req.Name,
		Desc:       req.Desc,
		Sort:       int32(req.Sort),
		Status:     0,
		CreateTime: int32(time.Now().Unix()),
	}
	return r.roleDao.CreateRole(c, &roleModel, req.NodeId)
}

// RoleList 分页获取角色
func (r *RoleService) RoleList(c *gin.Context, req *request.SearchRoleListParam) (err error, list *[]response.AuthRoleBaseRes, total int64) {
	return r.roleDao.RoleList(c, req)
}

// RoleAll 所有角色
func (r *RoleService) RoleAll(c *gin.Context) (err error, list *[]response.AuthRoleBaseRes, total int64) {
	return r.roleDao.RoleAll(c)
}

// RoleDelete 删除角色
func (r *RoleService) RoleDelete(c *gin.Context, req *commonreq.PrimaryIdParam) error {
	role, err := r.roleDao.GetRole(c, req.Id, "")
	if err != nil {
		return err
	}
	if role != nil && role.Id != req.Id {
		return errors.New(fmt.Sprintf("RoleDelete获取角色失败,paramRoleId=%d,resultRoleId=%d", req.Id, role.Id))
	}
	role.DeleteTime = int32(time.Now().Unix())
	rowsAffected, err := r.roleDao.RoleDelete(c, role)
	if err != nil {
		return err
	}
	if rowsAffected != 1 {
		return errors.New(fmt.Sprintf("删除角色异常,expectRowsAffected=1,actual=%d", rowsAffected))
	}
	return nil
}

// RoleUpdate 角色编辑
// 编辑角色的时候,将节点id和节点的父级存起来
func (r *RoleService) RoleUpdate(c *gin.Context, req *request.RoleUpdateParam) error {
	if utils.ContainsDuplicate(req.NodeId) {
		return errors.New("节点id不能重复")
	}

	if req.Id < 1 {
		return errors.New("角色id不能为空")
	}

	// 角色不允许重复
	count, rowsAffected, err := r.roleDao.RoleIsExist(c, req.Id, req.Name)
	if err != nil {
		return err
	}

	if count > 0 && rowsAffected > 0 {
		return errors.New("角色名称已经存在")
	}

	return r.roleDao.RoleUpdate(c, req)
}

// RoleDetail 角色详情
func (r *RoleService) RoleDetail(c *gin.Context, req *commonreq.PrimaryIdParam) (*response.RoleDetail, error) {
	var res response.RoleDetail
	var roleNodeList []*int32

	// 获取角色的节点
	list, err := r.roleDao.RoleNode(c, req.Id)
	if err != nil {
		return nil, err
	}
	for _, node := range list {
		roleNodeList = append(roleNodeList, &node.NodeId)
	}
	res.NodeId = roleNodeList

	// 获取角色基本信息
	roleDetail, err := r.roleDao.GetRole(c, req.Id, "")
	if err != nil {
		return nil, err
	}
	if roleDetail == nil || roleDetail.Id < 1 {
		return &res, nil
	}
	res.Info = &response.AuthRoleBaseRes{
		Id:   roleDetail.Id,
		Name: roleDetail.Name,
		Desc: roleDetail.Desc,
		Sort: int(roleDetail.Sort),
	}
	return &res, err
}
