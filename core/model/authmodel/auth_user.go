// Package authmodel auto generate file, do not edit it!
package authmodel

type TAuthUser struct {
	// 自增ID
	Id int32 `json:"id"`
	// 是否超级管理员:1=是,2=否
	IsRoot int8 `json:"is_root"`
	// 名称
	Name string `json:"name"`
	// 账号
	Account string `json:"account"`
	// 密码
	Password string `json:"password"`
	// 密码盐
	Salt string `json:"salt"`
	// 创建时间
	CreateTime int32 `json:"create_time"`
	// 修改时间
	UpdateTime int32 `json:"update_time"`
	// 最后登录时间
	LoginTime int32 `json:"login_time"`
	// 最后登录ip
	LoginIp string `json:"login_ip"`
	// 是否禁用:1=是,2=否
	Disable int8 `json:"disable"`
	// 0为非删除状态,非0位删除时间
	DeleteTime int32 `json:"delete_time"`
	// 登录详细地址
	LoginAddress string `json:"login_address"`
	// 是否删除:1=是,2=否
	IsDelete int8 `json:"is_delete"`
}

// TableName TAuthUser
func (model TAuthUser) TableName() string {
	return "t_auth_user"
}

// 是否超级管理员:1=是,2=否
const AuthUserIsRootTrue int8 = 1
const AuthUserIsRootFalse int8 = 2

// 是否禁用:1=是,2=否
const AuthUserDisableTrue int8 = 1
const AuthUserDisableFalse int8 = 2

// 是否删除:1=是,2=否
const AuthUserIsDeleteTrue int8 = 1
const AuthUserIsDeleteFalse int8 = 2
