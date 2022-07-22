package response

// SystemTraceInfo 基本信息
type SystemTraceInfo struct {
	// PK
	Id int64 `json:"id"`
	// 入参
	ReqParam string `json:"reqParam"`
	// 返参
	ResParam int32 `json:"resParam"`
	// 创建时间
	CreateTime int32 `json:"createTime"`
	// 错误信息
	ErrMsg string `json:"errMsg"`
	// 商户id
	MchId int32 `json:"mchId"`
	// 请求记录id
	Rid string `json:"rid"`
	// 调用路径
	FullPath string `json:"fullPath"`
}
