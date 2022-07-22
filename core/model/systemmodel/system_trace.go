// Package systemmodel auto generate file, do not edit it!
package systemmodel

type TSystemTrace struct {
	// PK
	Id int64 `json:"id"`
	// 入参
	ReqParam string `json:"req_param"`
	// 返参
	ResParam string `json:"res_param"`
	// 创建时间
	CreateTime int32 `json:"create_time"`
	// 错误信息
	ErrMsg string `json:"err_msg"`
	// 商户id
	MchId int32 `json:"mch_id"`
	// 请求记录id
	Rid string `json:"rid"`
	// 调用路径
	FullPath string `json:"full_path"`
}

// TableName TSystemTrace
func (model TSystemTrace) TableName() string {
	return "t_system_trace"
}
