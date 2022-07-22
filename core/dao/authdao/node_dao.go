package authdao

import (
	"equity/core/dao"
	"equity/core/model/authmodel"
	"equity/core/model/authmodel/request"
	"equity/core/model/authmodel/response"
	"equity/global"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type NodeDao struct {
	*dao.Dao
}

// CreateNode 新增节点
func (n *NodeDao) CreateNode(c *gin.Context, node *authmodel.TAuthNode) error {
	return global.DBEnging.Create(&node).Error
}

// Detail 节点详情
func (n *NodeDao) Detail(c *gin.Context, id int32) (*response.NodeDetail, error) {
	var res *response.NodeDetail
	err := global.DBEnging.Model(&authmodel.TAuthNode{}).
		Where("id = ?", id).
		Find(&res).Error
	return res, err
}

// UpdateNode  编辑节点
func (n *NodeDao) UpdateNode(c *gin.Context, req *request.AuthNodeUpdateParam) error {
	return global.DBEnging.Transaction(func(tx *gorm.DB) error {
		updatesMap := map[string]interface{}{
			"pid":         req.Pid,
			"name":        req.Name,
			"icon":        req.Icon,
			"url":         req.Url,
			"type":        req.Type,
			"sort":        req.Sort,
			"update_time": int32(time.Now().Unix()),
		}

		// 更新用户
		txUpdate := tx.Model(&authmodel.TAuthNode{}).Where("id = ?", req.Id).Updates(updatesMap)
		rowsAffected := txUpdate.RowsAffected
		if err := txUpdate.Error; err != nil {
			return err
		}
		if rowsAffected != 1 {
			return errors.New(fmt.Sprintf("编辑节点失败,rowsAffected = %d", rowsAffected))
		}
		return nil
	})
}

// ParentNodeIsExist 父节点是否存在
func (n *NodeDao) ParentNodeIsExist(c *gin.Context, pid int32) (count int64, err error) {
	db := global.DBEnging.Model(&authmodel.TAuthNode{}).Where("id = ?", pid)
	err = db.Count(&count).Error
	return count, err
}

// GetNodeIdsByRoleIds 通过角色id获取节点id
func (n *NodeDao) GetNodeIdsByRoleIds(c *gin.Context, roleIds []*int32) ([]*int32, error) {
	var nodeIds []*int32
	err := global.DBEnging.Model(&authmodel.TAuthRoleNode{}).
		Select("node_id").
		Where("role_id in ?", roleIds).
		Find(&nodeIds).Error
	return nodeIds, err
}

// GetNodeUrls 通过节点id获取节点url信息
func (n *NodeDao) GetNodeUrls(c *gin.Context, nodeIds []*int32) ([]*string, error) {
	var nodeUrls []*string
	err := global.DBEnging.Model(&authmodel.TAuthNode{}).
		Select("url").
		Where("id in ?", nodeIds).
		Find(&nodeUrls).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return nodeUrls, err
}

// NodeAll 获取所有节点
func (n *NodeDao) NodeAll(c *gin.Context, ids []*int32) ([]*authmodel.TAuthNode, error) {
	var nodeList []*authmodel.TAuthNode
	db := global.DBEnging.Model(&authmodel.TAuthNode{})
	if len(ids) > 0 {
		db = db.Where("id in ?", ids)
	}
	err := db.Order("sort desc").Scan(&nodeList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return nodeList, err
}
