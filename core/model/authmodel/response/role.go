package response

import "equity/utils/datetime"

// AuthRoleBaseRes 角色列表返回字段
type AuthRoleBaseRes struct {
	// 自增ID
	Id int32 `json:"id"`
	// 角色名称
	Name string `json:"name"`
	// 描述
	Desc string `json:"desc"`
	// 排序
	Sort int `json:"sort"`
	// 显隐:0-显示,1-隐藏
	Status int8 `json:"status"`
	// 创建时间
	CreateTime datetime.JSONTime `json:"createTime"`
	// 更新时间
	UpdateTime datetime.JSONTime `json:"updateTime"`
}

// RoleDetail 角色详情
type RoleDetail struct {
	Info   *AuthRoleBaseRes `json:"info"`   // 角色信息
	NodeId []*int32         `json:"nodeId"` // 节点信息
}
