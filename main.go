package main

import (
	"flag"
	"fmt"

	_ "oms-demo/grpc"
	_ "oms-demo/http"

	"github.com/go-spring/spring-base/log"
	"github.com/limitedlee/microservice/common/config"
	"github.com/maybgit/glog"
	"github.com/ppkg/stark"
	"github.com/ppkg/stark/app"
	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	defer glog.Flush()

	maxSize := 50 * 1024 * 1024
	application := &stark.GrpcApplication{
		Application: &stark.Application{
			Name:        "oms-demo",
			Environment: "dev",
			IsDebug:     true,
			LoadConfig: func() error {
				log.Info("加载其他配置...")
				return nil
			},
			SetupVars: func() error {
				log.Info("安装其他组件...")
				return nil
			},
		},
		Port:                     8080,
		UnaryServerInterceptors:  []grpc.UnaryServerInterceptor{},
		StreamServerInterceptors: []grpc.StreamServerInterceptor{},
		ServerOptions: []grpc.ServerOption{
			grpc.MaxRecvMsgSize(maxSize),
			grpc.MaxSendMsgSize(maxSize),
		},
	}

	// 添加redis连接
	err := application.PutDbConn(stark.DbConnInfo{
		Name: "redisClient",
		Url:  config.GetString("redis.Addr"),
		Type: stark.DbTypeRedis,
		Extras: map[string]interface{}{
			"password": config.GetString("redis.Password"),
			"db":       config.GetString("redis.DB"),
		},
	})
	if err != nil {
		fmt.Printf("添加redis连接配置异常:%+v \n", err)
		return
	}

	// 添加product-center数据库连接
	err = application.PutDbConn(stark.DbConnInfo{
		Name: "product-center",
		Url:  config.GetString("mysql.dc_product"),
		Type: stark.DbTypeMyql,
	})
	if err != nil {
		fmt.Printf("添加mysql连接配置异常:%+v \n", err)
	}

	// 添加datacenter数据库连接
	err = application.PutDbConn(stark.DbConnInfo{
		Name: "datacenter",
		Url:  config.GetString("mysql.datacenter"),
		Type: stark.DbTypeMyql,
	})
	if err != nil {
		fmt.Printf("添加mysql连接配置异常:%+v \n", err)
	}

	// 启动程序
	app.RunGrpcApplication(application)

}
