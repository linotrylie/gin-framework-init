package setting

import "time"

// 和 config.yaml 一一对应

type ServerSetting struct {
	RunMode     string
	HttpPort    string
	ReadTimeout time.Duration
	WritTimeout time.Duration
}

type JWTSetting struct {
	Secret string
	Issuer string
	Expire time.Duration
}

type RedisSetting struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
	DB          int
}

type AppSetting struct {
	DefaultPageSize int
	MaxPageSize     int
	LogSavePath     string
	LogFileName     string
	LogFileExt      string
}

type dataBase struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type DatabaseSettings struct {
	Master dataBase
	Slave  dataBase
}

type WebConfigSetting struct {
	HomeUrl   string
	UploadUrl string
	GetCount  string
	DownUrl   string
	Title     string
}

// 存储配置信息
var sections = make(map[string]interface{})

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	// !ok 说明sections中没有相应的配置,可以避免重复
	if _, ok := sections[k]; !ok {
		sections[k] = v
	}
	return nil
}

func (s *Setting) ReadAllSection() error {
	for k, v := range sections {
		err := s.ReadSection(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}
