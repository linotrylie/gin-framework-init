package codegenerator

// 导入数据map
var importMap = map[string]string{
	"time.Time": "time",
}

// 创建数据忽略的字段
var createIgnoreColumns = []string{
	"IsDelete",
	"DeleteTime",
	"CreateTime",
	"Id",
	"UpdateTime",
	"ModifyTime",
}

// 获取数据信息忽略的字段
var getDataIgnoreColumns = []string{
	"IsDelete",
	"DeleteTime",
	"Password",
	"PasswordSalt",
}

// 编辑数据忽略的字段
var updateIgnoreColumns = []string{
	"UpdateTime",
	"ModifyTime",
	"CreateTime",
	"IsDelete",
	"DeleteTime",
}

const (
	goFileExtension = ".go"
	requestPackage  = "package request"
	responsePackage = "package response"
	newLine         = "\r\n"
	autoIncrement   = "auto_increment"
	opCreate        = "Create"
	opUpdate        = "Update"
	opPageSearch    = "PageSearch"
	opInfo          = "Info"
	IsNullAbleYES   = "YES"
	IsNullAbleNO    = "NO"
	EmailColumn     = "Email"
)
