package codegenerator

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type DBModel struct {
	DBEngine *sql.DB
	DBInfo   *DBInfo
}

type DBInfo struct {
	DBType   string
	Host     string
	UserName string
	Password string
	Charset  string
}

// TableColumn 数据表字段信息
type TableColumn struct {
	ColumnName             string      // 字段
	DataType               string      // 数据类型
	IsNullable             string      // 是否允许为空 是=yes 否=no
	ColumnKey              string      // key
	ColumnType             string      // 数据类型
	ColumnComment          string      // 注释
	Extra                  string      // 自增字段 auto_increment 就存在这个字段
	CharacterMaximumLength interface{} // 以字符为单位的最大长度, 有可能是null
}

var DBTypeToStructType = map[string]string{
	"int":        "int32",
	"tinyint":    "int8",
	"smallint":   "int",
	"mediumint":  "int64",
	"bigint":     "int64",
	"bit":        "int",
	"bool":       "bool",
	"enum":       "string",
	"set":        "string",
	"varchar":    "string",
	"char":       "string",
	"tinytext":   "string",
	"mediumtext": "string",
	"text":       "string",
	"longtext":   "string",
	"blob":       "string",
	"tinyblob":   "string",
	"mediumblob": "string",
	"longblob":   "string",
	"date":       "time.Time",
	"datetime":   "time.Time",
	"timestamp":  "time.Time",
	"time":       "time.Time",
	"float":      "float64",
	"double":     "float64",
}

func NewDBModel(info *DBInfo) *DBModel {
	return &DBModel{DBInfo: info}
}

func (m *DBModel) Connect() error {
	var err error
	s := "%s:%s@tcp(%s)/information_schema?" +
		"charset=%s&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(
		s,
		m.DBInfo.UserName,
		m.DBInfo.Password,
		m.DBInfo.Host,
		m.DBInfo.Charset,
	)
	m.DBEngine, err = sql.Open(m.DBInfo.DBType, dsn)
	if err != nil {
		return err
	}

	return nil
}

// getTableFolderName 获取数据表文件夹名称
// eg: tb_order_comment ==> return: order
func getTableFolderName(dbName string) string {
	if !strings.Contains(dbName, "_") {
		return dbName
	}
	splits := strings.Split(dbName, "_")
	if len(splits) > 1 {
		return splits[1]
	}
	return dbName
}

func (m *DBModel) GetColumns(dbName, tableName string) ([]*TableColumn, error) {
	query := "SELECT COLUMN_NAME, DATA_TYPE, COLUMN_KEY, " +
		"IS_NULLABLE, COLUMN_TYPE, COLUMN_COMMENT,EXTRA,CHARACTER_MAXIMUM_LENGTH " +
		"FROM COLUMNS WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ? "
	rows, err := m.DBEngine.Query(query, dbName, tableName)
	if err != nil {
		return nil, err
	}
	if rows == nil {
		return nil, errors.New("没有数据")
	}
	defer func() {
		_ = rows.Close()
	}()

	var columns []*TableColumn
	for rows.Next() {
		var column TableColumn
		err := rows.Scan(
			&column.ColumnName,
			&column.DataType,
			&column.ColumnKey,
			&column.IsNullable,
			&column.ColumnType,
			&column.ColumnComment,
			&column.Extra,
			&column.CharacterMaximumLength,
		)
		if err != nil {
			return nil, err
		}
		columns = append(columns, &column)
	}
	return columns, nil
}
