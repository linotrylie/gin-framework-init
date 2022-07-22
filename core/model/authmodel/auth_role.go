// Package authmodel auto generate file, do not edit it!
package authmodel

type TAuthRole struct {
	// 自增ID
	Id int32 `json:"id"`
	// 角色名称
	Name string `json:"name"`
	// 描述
	Desc string `json:"desc"`
	// 排序
	Sort int32 `json:"sort"`
	// 显隐:1=显示,2=隐藏
	Status int8 `json:"status"`
	// 创建时间
	CreateTime int32 `json:"create_time"`
	// 更新时间
	UpdateTime int32 `json:"update_time"`
	// 是否删除:1=是,2=否
	IsDelete int8 `json:"is_delete"`
	// 删除时间
	DeleteTime int32 `json:"delete_time"`
}

// TableName TAuthRole
func (model TAuthRole) TableName() string {
	return "t_auth_role"
}

// 显隐:1=显示,2=隐藏
const AuthRoleStatusShow int8 = 1
const AuthRoleStatusHide int8 = 2

// 是否删除:1=是,2=否
const AuthRoleIsDeleteTrue int8 = 1
const AuthRoleIsDeleteFalse int8 = 2
