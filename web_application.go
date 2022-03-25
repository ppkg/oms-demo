package main

import (
	"bytes"
	"fmt"
	_ "oms-demo/grpc"
	_ "oms-demo/http"
	"strings"
	"time"

	"github.com/go-spring/spring-base/log"
	"github.com/go-spring/spring-base/util"
	"github.com/go-spring/spring-core/gs"
	"github.com/limitedlee/microservice/common/config"
	"github.com/maybgit/glog"
	"github.com/spf13/cast"

	_ "github.com/go-spring/starter-echo"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	_ "github.com/go-spring/starter-redigo"
)

type webApplication struct {
	// 是否debug模式
	isDebug bool
}

func NewWebApp() *webApplication {
	return &webApplication{}
}

// 初始化
func (s *webApplication) Init() error {
	s.adaptLog()

	gs.Property("spring.application.name", "oms-demo")
	gs.Property("web.server.port", 8080)
	s.isDebug = true
	return nil
}

// 适配日志框架
func (s *webApplication) adaptLog() {
	log.SetOutput(log.FuncOutput(func(level log.Level, msg *log.Message) {
		defer func() { msg.Reuse() }()
		logFn := glog.Infof
		if level >= log.ErrorLevel {
			logFn = glog.Errorf
		} else if level == log.WarnLevel {
			logFn = glog.Warningf
		}
		var buf bytes.Buffer
		for _, a := range msg.Args() {
			buf.WriteString(cast.ToString(a))
		}
		fileLine := util.Contract(fmt.Sprintf("%s:%d", msg.File(), msg.Line()), 48)
		logFn("[%s] %s\n", fileLine, buf.String())
	}))
}

// 安装组件
func (s *webApplication) Setup() error {
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
func (s *webApplication) setupRedis() error {
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
func (s *webApplication) setupDatabase() error {
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

func (s *webApplication) instanceDatabase(name, url string) error {
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
func (s *webApplication) Run() error {
	return gs.Run()
}
