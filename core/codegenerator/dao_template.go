package codegenerator

import (
	"equity/core/consts"
	"equity/utils/fileutil"
	"equity/utils/strutil"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

// GenerateModelDao 生成 Dao 层
// 如果有文件,不能覆盖
func GenerateModelDao(tableName string, tplColumns []*StructColumn, packageName string) error {
	if len(tplColumns) < 1 {
		return errors.New("没有数据表字段,请先完善信息")
	}

	modelPackageName := getModelPackageName(packageName)
	requestName := UnderscoreToLowerCamelCase(strings.Replace(tableName, consts.TablePrefix, "", 1))
	daoPackageName := "package " + requestName + "dao"
	requestName = strutil.LeftUpper(requestName)

	// 当前目录
	getwd, _ := os.Getwd()
	getwd = strings.Replace(getwd, "codegenerator", "", -1)

	// 目录
	filePath := getwd + "dao" + fileutil.FileSeparator() + packageName + "dao"

	// 文件名
	fileName := strings.Replace(tableName, consts.TablePrefix, "", 1)
	if !fileutil.DirIsExist(filePath) {
		if err := os.MkdirAll(filePath, 0777); err != nil {
			return err
		}
	}

	// 完整的文件路径
	fullPath := filePath + fileutil.FileSeparator() + fileName + "_dao" + goFileExtension
	if fileutil.FileIsExist(fullPath) {
		return errors.New(fmt.Sprintf("dao file exist,file=%s", fullPath))
	}

	allContents := strings.Builder{}
	// 新增数据的参数信息
	allContents.Write([]byte(daoPackageName + newLine))
	daoStruct := fmt.Sprintf(`type %s struct {
   *dao.Dao
}`, requestName+"Dao")

	allContents.Write([]byte(daoStruct + newLine))
	allContents.Write([]byte(generateCreate(modelPackageName, packageName, requestName, tableName) + newLine + newLine))
	allContents.Write([]byte(generatePageList(modelPackageName, packageName, requestName, tableName) + newLine + newLine))
	allContents.Write([]byte(generateGetDataById(modelPackageName, packageName, requestName, tableName) + newLine))
	bytes, err := fileutil.WriteContent(fullPath, os.O_RDWR|os.O_APPEND|os.O_CREATE, allContents.String())
	if err != nil {
		return err
	}
	if bytes < 1 {
		return errors.New(fmt.Sprintf("%s写入了0字节", fullPath))
	}
	log.Println(fmt.Sprintf("dao 文件生成成功,file=%s", fullPath))
	return nil
}

// generateGetDataById 通过Id获取数据
func generateGetDataById(modelName string, packageName string, requestName string, tableName string) string {
	firstLetter := getAbbreviationLetter(packageName)
	upperPKG := strutil.LeftUpper(packageName)
	funcName := upperPKG
	receiver := getRequestDaoName(requestName)
	template := `// Get%sById 通过Id获取数据
func (%s *%s) Get%sById(c *gin.Context, id int) (res []*%s, err error) {
	var model []*%s
	db := global.DBEnging.Where("id = ? ", id)
	err = db.Find(&model).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return model, err
	}
	return model, nil
}`
	return fmt.Sprintf(template,
		funcName,    // 注释
		firstLetter, // 接收者简写
		receiver,    // 接收者
		funcName,    // 函数名
		modelName+"."+UnderscoreToUpperCamelCase(tableName),
		modelName+"."+UnderscoreToUpperCamelCase(tableName),
	)
}

// 新增
func generateCreate(modelName string, packageName string, requestName string, tableName string) string {
	firstLetter := getAbbreviationLetter(packageName)
	upperPKG := strutil.LeftUpper(packageName)
	funcName := upperPKG + opCreate
	receiver := getRequestDaoName(requestName)

	template := `// %s 新增
func (%s *%s) %s(c *gin.Context, %s *%s) error {
	return global.DBEnging.Create(&%s).Error
}`
	return fmt.Sprintf(template,
		funcName,    // 注释
		firstLetter, // 接收者简写
		receiver,    // 接收者
		funcName,    // 函数名
		packageName,
		modelName+"."+UnderscoreToUpperCamelCase(tableName),
		packageName,
	)
}

// generatePageList 生成分页查询
func generatePageList(modelName, packageName, requestName, tableName string) string {
	firstLetter := getAbbreviationLetter(packageName)
	upperPKG := strutil.LeftUpper(packageName)
	funcName := upperPKG + opPageSearch
	receiver := getRequestDaoName(requestName)
	// 搜索参数结构体
	pageSearchModelStructName := getPageSearchModelStructName(requestName)
	// 返回结果结构体
	baseModelStructName := getBaseModelStructName(requestName)

	// 模型名称
	modelName = getModelPackageName(packageName) + "." + UnderscoreToUpperCamelCase(tableName)

	template := `// %s 分页查询
func (%s *%s) %s(c *gin.Context, req *request.%s) (list []*response.%s, total int64,err error,) {
	limit, offset := commonres.GetPager(req.PageNo, req.PageSize)
	db := global.DBEnging.Model(&%s{}).Where("is_delete = ?", "?需要手动填写参数")
    if req.Name != "" {
		db = db.Where("name LIKE ?",  commonres.LikeSearch(req.Name))
    }
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Order("id desc").Limit(limit).Offset(offset).Scan(&list).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return list, total,nil
	}
	return list, total,err
}`

	return fmt.Sprintf(template,
		funcName,    // 注释
		firstLetter, // 接收者简写
		receiver,    // 接收者
		funcName,
		pageSearchModelStructName, // 入参
		baseModelStructName,       // 返参
		modelName,                 // 数据模型名称,
	)
}

func getRequestDaoName(requestName string) string {
	return requestName + "Dao"
}
