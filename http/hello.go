package http

import (
	"context"
	"net/http"
	"oms-demo/model"
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
	resp := helloResponse{
		GoPath:      c.GOPATH,
		Message:     "请求成功",
		ProductList: list,
	}
	ctx.JSON(resp)
}

type helloResponse struct {
	GoPath      string           `json:"goPath"`
	Message     string           `json:"message"`
	ProductList []*model.Product `json:"productList"`
}

func (c *helloHttpServer) route() {
	r := c.server.GetMapping("/hello2", c.Hello)
	swagger.Path(r).
		WithID("getPetById").
		WithTags("pet").
		WithDescription("Returns a single pet").
		WithSummary("Find pet by ID").
		WithProduces("application/json", "application/xml").
		AddParam(SpringSwagger.PathParam("petId", "integer", "int64").WithDescription("ID of pet to return")).
		RespondsWith(http.StatusOK, SpringSwagger.NewBindResponse(new(helloResponse), "successful operation")).
		RespondsWith(http.StatusBadRequest, SpringSwagger.NewResponse("Invalid ID supplied")).
		RespondsWith(http.StatusNotFound, SpringSwagger.NewResponse("Pet not found")).
		SecuredWith("api_key", []string{}...)
}
