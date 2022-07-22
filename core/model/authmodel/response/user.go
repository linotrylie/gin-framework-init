package response

// AuthUserRes 用户信息
type AuthUserRes struct {
	Id           int32  `json:"id"`
	IsRoot       int8   `json:"isRoot"`
	Name         string `json:"name"`
	Account      string `json:"account"`
	Password     string `json:"password"`
	Salt         string `json:"salt"`
	CreateTime   int32  `json:"createTime"`
	UpdateTime   int32  `json:"updateTime"`
	LoginTime    int32  `json:"loginTime"`
	LoginIp      string `json:"loginIp"`
	Disable      int8   `json:"disable"`
	DeleteTime   int32  `json:"deleteTime"`
	LoginAddress string `json:"loginAddress"`
	IsDelete     int8   `json:"isDelete"`
}

// AuthUserBaseRes 用户列表信息
type AuthUserBaseRes struct {
	Id           int32  `json:"id"`
	IsRoot       int8   `json:"isRoot"`
	Name         string `json:"name"`
	Account      string `json:"account"`
	CreateTime   int32  `json:"createTime"`
	UpdateTime   int32  `json:"updateTime"`
	LoginTime    int32  `json:"loginTime"`
	LoginIp      string `json:"loginIp"`
	Disable      int8   `json:"disable"`
	LoginAddress string `json:"loginAddress"`
}

// AuthUserDetailRes 用户详情
type AuthUserDetailRes struct {
	Info *AuthUserBaseRes `json:"info"`
	Role []*int32         `json:"role"`
}

// AuthUserPwdAndSaltRes 获取用户密码和密码Salt
type AuthUserPwdAndSaltRes struct {
	Password string `json:"password"`
	Salt     string `json:"salt"`
}
