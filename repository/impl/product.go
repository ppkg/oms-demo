package impl

import (
	"context"
	"oms-demo/model"
	"oms-demo/repository"

	"github.com/go-spring/spring-base/log"
	"gorm.io/gorm"
)

type productRepositoryImpl struct {
	db *gorm.DB `autowire:"product-center"`
}

func NewProductRepository() repository.ProductRepository {
	return &productRepositoryImpl{}
}

func (s *productRepositoryImpl) List(ctx context.Context) ([]*model.Product, error) {
	var list []*model.Product
	err := s.db.Order("id desc").Limit(5).Find(&list).Error
	if err != nil {
		log.Errorf("查询产品列表异常:%+v", err)
		return nil, err
	}
	return list, nil
}
