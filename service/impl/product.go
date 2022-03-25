package impl

import (
	"context"
	"oms-demo/cache"
	"oms-demo/model"
	"oms-demo/repository"
	"oms-demo/service"

	"gorm.io/gorm"
)

type productServiceImpl struct {
	productRepository repository.ProductRepository `autowire:""`
	storeCache        cache.StoreCache             `autowire:""`
	productDb         *gorm.DB                     `autowire:"product-center"`
}

func NewProductService() service.ProductService {
	return &productServiceImpl{}
}

func (s *productServiceImpl) List(ctx context.Context) ([]*model.Product, error) {
	return s.productRepository.List(ctx,s.productDb)
}
