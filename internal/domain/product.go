package domain

import (
	"context"
	"github.com/google/uuid"
	"go-mini-ecommerce/internal/transport/req"
	"go-mini-ecommerce/utils/paging"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID          string    `gorm:"unique;not null;index;primary_key" json:"id"`
	CategoryID  string    `gorm:"not null" json:"category_id"`
	Slug        string    `gorm:"unique;not null;index:idx_product_slug" json:"slug"`
	Name        string    `gorm:"not null" json:"name"`
	Description string    `gorm:"not null" json:"description"`
	Price       int       `gorm:"not null" json:"price"`
	Stock       int64     `gorm:"not null" json:"stock"`
	IsActive    bool      `gorm:"not null" json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (product *Product) BeforeCreate(tx *gorm.DB) error {
	product.ID = uuid.New().String()
	product.IsActive = true
	return nil
}

//go:generate mockery --name=ProductRepository --output=../mocks
type ProductRepository interface {
	ListProducts(ctx context.Context, req *req.ListProductReq) ([]*ProductWithCategory, *paging.Pagination, error)
	Create(ctx context.Context, category *Product) error
	Update(ctx context.Context, category *Product) error
	GetByID(ctx context.Context, id string) (*Product, error)
	GetBySlug(ctx context.Context, slug string) (*ProductWithCategory, error)
}

//go:generate mockery --name=ProductUsecase --output=../mocks
type ProductUsecase interface {
	ListProducts(ctx context.Context, req *req.ListProductReq) ([]*ProductWithCategory, *paging.Pagination, error)
	Create(ctx context.Context, req *req.ProductCreateReq) (*Product, error)
	Update(ctx context.Context, id string, req *req.ProductUpdateReq) (*Product, error)
	GetBySlug(ctx context.Context, slug string) (*ProductWithCategory, error)
}

type ProductRes struct {
	ID          string    `json:"id"`
	Slug        string    `json:"slug"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	Stock       int64     `json:"stock"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (ProductRes) TableName() string {
	return "products"
}

type ProductWithCategory struct {
	ID          string      `json:"id"`
	Slug        string      `json:"slug"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Price       int         `json:"price"`
	Stock       int64       `json:"stock"`
	IsActive    bool        `json:"is_active"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	CategoryID  string      `json:"-"`
	Category    CategoryRes `json:"category"`
}

func (ProductWithCategory) TableName() string {
	return "products"
}
