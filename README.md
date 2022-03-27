## 命名规范

1. service层声明接口需要以XxxService方式命名，接口实现需要xxxServiceImpl方式命名
2. repository层声明接口需要以XxxRepository方式命名，接口实现需要xxxRepositoryImpl方式命名
3. cache层声明接口需要以XxxCache方式命名，接口实现需要xxxCacheImpl方式命名
4. grpc层声明struct需要以xxxGrpcServer方式命名
5. http层声明struct需要以xxxHttpServer方式命名

## 依赖注入

通过配置结构体标签来指定注入哪些对象

```go
type productServiceImpl struct {
	productRepository repository.ProductRepository `autowire:""`
	storeCache        cache.StoreCache             `autowire:""`
	productDb         *gorm.DB                     `autowire:"product-center"`
}
```

情况1：
根据数据类型来注入，找不到实例注入对象就会报错

```go
type productServiceImpl struct {
	productRepository repository.ProductRepository `autowire:""`
}
```

情况2：
在多个相同数据类型情况下，根据bean名称来注入，找不到实例注入对象就会报错

```go
type productServiceImpl struct {
	productDb         *gorm.DB                     `autowire:"product-center"`
}
```

情况3：
容器中存在就注入，不存在就不注入

```go
type productServiceImpl struct {
	storeCache        cache.StoreCache             `autowire:"?"`
}
```

## 注册实例

通过gs.Provide()来注册,第一个参数是实例化方法，后面是实例化方法参数
比如：NewProductService("zihua")

```go
func init() {
    gs.Provide(NewProductService,"zihua")
}
```

通过gs.Object()来注册

```go
func init() {
	gs.Object(new(greeterGrpcServer)).Init(func(s *greeterGrpcServer) {
		gs.GrpcServer("helloworld.Greeter", &grpc.Server{
			Register: helloworld.RegisterGreeterServer,
			Service:  s,
		})
	})
}
```

## proto生成

先安装proto生成工具

```
https://docs.buf.build/installation
```

执行生成pb文件

```
buf generate
```

