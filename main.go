package main

import (
	"context"
	"equity/global"
	"equity/initialize"
	"equity/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "go.uber.org/automaxprocs"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title 比邻星权益平台
// @version 1.0
// @description golang is a open source language
// @termsOfService https://www.google.com
func main() {
	err := initialize.ValidateCode()
	if err != nil {
		msg := fmt.Sprintf("initialize.ValidateCode err: %v", err)
		//message.SendMsgToFeiShu(fmt.Sprintf("initialize.ValidateCode err: %v", err.Error()))
		log.Fatalf(msg)
	}
	err = initialize.InitConfig()
	if err != nil {
		msg := fmt.Sprintf("initialize.InitConfig err: %v", err.Error())
		//message.SendMsgToFeiShu(msg)
		log.Fatalf(msg)
	}
	gin.SetMode(global.ServerSetting.RunMode)
	router := initialize.Routers()
	srv := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WritTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			msg := fmt.Sprintf("ListenAndServe err: %v", err)
			// message.SendMsgToFeiShu(msg)
			log.Fatal(msg)
		}
	}()

	// 优雅重启
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	if err := srv.Shutdown(ctx); err != nil {
		// message.SendMsgToFeiShu(err.Error())
	}

	if utils.RunModeIsRelease() {
		// message.SendMsgToFeiShu(fmt.Sprintf("start ok,now=%s", datetime.NowDateTime()))
	}
}
