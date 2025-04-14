package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"base/config"
	"base/initial"
	"base/router"
)

func main() {
	// 初始化Viper配置
	initial.InitViper()

	// 初始化Consul配置
	consulConfig := initial.InitConsul()

	// 注册服务
	if err := consulConfig.RegisterService(); err != nil {
		log.Fatalf("Failed to register service: %v", err)
	}

	// 初始化路由
	r := router.NewRouter()
	router := r.SetupRouter()

	// 优雅退出
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := router.Run(fmt.Sprintf(":%d", config.GlobalConfig.Server.Port)); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	<-quit
	// 注销服务
	if err := consulConfig.DeregisterService(); err != nil {
		log.Printf("Failed to deregister service: %v", err)
	}
	log.Println("Server exiting")
}
