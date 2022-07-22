package initialize

import (
	"equity/pkg/setting"
	"equity/utils"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
	"log"
	"os"
	"time"
)

type Model struct {
	Id         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

// NewDBEngine 实例化数据库连接
func NewDBEngine(databaseSetting *setting.DatabaseSettings) (*gorm.DB, error) {
	logLevel := logger.Warn
	if utils.RunModeIsDebug() {
		logLevel = logger.Info
	}
	// sql 日志
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标,前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logLevel,    // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,       // 禁用彩色打印
		},
	)

	// 主库
	dbMasterDsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local",
		databaseSetting.Master.UserName,
		databaseSetting.Master.Password,
		databaseSetting.Master.Host,
		databaseSetting.Master.DBName,
	)

	// 从库
	dbSlaveDsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local",
		databaseSetting.Slave.UserName,
		databaseSetting.Slave.Password,
		databaseSetting.Slave.Host,
		databaseSetting.Slave.DBName,
	)

	db, err := gorm.Open(mysql.Open(dbMasterDsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		return nil, err
	}
	err = db.Use(dbresolver.Register(dbresolver.Config{
		Sources:  []gorm.Dialector{mysql.Open(dbMasterDsn)},
		Replicas: []gorm.Dialector{mysql.Open(dbSlaveDsn)},
		Policy:   dbresolver.RandomPolicy{},
	}))

	if err != nil {
		return nil, err
	}

	sqlDb, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDb.SetMaxIdleConns(databaseSetting.Master.MaxIdleConns)
	sqlDb.SetMaxOpenConns(databaseSetting.Master.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Second * 600)
	return db, nil
}
