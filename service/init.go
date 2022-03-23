package service

import (
	_ "oms-demo/cache/impl"
	"oms-demo/proto/helloworld"
	_ "oms-demo/repository/impl"

	"github.com/go-spring/spring-core/grpc"
	"github.com/go-spring/spring-core/gs"
	"github.com/ppkg/starter-grpc/server/gateway"
)

func init() {
	gs.Object(new(GreeterService)).Init(func(s *GreeterService) {
		gs.GrpcServer("helloworld.Greeter", &grpc.Server{
			Register: helloworld.RegisterGreeterServer,
			Service:  s,
		})
	})
	gs.Object(&gateway.GatewayServer{
		Register: helloworld.RegisterGreeterHandlerFromEndpoint,
	}).Name("gateway.greeter")
}