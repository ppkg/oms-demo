package swagger

import (
	"oms-demo/dto"

	SpringSwagger "github.com/go-spring/spring-swag"
)

type bindDefinitionConfiguration struct {
	swagger *SpringSwagger.Swagger `autowire:""`
}

func (s *bindDefinitionConfiguration) register() {
	s.swagger.BindDefinitionWithTags(dto.HelloResponse{}, map[string]SpringSwagger.DefinitionField{
		"goPath": {
			Description: "goPath值",
		},
		"message": {
			Description: "请求信息",
		},
		"productList": {
			Description: "商品列表",
		},
	})
	s.swagger.BindDefinitions(dto.Product{})
}
