package authdao

import (
	"equity/core/dao"
	"equity/core/model/authmodel"
	"equity/core/model/authmodel/request"
	"equity/core/model/authmodel/response"
	"equity/core/model/commonres"
	"equity/global"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type RoleDao struct {
	*dao.Dao
}

// CreateRole 新增角色
func (r *RoleDao) CreateRole(c *gin.Context, role *authmodel.TAuthRole, nodeIds []int) error {
	return global.DBEnging.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&role).Error
		if err != nil {
			return err
		}

		if role.Id < 1 {
			return errors.New("新增角色,获取角色id失败")
		}

		// 一个用户可能有多个角色
		if len(nodeIds) > 0 {
			var roleNodes []authmodel.TAuthRoleNode
			for _, nodeId := range nodeIds {
				roleNodeItem := authmodel.TAuthRoleNode{
					RoleId: role.Id,
					NodeId: int32(nodeId),
				}
				roleNodes = append(roleNodes, roleNodeItem)
			}

			// 批量新增新的角色节点关系
			err := tx.Create(&roleNodes).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}

// GetRole 查角色
func (r *RoleDao) GetRole(c *gin.Context, id int32, name string) (tAuthRole *authmodel.TAuthRole, err error) {
	var roleModel authmodel.TAuthRole
	db := global.DBEnging.Where(" is_delete <> ? ", authmodel.AuthRoleIsDeleteTrue)
	if id > 0 {
		db = db.Where("id = ?", id)
	}
	if name != "" {
		db = db.Where("name = ?", name)
	}
	err = db.Find(&roleModel).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return &roleModel, err
	}
	return &roleModel, nil
}

// GetRoleByUserId 获取用户所属的角色id
func (r *RoleDao) GetRoleByUserId(c *gin.Context, userId int32) (tAuthRole []*authmodel.TAuthUserRole, err error) {
	var userRoleModel []*authmodel.TAuthUserRole
	db := global.DBEnging.Where(" user_id =  ? ", userId)
	err = db.Find(&userRoleModel).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return userRoleModel, err
	}
	return userRoleModel, nil
}

// RoleIsExist 角色名称是否唯一
func (r *RoleDao) RoleIsExist(c *gin.Context, id int, name string) (count, rowsAffected int64, err error) {
	db := global.DBEnging.Model(&authmodel.TAuthRole{}).Where(" is_delete <> ? ", authmodel.AuthRoleIsDeleteTrue)
	if id > 0 {
		// 是否其他角色也和当前角色名称相同
		db = db.Where(" id <> ?", id)
	}
	if name != "" {
		db = db.Where(" name = ?", name)
	}
	err = db.Count(&count).Error
	return count, db.RowsAffected, err
}

// RoleList 分页查询角色列表
func (r *RoleDao) RoleList(c *gin.Context, req *request.SearchRoleListParam) (err error, list *[]response.AuthRoleBaseRes, total int64) {
	limit, offset := commonres.GetPager(req.PageNo, req.PageSize)
	var roleList []response.AuthRoleBaseRes
	var roleModel authmodel.TAuthRole

	// 创建db
	db := global.DBEnging.Model(&roleModel).Where("is_delete = ? ", authmodel.AuthRoleIsDeleteFalse)
	if req.Name != "" {
		db = db.Where("name LIKE ?", commonres.LikeSearch(req.Name))
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Order("id DESC").Limit(limit).Offset(offset).Scan(&roleList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, &roleList, total
	}
	return err, &roleList, total
}

// RoleAll 查询所有角色列表
func (r *RoleDao) RoleAll(c *gin.Context) (err error, list *[]response.AuthRoleBaseRes, total int64) {
	var roleList []response.AuthRoleBaseRes
	var roleModel authmodel.TAuthRole

	// 创建db
	db := global.DBEnging.Model(&roleModel).Where("is_delete = ? ", authmodel.AuthRoleIsDeleteFalse)
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Order("id DESC").Scan(&roleList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, &roleList, total
	}
	return err, &roleList, total
}

// RoleDelete 设置删除状态
func (r *RoleDao) RoleDelete(c *gin.Context, role *authmodel.TAuthRole) (rowsAffected int64, err error) {
	db := global.DBEnging.Model(&role).
		Updates(map[string]interface{}{"delete_time": role.DeleteTime, "is_delete": authmodel.AuthRoleIsDeleteTrue})
	return db.RowsAffected, db.Error
}

// RoleNode 角色对应的节点
func (r *RoleDao) RoleNode(c *gin.Context, roleId int32) (roleNode []*authmodel.TAuthRoleNode, err error) {
	var roleNodeModel []*authmodel.TAuthRoleNode
	db := global.DBEnging.Where(" role_id =  ? ", roleId)
	err = db.Find(&roleNodeModel).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return roleNodeModel, err
	}
	return roleNodeModel, nil
}

// RoleUpdate 角色编辑
func (r *RoleDao) RoleUpdate(c *gin.Context, req *request.RoleUpdateParam) error {
	var roleModel authmodel.TAuthRole
	return global.DBEnging.Transaction(func(tx *gorm.DB) error {
		// 编辑角色
		roleModelUpdatesMap := map[string]interface{}{
			"update_time": int32(time.Now().Unix()),
			"name":        req.Name,
			"desc":        req.Desc,
			"sort":        req.Sort,
		}

		txUpdate := tx.Model(&roleModel).Where("id = ?", req.Id).Updates(roleModelUpdatesMap)
		rowsAffected := txUpdate.RowsAffected
		if err := txUpdate.Error; err != nil {
			return err
		}
		if rowsAffected != 1 {
			return errors.New("编辑角色失败")
		}

		// 批量删除角色和节点的关系
		if err := tx.Unscoped().Where("role_id = ?", req.Id).Delete(&authmodel.TAuthRoleNode{}).Error; err != nil && err != gorm.ErrRecordNotFound {
			return err
		}

		// 一个角色对应多个节点
		if len(req.NodeId) > 0 {
			var batchCreateRoleNode []authmodel.TAuthRoleNode
			for _, nodeId := range req.NodeId {
				item := authmodel.TAuthRoleNode{
					RoleId: int32(req.Id),
					NodeId: int32(nodeId),
				}
				batchCreateRoleNode = append(batchCreateRoleNode, item)
			}

			// 批量新增新的角色节点关系
			err := tx.Create(&batchCreateRoleNode).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}
