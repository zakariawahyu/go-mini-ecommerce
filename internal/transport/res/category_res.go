package res

import (
	"go-mini-ecommerce/internal/domain"
	"go-mini-ecommerce/utils/paging"
)

type ListCategories struct {
	Categories []*domain.Category `json:"categories"`
	Pagination *paging.Pagination `json:"pagination"`
}
