package initialize

import (
	"equity/global"
	"equity/pkg/logger"
	"equity/pkg/setting"
	"equity/pkg/tracer"
	"errors"
	"github.com/go-redis/redis"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"time"
)

// InitConfig 初始化配置
func InitConfig() error {
	err := setUpSetting()
	if err != nil {
		return err
	}

	err = setupDBEngine()
	if err != nil {
		return err
	}

	err = setupLogger()
	if err != nil {
		return err
	}
	err = setupTracer()
	if err != nil {
		return err
	}

	err = SetUpRedis()
	if err != nil {
		return err
	}
	return nil
}

// 数据库配置
func setupDBEngine() error {
	var err error
	global.DBEnging, err = NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

// 初始化日志
func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" +
		global.AppSetting.LogFileName +
		global.AppSetting.LogFileExt

	// 使用 lumberjack 作为日志库的 io.writer
	lumberjackLogger := &lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   60, // 最大 600MB
		MaxAge:    1,  //  最大生存时间为10天
		LocalTime: true,
	}

	global.Logger = logger.NewLogger(lumberjackLogger, "", log.LstdFlags)
	global.Logger.WithCaller(2)
	return nil
}

// 初始化链路追踪
func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer(
		"blx-equity-platform",
		"127.0.0.1:6831",
	)
	if err != nil {
		return err
	}
	global.Tracer = jaegerTracer
	return nil
}

// 读取 yaml 配置文件
func setUpSetting() error {
	settings, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = settings.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = settings.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = settings.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	err = settings.ReadSection("Web", &global.WebSetting)
	if err != nil {
		return err
	}

	err = settings.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}

	err = settings.ReadSection("Redis", &global.RedisSetting)
	if err != nil {
		return err
	}

	global.RedisSetting.IdleTimeout *= time.Second
	global.JWTSetting.Expire *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WritTimeout *= time.Second
	return nil
}

// 初始化redis连接
func SetUpRedis() error {
	redisCfg := global.RedisSetting
	if redisCfg.Host == "" {
		return errors.New("获取redis配置文件失败")
	}
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Host,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})
	pong, err := client.Ping().Result()
	if err != nil {
		return err
	}
	global.Logger.InfoF("SetUpRedis ping :%s", pong)
	global.RedisClient = client
	return nil
}
