package main

import (
	"github.com/go-spring/spring-base/log"
	"github.com/go-spring/spring-core/gs"
	"github.com/go-spring/spring-core/web"
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

func init() {
	gs.Object(new(Controller)).Init(func(c *Controller) {
		gs.GetMapping("/", c.Hello)
	})
}

type Controller struct {
	GOPATH string `value:"${GOPATH}"`
}

func (c *Controller) Hello(ctx web.Context) {
	ctx.String("%s - hello world!", c.GOPATH)
}
