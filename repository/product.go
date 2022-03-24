package repository

import (
	"context"
	"oms-demo/model"

	"gorm.io/gorm"
)

type ProductRepository interface {
	List(ctx context.Context, db *gorm.DB) ([]*model.Product, error)
}
