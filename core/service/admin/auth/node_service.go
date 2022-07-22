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

// NodeService 节点相关服务
type NodeService struct {
	nodeDao *authdao.NodeDao
}

var userService UserService

// CreateNode 新增节点
func (n *NodeService) CreateNode(c *gin.Context, req *request.AuthNodeCreateParam) error {
	// 有父节点的,判断父节点必须存在
	if req.Pid > 0 {
		count, err := n.nodeDao.ParentNodeIsExist(c, req.Pid)
		if err != nil {
			return err
		}
		if count != 1 {
			return errors.New("父节点不存在")
		}
	}
	nodeModel := authmodel.TAuthNode{
		Pid:        req.Pid,
		Name:       req.Name,
		Icon:       req.Icon,
		Url:        req.Url,
		Type:       int8(req.Type),
		Sort:       req.Sort,
		CreateTime: int32(time.Now().Unix()),
		UpdateTime: int32(time.Now().Unix()),
	}
	return n.nodeDao.CreateNode(c, &nodeModel)
}

// UpdateNode 编辑节点
func (n *NodeService) UpdateNode(c *gin.Context, req *request.AuthNodeUpdateParam) error {
	if req.Pid > 0 {
		count, err := n.nodeDao.ParentNodeIsExist(c, int32(req.Pid))
		if err != nil {
			return err
		}
		if count != 1 {
			return errors.New(fmt.Sprintf("父节点不存在,pid=%d", req.Pid))
		}
	}
	return n.nodeDao.UpdateNode(c, req)
}

// NodeDetail 节点详情
func (n *NodeService) NodeDetail(c *gin.Context, req *commonreq.PrimaryIdParam) (*response.NodeDetail, error) {
	return n.nodeDao.Detail(c, req.Id)
}

// UserNodeTree 获取用户的节点树
func (n *NodeService) UserNodeTree(c *gin.Context) (*response.UserNodeTree, error) {
	// 节点列表
	var nodeList []*authmodel.TAuthNode

	// 通过token获取用户id
	uid, err := utils.GetUidByToken(c)
	if err != nil {
		return nil, err
	}

	// 判断用户是否存在
	user, err := userService.userDao.GetUserById(c, int32(uid))
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("用户获取失败")
	}

	// 管理员拥有所有节点树
	// 区分超级管理员和普通管理员
	nodeList, err = n.userNode(c, user.IsRoot, int32(int(user.Id)))
	if err != nil {
		return nil, err
	}

	if len(nodeList) < 1 {
		return nil, nil
	}

	nodeTrees := nodeTreeFormat(nodeList, true)
	res := &response.UserNodeTree{
		// 用户拥有的目录,菜单,权限
		NodeTree: getTreeRecursive(nodeTrees, 0),
		// 用户拥有的button权限
		ButtonUrl: nodeButtons(nodeList),
	}
	return res, nil
}

// userNode 用户拥有的节点,
// 区分超级用户和普通用户
func (n *NodeService) userNode(c *gin.Context, isRoot int8, uid int32) ([]*authmodel.TAuthNode, error) {
	// 超级管理员
	if isRoot == authmodel.AuthUserIsRootTrue {
		return n.nodeDao.NodeAll(c, nil)
	}

	// 普通用户通过角色获取
	// 获取用户角色id
	userDetailParam := commonreq.PrimaryIdParam{Id: uid}
	userDetail, err := userService.UserDetail(c, &userDetailParam)
	if err != nil {
		return nil, err
	}

	// 角色id
	roleIds := userDetail.Role

	// 没有绑定角色
	if len(roleIds) < 1 {
		return nil, nil
	}

	// 通过角色获取节点id
	nodeIds, err := n.nodeDao.GetNodeIdsByRoleIds(c, roleIds)
	if err != nil {
		return nil, err
	}

	// 没有节点id
	if len(nodeIds) < 1 {
		return nil, nil
	}

	// 用户拥有的节点
	return n.nodeDao.NodeAll(c, nodeIds)
}

// NodeTree 节点树
func (n *NodeService) NodeTree(c *gin.Context) ([]*response.NodeTree, error) {
	nodeList, err := n.nodeDao.NodeAll(c, nil)
	if err != nil {
		return nil, err
	}
	nodeTrees := nodeTreeFormat(nodeList, false)
	return getTreeRecursive(nodeTrees, 0), nil
}

// getTreeRecursive 递归获取分类树
func getTreeRecursive(list []*response.NodeTree, parentId int32) []*response.NodeTree {
	res := make([]*response.NodeTree, 0)
	for _, v := range list {
		if v.Pid == parentId {
			v.Children = getTreeRecursive(list, v.Id)
			res = append(res, v)
		}
	}
	return res
}

// nodeButtons 获取button类型的节点
func nodeButtons(nodeList []*authmodel.TAuthNode) []*string {
	nodeButtons := make([]*string, 0)
	for _, node := range nodeList {
		if node.Type == authmodel.AuthNodeTypeButton {
			nodeButtons = append(nodeButtons, &node.Url)
		}
	}
	return nodeButtons
}

// nodeTreeFormat 节点格式转换
func nodeTreeFormat(nodeList []*authmodel.TAuthNode, isIgnoreButton bool) []*response.NodeTree {
	nodeTrees := make([]*response.NodeTree, 0)
	for _, node := range nodeList {
		// button 是否参与构建资源树
		if isIgnoreButton == true && node.Type == authmodel.AuthNodeTypeButton {
			continue
		}
		nodeTree := &response.NodeTree{
			Id:         node.Id,
			Pid:        node.Pid,
			Name:       node.Name,
			Icon:       node.Icon,
			Url:        node.Url,
			Type:       node.Type,
			Sort:       node.Sort,
			CreateTime: node.CreateTime,
			UpdateTime: node.UpdateTime,
			Children:   nil,
		}
		nodeTrees = append(nodeTrees, nodeTree)
	}
	return nodeTrees
}
