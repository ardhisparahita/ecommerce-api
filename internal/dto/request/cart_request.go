package request

type AddToCartRequest struct {
	ProductID uint64 `json:"product_id"`
	Quantity  int    `json:"quantity"`
}
type UpdateCartRequest struct {
	Quantity int `json:"quantity"`
}
