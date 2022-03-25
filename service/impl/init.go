package impl

import (
	_ "oms-demo/cache/impl"
	_ "oms-demo/repository/impl"

	"github.com/go-spring/spring-core/gs"
)

func init() {
	gs.Provide(NewProductService)
}
