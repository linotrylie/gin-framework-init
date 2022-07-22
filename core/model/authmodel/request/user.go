package request

import (
	"equity/core/consts"
	"equity/core/model/authmodel"
	validation "github.com/go-ozzo/ozzo-validation"
)

// UserLoginParam  管理员登录参数
type UserLoginParam struct {
	UserName string // 用户名
	Password string // 密码
}

func (param *UserLoginParam) Validate() error {
	return validation.ValidateStruct(param,
		validation.Field(&param.UserName, validation.Required.Error(consts.ValidateColumnRequired)),
		validation.Field(&param.Password,
			validation.Required.Error(consts.ValidateColumnRequired),
			validation.Length(6, 32).Error("密码长度在6到32之间"),
		),
	)
}

// CreateUserParam  新增管理员参数
type CreateUserParam struct {
	UserName string // 用户名
	Account  string // 账号
	Password string // 密码
	RoleIds  []int  // 角色id数组
}

func (param *CreateUserParam) Validate() error {
	return validation.ValidateStruct(param,
		validation.Field(&param.Account, validation.Required.Error(consts.ValidateColumnRequired)),
		validation.Field(&param.UserName, validation.Required.Error(consts.ValidateColumnRequired)),
		validation.Field(&param.Password,
			validation.Required.Error(consts.ValidateColumnRequired),
			validation.Length(6, 32).Error("密码长度在6到32之间"),
		),
	)
}

// UpdateUserParam  编辑管理员参数
type UpdateUserParam struct {
	Id       int32  // 用户id
	UserName string // 用户名
	RoleIds  []int  // 角色id数组
}

func (param *UpdateUserParam) Validate() error {
	return validation.ValidateStruct(param,
		validation.Field(&param.Id,
			validation.Required.Error(consts.ValidateColumnRequired),
			validation.Min(1).Error(consts.MinIsOne),
		),
		validation.Field(&param.UserName, validation.Required.Error(consts.ValidateColumnRequired)),
	)
}

// SetUserStatusParam  设置管理员状态参数
type SetUserStatusParam struct {
	UserId int32 // 用户id
	Status int   // 用户状态
}

func (param *SetUserStatusParam) Validate() error {
	return validation.ValidateStruct(param,
		validation.Field(&param.UserId,
			validation.Required.Error(consts.ValidateColumnRequired),
			validation.Min(1).Error(consts.MinIsOne),
		),
		validation.Field(&param.Status, validation.In(authmodel.AuthUserDisableTrue, authmodel.AuthUserDisableFalse).Error(consts.DataNotInAllowedRange)),
	)
}

// ChangeUserPWDParam 修改用户密码
type ChangeUserPWDParam struct {
	UserId          int32  // 用户id
	Password        string // 密码
	ConfirmPassword string // 确认密码
}

func (param *ChangeUserPWDParam) Validate() error {
	return validation.ValidateStruct(param,
		validation.Field(&param.UserId,
			validation.Required.Error(consts.ValidateColumnRequired),
			validation.Min(1).Error(consts.MinIsOne),
		),
		validation.Field(&param.Password,
			validation.Required.Error(consts.ValidateColumnRequired),
			validation.Length(6, 32).Error("密码长度在6到32之间"),
		),
		validation.Field(&param.ConfirmPassword,
			validation.Required.Error(consts.ValidateColumnRequired),
			validation.Length(6, 32).Error("确认密码长度在6到32之间"),
		),
	)
}

// SearchUserListParam 搜索用户列表
type SearchUserListParam struct {
	Name     string
	Account  string
	PageNo   int
	PageSize int
}

func (param *SearchUserListParam) Validate() error {
	return nil
}
