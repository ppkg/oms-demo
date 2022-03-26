package http

import (
	_ "oms-demo/service/impl"

	"github.com/go-spring/spring-core/gs"
)

func init() {
	gs.Object(new(helloHttpServer)).Init(func(c *helloHttpServer) {
		c.route()
	})
}
