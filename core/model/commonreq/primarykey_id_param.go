package commonreq

import (
	"equity/core/consts"
	validation "github.com/go-ozzo/ozzo-validation"
)

// PrimaryIdParam Find by id structure
type PrimaryIdParam struct {
	Id int32 `json:"id" form:"id"`
}
type PrimaryIdParams struct {
	Ids []int32 `json:"id" form:"id"`
}

func (param *PrimaryIdParam) Validate() error {
	return validation.ValidateStruct(param,
		validation.Field(&param.Id,
			validation.Required.Error(consts.ValidateColumnRequired),
			validation.Min(1).Error(consts.MinIsOne),
			validation.Max(consts.MaxInt32Value).Error(consts.CannotGreatThanInt32MaxValue),
		),
	)
}

func (param *PrimaryIdParams) Validate() error {
	return nil
}
