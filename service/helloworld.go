package service

import (
	"context"
	"oms-demo/cache"
	"oms-demo/proto/helloworld"
	"oms-demo/repository"

	"github.com/go-spring/spring-base/log"
	"github.com/go-spring/spring-core/gs"
	"gorm.io/gorm"
)

type GreeterService struct {
	productRepository repository.ProductRepository `autowire:""`
	storeCache        cache.StoreCache             `autowire:""`
	productDb         *gorm.DB                     `autowire:"product-center"`
	AppName           string                       `value:"${spring.application.name}"`
}

func (s *GreeterService) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	log.Infof("Received: %v", in.GetName())
	list, err := s.productRepository.List(ctx, s.productDb)
	resp := &helloworld.HelloReply{Message: "Hello " + in.GetName() + " from " + s.AppName}
	if err != nil {
		return nil, err
	}
	for _, v := range list {
		resp.Message += "|" + v.Name
	}
	storeInfo, err := s.storeCache.GetByFinanceCode(ctx, "CX0004")
	if err != nil {
		return nil, err
	}
	resp.Message += "-->" + storeInfo
	gs.NewApp()
	return resp, nil
}
