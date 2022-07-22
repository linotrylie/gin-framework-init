// Package authmodel auto generate file, do not edit it!
package authmodel

type TAuthNode struct {
	// 自增ID
	Id int32 `json:"id"`
	// 父级ID
	Pid int32 `json:"pid"`
	// 名称
	Name string `json:"name"`
	// 图标
	Icon string `json:"icon"`
	// 路由规则
	Url string `json:"url"`
	// 类型:1=目录,2=菜单,3=按钮
	Type int8 `json:"type"`
	// 排序
	Sort int32 `json:"sort"`
	// 创建时间
	CreateTime int32 `json:"create_time"`
	// 更新时间
	UpdateTime int32 `json:"update_time"`
}

// TableName TAuthNode
func (model TAuthNode) TableName() string {
	return "t_auth_node"
}

// 类型:1=目录,2=菜单,3=按钮
const AuthNodeTypeMenu int8 = 1
const AuthNodeTypeDir int8 = 2
const AuthNodeTypeButton int8 = 3
