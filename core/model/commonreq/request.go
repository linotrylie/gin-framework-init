package commonreq

// FrontRequest 商户请求
type FrontRequest struct {
	Header header `json:"header" form:"header" binding:"required"`
	Body   string `json:"body" form:"body" binding:"required"`
}

type header struct {
	RequestTime string `json:"requestTime" form:"requestTime" binding:"required"`
	AppKey      string `json:"appKey" form:"appKey" binding:"required"`
	Sign        string `json:"sign" form:"sign" binding:"required"`
	Method      string `json:"method" form:"method" binding:"required"`
	ResourceId  string `json:"resourceId" form:"resourceId" binding:"required"`
}
