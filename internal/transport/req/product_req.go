package req

import validation "github.com/go-ozzo/ozzo-validation"

type ProductCreateReq struct {
	CategoryID  string `json:"category_id"`
	Slug        string `json:"slug"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Stock       int64  `json:"stock"`
}

func (r ProductCreateReq) Validate() error {
	return validation.ValidateStruct(
		&r,
		validation.Field(&r.CategoryID, validation.Required),
		validation.Field(&r.Slug, validation.Required),
		validation.Field(&r.Name, validation.Required),
		validation.Field(&r.Description, validation.Required),
		validation.Field(&r.Price, validation.Required),
		validation.Field(&r.Stock, validation.Required),
	)
}

type ProductUpdateReq struct {
	CategoryID  string `json:"category_id"`
	Slug        string `json:"slug"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Stock       int64  `json:"stock"`
}

func (r ProductUpdateReq) Validate() error {
	return validation.ValidateStruct(
		&r,
		validation.Field(&r.CategoryID, validation.Required),
		validation.Field(&r.Slug, validation.Required),
		validation.Field(&r.Name, validation.Required),
		validation.Field(&r.Description, validation.Required),
		validation.Field(&r.Price, validation.Required),
		validation.Field(&r.Stock, validation.Required),
	)
}

type ListProductReq struct {
	Name      string
	Slug      string
	Category  string
	Limit     int64
	Page      int64
	OrderBy   string
	OrderDesc bool
}
