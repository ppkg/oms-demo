package http

import (
	"context"
	"fmt"
	"net/http"
	"oms-demo/model"
	"oms-demo/service"

	"github.com/go-spring/spring-core/web"
	SpringSwagger "github.com/go-spring/spring-swag"
)

type helloHttpServer struct {
	productService service.ProductService `autowire:""`
	GOPATH         string                 `value:"${GOPATH}"`
	swagger        *SpringSwagger.Swagger `autowire:""`
}

func (c *helloHttpServer) Hello(ctx web.Context) {
	list, err := c.productService.List(context.Background())
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		ctx.String(err.Error())
		return
	}
	resp := helloResponse{
		GoPath:      c.GOPATH,
		Message:     "请求成功",
		ProductList: list,
	}
	fmt.Println(c.swagger.ReadDoc())
	ctx.JSON(resp)
}

type helloResponse struct {
	GoPath      string           `json:"goPath"`
	Message     string           `json:"message"`
	ProductList []*model.Product `json:"productList"`
}
