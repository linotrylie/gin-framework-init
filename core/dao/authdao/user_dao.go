package authdao

import (
	"equity/core/dao"
	"equity/core/model/authmodel"
	"equity/core/model/authmodel/request"
	"equity/core/model/authmodel/response"
	"equity/core/model/commonres"
	"equity/global"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type UserDao struct {
	*dao.Dao
}

// GetUser 登录查用户
func (u *UserDao) GetUser(c *gin.Context, req *request.UserLoginParam) (tAuthUser *authmodel.TAuthUser, err error) {
	var userModel authmodel.TAuthUser
	err = global.DBEnging.Where("account = ?", req.UserName).Find(&userModel).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return &userModel, err
	}
	return &userModel, nil
}

// GetUserCountByName 通过用户名查用户数
func (u *UserDao) GetUserCountByName(c *gin.Context, name string) (int64, error) {
	var count int64
	db := global.DBEnging.Model(authmodel.TAuthUser{}).Where("name = ? AND is_delete <> ?", name, authmodel.AuthUserIsDeleteTrue)
	db.Count(&count)
	err := db.Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, err
	}
	return count, nil
}

// UserNameIsExist 用户名是否存在
func (u *UserDao) UserNameIsExist(c *gin.Context, name string, userId int32) (int64, error) {
	var count int64
	db := global.DBEnging.Model(authmodel.TAuthUser{}).Where(
		"name = ? AND is_delete <> ? AND id <> ?",
		name,
		authmodel.AuthUserIsDeleteTrue,
		userId)
	db.Count(&count)
	err := db.Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, err
	}
	return count, nil
}

// GetUserByAccount 通过账号查用户
func (u *UserDao) GetUserByAccount(c *gin.Context, account string) (count int64, err error) {
	db := global.DBEnging.Model(authmodel.TAuthUser{}).Where(
		"account = ? AND is_delete <> ?",
		account,
		authmodel.AuthUserIsDeleteTrue,
	)
	db.Count(&count)
	err = db.Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, err
	}
	return count, err
}

// GetUserById 通过id查用户
func (u *UserDao) GetUserById(c *gin.Context, id int32) (tAuthUser *authmodel.TAuthUser, err error) {
	var userModel authmodel.TAuthUser
	err = global.DBEnging.Where("id = ? AND is_delete <> ? ", id, authmodel.AuthUserIsDeleteTrue).Find(&userModel).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	if userModel.Id < 1 {
		return nil, nil
	}
	return &userModel, nil
}

// CreateUser 新增用户
func (u *UserDao) CreateUser(c *gin.Context, user *authmodel.TAuthUser, roleIds []int) error {
	return global.DBEnging.Transaction(func(tx *gorm.DB) error {
		// 创建用户
		txCreate := tx.Create(&user)
		err := txCreate.Error
		if err != nil {
			return err
		}

		if txCreate.RowsAffected != 1 {
			return errors.New(fmt.Sprintf("创建用户失败,期望成功条数=1,实际条数=%d", txCreate.RowsAffected))
		}

		// 一个用户可能有多个角色
		if len(roleIds) > 0 {
			var userRoles []authmodel.TAuthUserRole
			for _, roleId := range roleIds {
				authUserRole := authmodel.TAuthUserRole{
					RoleId: int32(roleId),
					UserId: user.Id,
				}
				userRoles = append(userRoles, authUserRole)
			}

			// 新增新的用户和角色关系
			err := tx.Create(&userRoles).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}

// UpdateUser  编辑用户
func (u *UserDao) UpdateUser(c *gin.Context, req request.UpdateUserParam) error {
	return global.DBEnging.Transaction(func(tx *gorm.DB) error {
		var userModel authmodel.TAuthUser
		userModelUpdatesMap := map[string]interface{}{
			"update_time": int32(time.Now().Unix()),
			"name":        req.UserName,
		}

		// 更新用户
		txUpdate := tx.Model(&userModel).Where("id = ?", req.Id).Updates(userModelUpdatesMap)
		rowsAffected := txUpdate.RowsAffected
		if err := txUpdate.Error; err != nil {
			return err
		}

		if rowsAffected != 1 {
			return errors.New("更新失败")
		}

		// 删除原来的角色id
		if err := tx.Unscoped().Where("user_id = ?", req.Id).Delete(&authmodel.TAuthUserRole{}).Error; err != nil {
			return err
		}

		// 一个用户可以对应有多个角色
		if len(req.RoleIds) > 0 {
			var userRoles []authmodel.TAuthUserRole
			// 新增新的用户和角色关系
			for _, roleId := range req.RoleIds {
				authUserRole := authmodel.TAuthUserRole{
					RoleId: int32(roleId),
					UserId: req.Id,
				}
				userRoles = append(userRoles, authUserRole)
			}
			err := tx.Create(&userRoles).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}

// SetUserStatus 设置用户状态
func (u *UserDao) SetUserStatus(c *gin.Context, user *authmodel.TAuthUser) (rowsAffected int64, err error) {
	db := global.DBEnging.Model(&user).
		Updates(map[string]interface{}{"update_time": user.UpdateTime, "disable": user.Disable})
	return db.RowsAffected, db.Error
}

// ChangePassword 修改密码
func (u *UserDao) ChangePassword(c *gin.Context, user *authmodel.TAuthUser) (rowsAffected int64, err error) {
	db := global.DBEnging.Model(&user).
		Updates(map[string]interface{}{
			"update_time": user.UpdateTime,
			"password":    user.Password,
			"salt":        user.Salt,
		})
	return db.RowsAffected, db.Error
}

// Delete 设置用户删除状态
func (u *UserDao) Delete(c *gin.Context, user *authmodel.TAuthUser) (rowsAffected int64, err error) {
	db := global.DBEnging.Model(&user).
		Updates(map[string]interface{}{"delete_time": user.DeleteTime, "is_delete": authmodel.AuthUserIsDeleteTrue})
	return db.RowsAffected, db.Error
}

// UserList 分页查询用户列表
func (u *UserDao) UserList(c *gin.Context, req *request.SearchUserListParam) (err error, list *[]response.AuthUserBaseRes, total int64) {
	var userModel authmodel.TAuthUser
	var userList []response.AuthUserBaseRes
	limit, offset := commonres.GetPager(req.PageNo, req.PageSize)

	// 创建db
	db := global.DBEnging.Model(&userModel).Where("is_delete = ?", authmodel.AuthUserIsDeleteFalse)
	if req.Name != "" {
		db = db.Where("name LIKE ?", commonres.LikeSearch(req.Name))
	}
	if req.Account != "" {
		db = db.Where("account LIKE ?", commonres.LikeSearch(req.Account))
	}
	db.Select([]string{"id,is_root,name,account,create_time,update_time,login_time,login_ip,disable,login_address"})
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Order("id desc").Limit(limit).Offset(offset).Scan(&userList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, &userList, total
	}
	return err, &userList, total
}
