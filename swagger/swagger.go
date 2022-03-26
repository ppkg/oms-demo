package swagger

import (
	"github.com/go-openapi/spec"
	"github.com/go-spring/spring-core/gs"
	"github.com/go-spring/spring-core/web"
	SpringSwagger "github.com/go-spring/spring-swag"
	"github.com/go-spring/spring-swag/swagger"
)

func init() {
	gs.Provide(injectSwagger, "")
}

func injectSwagger(server web.Server) *SpringSwagger.Swagger {
	rootSW := swagger.Doc(server).
		WithDescription("This is a sample server Petstore server.  You can find out more about Swagger at [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/).  For this sample, you can use the api key `special-key` to test the authorization filters.").
		WithVersion("1.0.5").
		WithTitle("Swagger Petstore").
		WithTermsOfService("http://swagger.io/terms/").
		WithContact("", "", "apiteam@swagger.io").
		WithLicense("Apache 2.0", "http://www.apache.org/licenses/LICENSE-2.0.html").
		WithHost("petstore.swagger.io").
		WithBasePath("/v2").
		WithTags(
			spec.NewTag("pet", "Everything about your Pets", &spec.ExternalDocumentation{
				Description: "Find out more",
				URL:         "http://swagger.io",
			}),
			spec.NewTag("store", "Access to Petstore orders", nil),
			spec.NewTag("user", "Operations about user", &spec.ExternalDocumentation{
				Description: "Find out more about our store",
				URL:         "http://swagger.io",
			}),
		).
		WithSchemes("https", "http").
		WithExternalDocs(&spec.ExternalDocumentation{
			Description: "Find out more about Swagger",
			URL:         "http://swagger.io",
		}).
		AddApiKeySecurityDefinition("api_key", "header").
		AddOauth2ImplicitSecurityDefinition("petstore_auth",
			"https://petstore.swagger.io/oauth/authorize",
			map[string]string{
				"read:pets":  "read your pets",
				"write:pets": "modify pets in your account",
			})
	return rootSW
}
