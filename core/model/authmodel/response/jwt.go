package response

// AdminJwt 管理员jwt Payload 信息
type AdminJwtRes struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	Uid       int    `json:"uid"`
	Exp       int    `json:"exp"`
	Iss       string `json:"iss"`
}
