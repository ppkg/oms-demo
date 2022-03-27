package dto

type HelloResponse struct {
	// goPath值
	GoPath string `json:"goPath"`
	// 请求信息
	Message string `json:"message"`
	// 商品列表
	ProductList []Product `json:"productList"`
}

type Product struct {
	// 商品ID
	Id int64 `json:"id"`
	// 商品名称
	Name string `json:"name"`
}
