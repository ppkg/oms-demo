package service

import (
	"oms-demo/proto/helloworld"
	_ "oms-demo/service/impl"

	"github.com/go-spring/spring-core/grpc"
	"github.com/go-spring/spring-core/gs"
)

func init() {
	gs.Object(new(greeterGrpcServer)).Init(func(s *greeterGrpcServer) {
		gs.GrpcServer("helloworld.Greeter", &grpc.Server{
			Register: helloworld.RegisterGreeterServer,
			Service:  s,
		})
	})
}
