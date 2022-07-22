package request

import (
	"equity/core/consts"
	validation "github.com/go-ozzo/ozzo-validation"
)

// RoleCreateParam 新增角色
type RoleCreateParam struct {
	Sort   int    // 排序
	Name   string // 角色名称
	Desc   string // 描述
	NodeId []int  // 节点id
}

func (param *RoleCreateParam) Validate() error {
	return validation.ValidateStruct(param,
		validation.Field(&param.Name, validation.Required.Error(consts.ValidateColumnRequired)),
	)
}

// SearchRoleListParam 搜索角色列表
type SearchRoleListParam struct {
	Name     string
	PageNo   int
	PageSize int
}

func (param *SearchRoleListParam) Validate() error {
	return nil
}

// RoleUpdateParam 角色编辑
type RoleUpdateParam struct {
	Id     int    // 角色id
	Name   string // 角色名称
	Desc   string // 角色描述
	NodeId []int  // 访问节点id列表
	Sort   int    // 排序
}

func (param *RoleUpdateParam) Validate() error {
	return validation.ValidateStruct(param,
		validation.Field(&param.Id,
			validation.Required.Error(consts.ParamRequired),
			validation.Min(1).Error(consts.MinIsOne),
		),
		validation.Field(&param.Name, validation.Required.Error(consts.ValidateColumnRequired)),
	)
}
