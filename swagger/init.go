package swagger

import (
	"github.com/go-spring/spring-core/gs"
	"github.com/ppkg/stark/swagger"
)

func init() {
	doc := &swagger.Document{
		Description:    "oms商品接口",
		Version:        "1.0.0",
		Title:          "商品管理",
		TermsOfService: "",
		Host:           "10.1.1.248",
		BasePath:       "",
		Schemes: []string{
			"http",
		},
		Id: "product",
		ApiKeySecurityDefinition: map[string]string{
			"token": "header",
		},
	}
	gs.Object(doc)
	gs.Object(new(bindDefinitionConfiguration)).Init(func(bindConf *bindDefinitionConfiguration) {
		bindConf.register()
	})
}
