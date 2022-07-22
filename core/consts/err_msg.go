package consts

// 数据验证错误信息
const (
	ParamRequired                = "请求参数不能为空"
	ValidateColumnRequired       = "必填字段"
	MinIsOne                     = "最小为1"
	MinIsZero                    = "最小为0"
	MaxIsTwo                     = "最大为2"
	DataNotInAllowedRange        = "不在允许的范围"
	CannotGreatThanInt8MaxValue  = "最大值不能超过127"
	CannotGreatThanInt16MaxValue = "最大值不能超过32767"
	CannotGreatThanInt32MaxValue = "最大值不能超过2147483647"
	CannotGreatThanInt64MaxValue = "最大值不能超过9223372036854775807"
	PhoneFormatError             = "手机号格式错误"
	EmailFormatError             = "邮箱号格式错误"
	DateYYMMDDFormatError        = "日期格式错误,正确格式:yyyy-MM-dd"
	MaxPageSizeIsTwenty          = "分页最大条数为20"
	ReadBodyError                = "读取body参数失败"
)
