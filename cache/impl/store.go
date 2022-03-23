package impl

import (
	"context"
	"oms-demo/cache"

	"github.com/go-spring/spring-core/redis"
)

type storeCacheImpl struct {
	client redis.Client `autowire:""`
}

func NewStoreCache() cache.StoreCache {
	return &storeCacheImpl{}
}

func (s *storeCacheImpl) GetByFinanceCode(ctx context.Context, financeCode string) (string, error) {
	return s.client.HGet(ctx, "warehouse:store:relation", financeCode)
}
