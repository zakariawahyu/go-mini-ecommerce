package req

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type RegisterReq struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r RegisterReq) Validate() error {
	return validation.ValidateStruct(
		&r,
		validation.Field(&r.FullName, validation.Required),
		validation.Field(&r.Email, validation.Required, is.Email),
		validation.Field(&r.Username, validation.Required),
		validation.Field(&r.Password, validation.Required),
	)
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r LoginReq) Validate() error {
	return validation.ValidateStruct(
		&r,
		validation.Field(&r.Email, validation.Required),
		validation.Field(&r.Password, validation.Required),
	)
}
