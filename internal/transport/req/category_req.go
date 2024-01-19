package req

import validation "github.com/go-ozzo/ozzo-validation"

type CategoryCreateReq struct {
	Slug string `json:"slug"`
	Name string `json:"name"`
}

func (r CategoryCreateReq) Validate() error {
	return validation.ValidateStruct(
		&r,
		validation.Field(&r.Slug, validation.Required),
		validation.Field(&r.Name, validation.Required),
	)
}

type CategoryUpdateReq struct {
	Slug string `json:"slug"`
	Name string `json:"name"`
}

func (r CategoryUpdateReq) Validate() error {
	return validation.ValidateStruct(
		&r,
		validation.Field(&r.Slug, validation.Required),
		validation.Field(&r.Name, validation.Required),
	)
}

type ListCategoryReq struct {
	Name      string
	Limit     int64
	Page      int64
	OrderBy   string
	OrderDesc bool
}
