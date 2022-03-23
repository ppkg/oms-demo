package cache

import "context"

type StoreCache interface {
	GetByFinanceCode(ctx context.Context, financeCode string) (string, error)
}
