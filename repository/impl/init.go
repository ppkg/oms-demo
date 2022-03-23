package impl

import (
	"github.com/go-spring/spring-core/gs"
)

func init() {
	gs.Provide(NewProductRepository)
}
