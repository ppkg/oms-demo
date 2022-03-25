package http

import (
	"context"
	"net/http"
	"oms-demo/service"

	"github.com/go-spring/spring-core/web"
)

type helloHttpServer struct {
	productService service.ProductService `autowire:""`
	GOPATH         string                 `value:"${GOPATH}"`
}

func (c *helloHttpServer) Hello(ctx web.Context) {
	list, err := c.productService.List(context.Background())
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		ctx.String(err.Error())
		return
	}
	ctx.String("%s - hello world!", c.GOPATH)
	ctx.JSON(list)
}
