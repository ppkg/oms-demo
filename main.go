package main

import (
	"flag"

	"github.com/go-spring/spring-base/log"
	"github.com/maybgit/glog"
)

func main() {
	flag.Parse()
	defer glog.Flush()

	app := NewWebApp()
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
