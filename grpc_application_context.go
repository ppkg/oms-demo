package main

import (
	"fmt"
	_ "oms-demo/service"
	"strings"
	"time"

	"github.com/go-spring/spring-base/log"
	"github.com/go-spring/spring-core/gs"
	"github.com/limitedlee/microservice/common/config"

	_ "github.com/go-spring/starter/starter-echo"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	_ "github.com/go-spring/starter-redigo"
)

type grpcApplicationContext struct {
	// 是否debug模式
	isDebug bool
}

func NewGrpcApp() *grpcApplicationContext {
	return &grpcApplicationContext{}
}

// 初始化
func (s *grpcApplicationContext) Init() error {
	gs.Property("spring.application.name", "oms-demo")
	gs.Property("web.server.port", 8080)
	s.isDebug = true
	return nil
}

// 安装组件
func (s *grpcApplicationContext) Setup() error {
	// 安装数据库
	err := s.setupDatabase()
	if err != nil {
		return err
	}

	// 安装redis
	err = s.setupRedis()
	if err != nil {
		return err
	}

	return nil
}

// 安装redis
func (s *grpcApplicationContext) setupRedis() error {
	addr := config.GetString("redis.Addr")
	urlInfo := strings.Split(addr, ":")
	if len(urlInfo) != 2 {
		return fmt.Errorf("redis连接地址不正确")
	}
	gs.Property("redis.host", urlInfo[0])
	gs.Property("redis.port", urlInfo[1])
	gs.Property("redis.password", config.GetString("redis.Password"))
	gs.Property("redis.database", config.GetString("redis.DB"))
	return nil
}

// 安装数据库
func (s *grpcApplicationContext) setupDatabase() error {
	err := s.instanceDatabase("product-center", config.GetString("mysql.dc_product"))
	if err != nil {
		return err
	}

	err = s.instanceDatabase("datacenter", config.GetString("mysql.datacenter"))
	if err != nil {
		return err
	}
	return nil

}

func (s *grpcApplicationContext) instanceDatabase(name, url string) error {
	gormConf := &gorm.Config{}
	if s.isDebug {
		gormConf.Logger = logger.Default.LogMode(logger.Info)
	}
	db, err := gorm.Open(mysql.Open(url), gormConf)
	if err != nil {
		log.Errorf("初始化数据库异常:%+v", err)
		return err
	}
	// 设置数据库连接池
	sqlDB, err := db.DB()
	if err != nil {
		log.Errorf("设置数据库连接池异常:%+v", err)
		return err
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(1000)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	gs.Object(db).Name(name).Destroy(func(db *gorm.DB) {
		err = sqlDB.Close()
		if err != nil {
			log.Errorf("关闭数据库异常:%+v", err)
		}
	})
	return nil
}

// 执行业务
func (s *grpcApplicationContext) Run() error {
	return gs.Run()
}
