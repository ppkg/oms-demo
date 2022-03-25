package main

import (
	"github.com/go-spring/spring-base/log"
)

func main() {
	app := NewGrpcApp()
	err := app.Init()
	if err != nil {
		log.Errorf("服务初始化异常:%+v", err)
		return
	}
	err = app.Setup()
	if err != nil {
		log.Errorf("服务安装组件异常:%+v", err)
		return
	}
	err = app.Run()
	if err != nil {
		log.Errorf("服务运行异常:%+v", err)
	}
}
