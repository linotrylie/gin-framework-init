package request

import (
	"equity/core/consts"
	"equity/core/model/authmodel"
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
)

// AuthNodeCreateParam 新增节点
type AuthNodeCreateParam struct {
	Pid  int32
	Name string
	Icon string
	Url  string
	Type int8
	Sort int32
}

func (param *AuthNodeCreateParam) Validate() error {
	if param == nil {
		return errors.New(consts.ParamRequired)
	}
	return validation.ValidateStruct(param,
		validation.Field(&param.Pid,
			validation.Min(0).Error(consts.MinIsZero),
			validation.Max(consts.MaxInt32Value).Error(consts.CannotGreatThanInt32MaxValue),
		),
		validation.Field(&param.Name, validation.Required.Error(consts.ValidateColumnRequired)),
		validation.Field(&param.Url, validation.Required.Error(consts.ValidateColumnRequired)),
		validation.Field(&param.Type,
			validation.Required.Error(consts.ValidateColumnRequired),
			validation.In(
				authmodel.AuthNodeTypeMenu,
				authmodel.AuthNodeTypeDir,
				authmodel.AuthNodeTypeButton,
			).Error(consts.DataNotInAllowedRange),
		),
		validation.Field(&param.Sort,
			validation.Required, validation.Min(1).Error(consts.MinIsOne),
		),
	)
}

// AuthNodeUpdateParam 编辑节点
type AuthNodeUpdateParam struct {
	Id   int
	Pid  int
	Name string
	Icon string
	Url  string
	Type int8
	Sort int32
}

func (param *AuthNodeUpdateParam) Validate() error {
	if param == nil {
		return errors.New(consts.ParamRequired)
	}
	return validation.ValidateStruct(param,
		validation.Field(&param.Pid,
			validation.Min(0).Error(consts.MinIsZero),
			validation.Max(consts.MaxInt32Value).Error(consts.CannotGreatThanInt32MaxValue),
		),
		validation.Field(&param.Id,
			validation.Min(1).Error(consts.MinIsOne),
			validation.Required.Error(consts.ValidateColumnRequired),
			validation.Max(consts.MaxInt32Value).Error(consts.CannotGreatThanInt32MaxValue),
		),
		validation.Field(&param.Name, validation.Required.Error(consts.ValidateColumnRequired)),
		validation.Field(&param.Url, validation.Required.Error(consts.ValidateColumnRequired)),
		validation.Field(&param.Type,
			validation.Required.Error(consts.ValidateColumnRequired),
			validation.In(
				authmodel.AuthNodeTypeMenu,
				authmodel.AuthNodeTypeDir,
				authmodel.AuthNodeTypeButton,
			).Error(consts.DataNotInAllowedRange),
		),
		validation.Field(&param.Sort,
			validation.Required.Error(consts.ValidateColumnRequired),
			validation.Min(1).Error(consts.MinIsOne),
			validation.Max(consts.MaxInt32Value).Error(consts.CannotGreatThanInt32MaxValue),
		),
	)
}
