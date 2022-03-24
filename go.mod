module oms-demo

go 1.18

require (
	github.com/go-spring/spring-base v1.2.0
	github.com/go-spring/spring-core v1.2.0
	github.com/go-spring/starter-redigo v1.1.0-rc3
	github.com/go-spring/starter/starter-echo v0.0.0-00010101000000-000000000000
	github.com/golang/protobuf v1.5.2
	github.com/limitedlee/microservice v0.1.7
	google.golang.org/grpc v1.45.0
	gorm.io/driver/mysql v1.3.2
	gorm.io/gorm v1.23.3
)

require (
	github.com/BurntSushi/toml v0.3.1 // indirect
	github.com/go-spring/spring-echo v1.1.0-rc3 // indirect
	github.com/go-spring/spring-redigo v1.1.0-rc3 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/gomodule/redigo v1.8.5 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	github.com/labstack/echo/v4 v4.6.1 // indirect
	github.com/labstack/gommon v0.3.0 // indirect
	github.com/magiconair/properties v1.8.5 // indirect
	github.com/mattn/go-colorable v0.1.8 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/pelletier/go-toml v1.9.4 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.1 // indirect
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5 // indirect
	golang.org/x/net v0.0.0-20210913180222-943fd674d43e // indirect
	golang.org/x/sys v0.0.0-20210910150752-751e447fb3d0 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220322021311-435b647f9ef2 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace (
	// github.com/go-spring/spring-base v1.2.0 => github.com/ppkg/go-spring/spring/spring-base v1.2.0
	// github.com/go-spring/spring-core v1.2.0 => github.com/ppkg/go-spring/spring/spring-core v1.2.0
	// github.com/go-spring/starter/starter-echo v1.2.0 => github.com/ppkg/go-spring/starter/starter-echo v1.2.0

	github.com/go-spring/spring-base => /home/zihua/Documents/goPath/src/github.com/go-spring/go-spring/spring/spring-base
	github.com/go-spring/spring-core => /home/zihua/Documents/goPath/src/github.com/go-spring/go-spring/spring/spring-core
	github.com/go-spring/starter/starter-echo => /home/zihua/Documents/goPath/src/github.com/go-spring/go-spring/starter/starter-echo
)
