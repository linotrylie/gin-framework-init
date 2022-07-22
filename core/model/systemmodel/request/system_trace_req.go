package request

import (
	"equity/core/consts"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/pkg/errors"
)

// SystemTraceCreate 创建
type SystemTraceCreate struct {
	// 入参
	ReqParam string `json:"reqParam"`
	// 返参
	ResParam string `json:"resParam"`
	// 错误信息
	ErrMsg string `json:"errMsg"`
	// 请求记录id
	Rid string `json:"rid"`
	// 调用路径
	FullPath string `json:"fullPath"`

	AppKey string `json:"appKey"`
}

func (param *SystemTraceCreate) Validate() error {
	if param == nil {
		return errors.New(consts.ParamRequired)
	}
	return validation.ValidateStruct(param,
		validation.Field(&param.ReqParam,
			validation.Length(0, 65535).Error("长度在0到65535之间"),
		),
		validation.Field(&param.ResParam,
			validation.Min(0).Error(consts.MinIsZero),
			validation.Max(consts.MaxInt32Value).Error(consts.CannotGreatThanInt32MaxValue),
		),
		validation.Field(&param.ErrMsg,
			validation.Length(0, 65535).Error("长度在0到65535之间"),
		),
		validation.Field(&param.Rid,
			validation.Length(0, 32).Error("长度在0到32之间"),
		),
		validation.Field(&param.FullPath,
			validation.Length(0, 250).Error("长度在0到250之间"),
		),
	)

}

// SystemTraceUpdate 编辑
type SystemTraceUpdate struct {
	// PK
	Id int64 `json:"id"`
	// 入参
	ReqParam string `json:"reqParam"`
	// 返参
	ResParam int32 `json:"resParam"`
	// 错误信息
	ErrMsg string `json:"errMsg"`
	// 商户id
	MchId int32 `json:"mchId"`
	// 请求记录id
	Rid string `json:"rid"`
	// 调用路径
	FullPath string `json:"fullPath"`
}

func (param *SystemTraceUpdate) Validate() error {
	if param == nil {
		return errors.New(consts.ParamRequired)
	}
	return validation.ValidateStruct(param,
		validation.Field(&param.Id,
			validation.Min(0).Error(consts.MinIsZero),
			validation.Required.Error(consts.ValidateColumnRequired),
			validation.Max(consts.MaxInt64Value).Error(consts.CannotGreatThanInt64MaxValue),
		),
		validation.Field(&param.ReqParam,
			validation.Length(0, 65535).Error("长度在0到65535之间"),
		),
		validation.Field(&param.ResParam,
			validation.Min(0).Error(consts.MinIsZero),
			validation.Max(consts.MaxInt32Value).Error(consts.CannotGreatThanInt32MaxValue),
		),
		validation.Field(&param.ErrMsg,
			validation.Length(0, 65535).Error("长度在0到65535之间"),
		),
		validation.Field(&param.MchId,
			validation.Min(0).Error(consts.MinIsZero),
			validation.Max(consts.MaxInt32Value).Error(consts.CannotGreatThanInt32MaxValue),
		),
		validation.Field(&param.Rid,
			validation.Length(0, 32).Error("长度在0到32之间"),
		),
		validation.Field(&param.FullPath,
			validation.Length(0, 250).Error("长度在0到250之间"),
		),
	)

}

// SystemTracePageSearch 分页获取
type SystemTracePageSearch struct {
	Rid      string
	MchId    int32
	PageNo   int
	PageSize int
}

func (param *SystemTracePageSearch) Validate() error {
	if param == nil {
		return errors.New(consts.ParamRequired)
	}
	return validation.ValidateStruct(param,
		validation.Field(&param.PageNo,
			validation.Min(1).Error(consts.MinIsOne),
		),
		validation.Field(&param.PageSize,
			validation.Min(1).Error(consts.MinIsOne),
			validation.Max(20).Error(consts.MaxPageSizeIsTwenty),
		),
	)
}
