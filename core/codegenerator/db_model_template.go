package codegenerator

import (
	"equity/core/consts"
	"equity/utils/fileutil"
	"errors"
	"fmt"
	"os"
	"strings"
	"text/template"
)

const strcutTpl = `// Package {{.PackageName}} auto generate file, do not edit it!
package {{.PackageName}}

type {{.TableName | ToCamelCase}} struct {
{{range .Columns}} {{$length := len .Comment}} {{ if gt $length 0 }}  // {{.Comment}} {{else}} // {{.Name}}{{ end }}
 {{$typeLen := len .Type }} {{if gt $typeLen 0 }}  {{.Name | ToCamelCase}} {{.Type}} {{.Tag}}{{ else }}{{.Name}}{{ end }}
{{end}}}

// TableName {{.TableName | ToCamelCase}} 
func (model {{.TableName | ToCamelCase}}) TableName() string {
	return "{{.TableName}}"
}`

type StructTemplate struct {
	strcutTpl string
}

type StructColumn struct {
	Name                   string
	Type                   string
	Tag                    string
	Comment                string
	Extra                  string
	CharacterMaximumLength interface{}
	IsNullAble             string
}

type StructTemplateDB struct {
	TableName   string
	Columns     []*StructColumn
	PackageName string
}

func NewStructTemplate() *StructTemplate {
	return &StructTemplate{strcutTpl: strcutTpl}
}

func (t *StructTemplate) AssemblyColumns(tbColumns []*TableColumn) []*StructColumn {
	tplColumns := make([]*StructColumn, 0, len(tbColumns))
	for _, column := range tbColumns {
		tag := fmt.Sprintf("`"+"json:"+"\"%s\""+"`", column.ColumnName)
		tplColumns = append(tplColumns, &StructColumn{
			Name:                   column.ColumnName,
			Type:                   DBTypeToStructType[column.DataType],
			Tag:                    tag,
			Comment:                column.ColumnComment,
			Extra:                  column.Extra,
			CharacterMaximumLength: column.CharacterMaximumLength,
			IsNullAble:             column.IsNullable,
		})
	}

	return tplColumns
}

// GenerateTableModel 生成数据表模型
func (t *StructTemplate) GenerateTableModel(tableName string, tplColumns []*StructColumn, packageName string) error {
	packageName = getModelPackageName(packageName)
	tpl := template.Must(template.New("sql2struct").Funcs(template.FuncMap{
		"ToCamelCase": UnderscoreToUpperCamelCase,
	}).Parse(t.strcutTpl))

	tplDB := StructTemplateDB{
		TableName:   tableName,
		Columns:     tplColumns,
		PackageName: packageName,
	}

	getwd, _ := os.Getwd()
	getwd = strings.Replace(getwd, "codegenerator", "", -1)
	filePath := getwd + "model" + fileutil.FileSeparator() + packageName

	// Create the path
	parentPath := fmt.Sprintf("%s", filePath)
	if !fileutil.DirIsExist(parentPath) {
		err := os.MkdirAll(parentPath, 0777)
		if err != nil {
			return err
		}
	}

	if strings.HasPrefix(tableName, consts.TablePrefix) {
		tableName = strings.Replace(tableName, consts.TablePrefix, "", 1)
	}

	// create the file
	fileName := fmt.Sprintf("%s%s%s%s", filePath, fileutil.FileSeparator(), tableName, goFileExtension)

	// 覆盖已有的文件
	if fileutil.FileIsExist(fileName) {
		return errors.New("file exist")
	}

	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer func() {
		_ = f.Close()
	}()

	err = tpl.Execute(f, tplDB)
	if err != nil {
		return err
	}
	fmt.Printf("\nfile generate succss:\nfile:%s\n", fileName)
	fmt.Println("append const data to file begin:\n", fileName)
	err = GenerateTableConst(tableName, tplColumns, packageName, fileName)
	fmt.Println("append const data to file end")
	fmt.Println()
	return err
}
