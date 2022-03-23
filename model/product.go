package model

type Product struct {
	Id   int64
	Name string
}

func (s Product) TableName() string {
	return "product"
}
