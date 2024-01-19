package res

import (
	"go-mini-ecommerce/internal/domain"
	"go-mini-ecommerce/utils/paging"
)

type ListProducts struct {
	Products   []*domain.ProductWithCategory `json:"products"`
	Pagination *paging.Pagination            `json:"pagination"`
}
