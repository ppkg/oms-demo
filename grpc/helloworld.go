package service

import (
	"context"
	"oms-demo/cache"
	"oms-demo/proto/helloworld"
	"oms-demo/service"

	"github.com/go-spring/spring-base/log"
)

type greeterGrpcServer struct {
	productService service.ProductService `autowire:""`
	storeCache     cache.StoreCache       `autowire:""`
	AppName        string                 `value:"${spring.application.name}"`
}

func (s *greeterGrpcServer) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	log.Infof("Received: %v", in.GetName())
	list, err := s.productService.List(ctx)
	resp := &helloworld.HelloReply{Message: "Hello " + in.GetName() + " from " + s.AppName}
	if err != nil {
		return nil, err
	}
	for _, v := range list {
		resp.Message += "|" + v.Name
	}
	storeInfo, err := s.storeCache.GetByFinanceCode(ctx, "CX0004")
	if err != nil {
		log.Errorf("获取门店缓存失败:%+v", err)
		return nil, err
	}
	resp.Message += "-->" + storeInfo
	return resp, nil
}
