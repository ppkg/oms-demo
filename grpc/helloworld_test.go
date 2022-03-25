package service

import (
	"context"
	"oms-demo/proto/helloworld"
	"testing"

	"github.com/go-spring/spring-base/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Test_helloworld(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{
			desc: "grpc连接",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				log.Errorf("建立连接失败：%+v", err)
			}
			client := helloworld.NewGreeterClient(conn)
			resp, err := client.SayHello(context.Background(), &helloworld.HelloRequest{
				Name: "zihua",
			})
			if err != nil {
				log.Errorf("grpc请求失败:%+v", err)
				return
			}
			log.Info(resp.Message)

		})
	}
}
