package http

import (
	"net/http"
	_ "oms-demo/service/impl"

	"github.com/go-spring/spring-core/gs"
	"github.com/go-spring/spring-core/web"
	SpringSwagger "github.com/go-spring/spring-swag"
	"github.com/go-spring/spring-swag/swagger"
)

func init() {
	web.RegisterSwaggerHandler(func(r web.Router, doc string) {})
	gs.Object(new(helloHttpServer)).Init(func(c *helloHttpServer) {
		r := gs.GetMapping("/", c.Hello)
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
	})
}
