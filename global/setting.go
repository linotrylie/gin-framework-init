package global

import (
	"equity/pkg/logger"
	"equity/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSetting
	AppSetting      *setting.AppSetting
	DatabaseSetting *setting.DatabaseSettings
	WebSetting      *setting.WebConfigSetting
	Logger          *logger.Logger
	JWTSetting      *setting.JWTSetting
	RedisSetting    *setting.RedisSetting
)
