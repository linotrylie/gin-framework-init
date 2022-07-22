// Package authmodel auto generate file, do not edit it!
package authmodel

type TAuthRoleNode struct {
	// 角色ID
	RoleId int32 `json:"role_id"`
	// 权限ID
	NodeId int32 `json:"node_id"`
}

// TableName TAuthRoleNode
func (model TAuthRoleNode) TableName() string {
	return "t_auth_role_node"
}
