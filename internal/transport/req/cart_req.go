package req

import validation "github.com/go-ozzo/ozzo-validation"

type CartCreateReq struct {
	CustomerID string `json:"user_id"`
	ProductID  string `json:"product_id"`
	Quantity   int    `json:"quantity"`
}

func (r CartCreateReq) Validate() error {
	return validation.ValidateStruct(
		&r,
		validation.Field(&r.ProductID, validation.Required),
		validation.Field(&r.Quantity, validation.Required),
	)
}

type CartUpdateReq struct {
	CustomerID string `json:"user_id"`
	ProductID  string `json:"product_id"`
	Quantity   int    `json:"quantity"`
}

func (r CartUpdateReq) Validate() error {
	return validation.ValidateStruct(
		&r,
		validation.Field(&r.ProductID, validation.Required),
		validation.Field(&r.Quantity, validation.Required),
	)
}

type ListCartReq struct {
	Limit     int64
	Page      int64
	OrderBy   string
	OrderDesc bool
}
