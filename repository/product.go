package repository

import (
	"context"
	"oms-demo/model"
)

type ProductRepository interface {
	List(ctx context.Context) ([]*model.Product, error)
}
