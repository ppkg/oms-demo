package service

import (
	"context"
	"oms-demo/model"
)

type ProductService interface {
	List(ctx context.Context) ([]*model.Product, error)
}
