package domain

import (
	"context"
	"github.com/google/uuid"
	"go-mini-ecommerce/internal/transport/req"
	"go-mini-ecommerce/utils/paging"
	"gorm.io/gorm"
	"time"
)

type Category struct {
	ID        string    `gorm:"unique;not null;index;primary_key" json:"id"`
	Slug      string    `gorm:"unique;not null;index:idx_category_slug" json:"slug"`
	Name      string    `gorm:"not null" json:"name"`
	IsActive  bool      `gorm:"not null" json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (category *Category) BeforeCreate(tx *gorm.DB) error {
	category.ID = uuid.New().String()
	category.IsActive = true
	return nil
}

//go:generate mockery --name=CategoryRepository --output=../mocks
type CategoryRepository interface {
	ListCategories(ctx context.Context, req *req.ListCategoryReq) ([]*Category, *paging.Pagination, error)
	Create(ctx context.Context, category *Category) error
	Update(ctx context.Context, category *Category) error
	GetByID(ctx context.Context, id string) (*Category, error)
	GetBySlug(ctx context.Context, slug string) (*Category, error)
}

//go:generate mockery --name=CategoryUsecase --output=../mocks
type CategoryUsecase interface {
	ListCategories(ctx context.Context, req *req.ListCategoryReq) ([]*Category, *paging.Pagination, error)
	Create(ctx context.Context, req *req.CategoryCreateReq) (*Category, error)
	Update(ctx context.Context, id string, req *req.CategoryUpdateReq) (*Category, error)
	GetBySlug(ctx context.Context, slug string) (*Category, error)
}

type CategoryRes struct {
	ID       string `json:"id"`
	Slug     string `json:"slug"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}

func (CategoryRes) TableName() string {
	return "categories"
}
