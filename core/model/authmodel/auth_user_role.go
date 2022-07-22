// Package authmodel auto generate file, do not edit it!
package authmodel

type TAuthUserRole struct {
	// 用户id
	UserId int32 `json:"user_id"`
	// 角色id
	RoleId int32 `json:"role_id"`
}

// TableName TAuthUserRole
func (model TAuthUserRole) TableName() string {
	return "t_auth_user_role"
}
