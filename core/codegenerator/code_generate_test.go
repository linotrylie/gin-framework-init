package codegenerator

import (
	"fmt"
	"log"
	"testing"
)

var dbModel *DBModel

func init() {
	dbInfo := &DBInfo{
		DBType:   "mysql",
		Host:     "139.9.153.162:3306",
		UserName: "root",
		Password: "Xingli@1234",
		Charset:  "utf8",
	}
	dbModel = NewDBModel(dbInfo)
	err := dbModel.Connect()
	if err != nil {
		log.Fatalf("dbModel connect error: %v", err)
	}
	fmt.Println("connect db success!")
}

// TestGenerate 数据模型自动生成
func TestGenerate(t *testing.T) {
	tables := []string{
		"t_scenic_spot_order_refund_serial",
	}
	for _, tableName := range tables {
		if tableName == "" {
			break
		}
		columns, err := dbModel.GetColumns("blx-equity-platform", tableName)
		if err != nil {
			log.Fatalf("dbModel GetColumns error: %v", err)
		}
		packageName := getTableFolderName(tableName)
		template := NewStructTemplate()
		assemblyColumns := template.AssemblyColumns(columns)

		// 生成数据表,文件存在不允许覆盖
		err = template.GenerateTableModel(tableName, assemblyColumns, packageName)
		if err != nil {
			log.Println(err.Error())
		}

		// 生成 request 文件存在不允许覆盖
		err = GenerateModelRequest(tableName, assemblyColumns, packageName)
		if err != nil {
			log.Println(err.Error())
		}

		// 生成 response 文件存在不允许覆盖
		err = GenerateModelResponse(tableName, assemblyColumns, packageName)
		if err != nil {
			log.Println(err.Error())
		}

		// 生成数据库 dao 文件,存在则不允许覆盖
		// GenerateModelDao
		err = GenerateModelDao(tableName, assemblyColumns, packageName)
		if err != nil {
			log.Println(err.Error())
		}

		// todo 生成 服务层CRUD 文件存在则不允许覆盖
		// todo 生成 控制器 文件存在则不允许覆盖
	}
}
