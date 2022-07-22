package errcode

var (
	Success                            = NewError(0, "成功")
	ServerError                        = NewError(10000000, "服务器内部错误")
	InvalidParams                      = NewError(10000001, "Token 不能为空")
	NotFond                            = NewError(10000002, "找不到")
	UnauthorizedAuthNotExist           = NewError(10000003, "鉴权失败,找不到对应的AppKey 和 AppSecret")
	UnauthorizedAuthTokenError         = NewError(10000004, "鉴权失败,Token 错误")
	UnauthorizedAuthTokenTimeout       = NewError(10000005, "鉴权失败,Token 已过期,请重新登录")
	UnauthorizedAuthTokenGenerateError = NewError(10000006, "鉴权失败,Token 生成失败")
	TooManyRequests                    = NewError(10000007, "请求过多")
)
