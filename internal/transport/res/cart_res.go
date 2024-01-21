package res

import (
	"go-mini-ecommerce/internal/domain"
	"go-mini-ecommerce/utils/paging"
)

type ListCarts struct {
	Carts      []*domain.CartWithProduct `json:"carts"`
	Pagination *paging.Pagination        `json:"pagination"`
}
