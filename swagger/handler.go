package swagger

import (
	"net/http"

	"github.com/go-spring/spring-core/gs"
	"github.com/go-spring/spring-core/web"
	SpringSwagger "github.com/go-spring/spring-swag"
)

func init() {
	gs.Object(new(swaggerHttpServer)).Init(func(server *swaggerHttpServer) {
		gs.GetMapping("/swagger/api.json", server.Api)
		gs.HandleGet("/swagger/*", web.WrapH(http.FileServer(http.Dir("./static"))))
	})
}

type swaggerHttpServer struct {
	swagger *SpringSwagger.Swagger `autowire:""`
}

func (s *swaggerHttpServer) Api(ctx web.Context) {
	ctx.SetContentType("application/json; charset=UTF-8")
	ctx.ResponseWriter().Write([]byte((s.swagger.ReadDoc())))
}
