package req

type OrderCreateReq struct {
	CustomerID string  `json:"customer_id"`
	Carts      []Carts `json:"carts,omitempty"`
}

type Carts struct {
	CartID string `json:"cart_id"`
}

type OrderItems struct {
	ProductID string  `json:"product_id,omitempty"`
	Quantity  int     `json:"quantity"`
	Amount    float64 `json:"amount"`
}
