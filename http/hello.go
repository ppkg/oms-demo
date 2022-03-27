package http

import (
	"context"
	"net/http"
	"oms-demo/dto"
	"oms-demo/service"

	"github.com/go-spring/spring-core/web"
	SpringSwagger "github.com/go-spring/spring-swag"
	"github.com/go-spring/spring-swag/swagger"
)

type helloHttpServer struct {
	productService service.ProductService `autowire:""`
	GOPATH         string                 `value:"${GOPATH}"`
	server         web.Server             `autowire:""`
}

func (c *helloHttpServer) Hello(ctx web.Context) {
	list, err := c.productService.List(context.Background())
	if err != nil {
		ctx.SetStatus(http.StatusInternalServerError)
		ctx.String(err.Error())
		return
	}
	resp := dto.HelloResponse{
		GoPath:  c.GOPATH,
		Message: "请求成功",
	}

	for _, v := range list {
		resp.ProductList = append(resp.ProductList, dto.Product{
			Id:   v.Id,
			Name: v.Name,
		})
	}
	ctx.JSON(resp)
}

// 注册路由
func (c *helloHttpServer) route() {
	r := c.server.GetMapping("/hello2", c.Hello)
	swagger.Path(r).
		WithID("查询商品").
		WithTags("商品管理").
		WithDescription("返回前面10条商品信息详情").
		WithProduces("application/json").
		AddParam(SpringSwagger.PathParam("name", "string", "").WithDescription("你的名称")).
		RespondsWith(http.StatusOK, SpringSwagger.NewBindResponse(new(dto.HelloResponse), "请求成功")).
		RespondsWith(http.StatusBadRequest, SpringSwagger.NewResponse("请求失败"))
}
