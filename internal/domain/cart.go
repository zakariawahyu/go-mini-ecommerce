package domain

import (
	"context"
	"github.com/google/uuid"
	"go-mini-ecommerce/internal/transport/req"
	"go-mini-ecommerce/utils/paging"
	"gorm.io/gorm"
	"time"
)

type Cart struct {
	ID         string    `gorm:"unique;not null;index;primary_key" json:"id"`
	CustomerID string    `gorm:"not null" json:"user_id"`
	ProductID  string    `gorm:"not null" json:"product_id"`
	Quantity   int       `gorm:"not null" json:"quantity"`
	IsActive   bool      `gorm:"not null" json:"is_active"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (category *Cart) BeforeCreate(tx *gorm.DB) error {
	category.ID = uuid.New().String()
	category.IsActive = true
	return nil
}

//go:generate mockery --name=CartRepository --output=../mocks
type CartRepository interface {
	ListCarts(ctx context.Context, customerID string, req *req.ListCartReq) ([]*CartWithProduct, *paging.Pagination, error)
	Create(ctx context.Context, category *Cart) error
	Update(ctx context.Context, category *Cart) error
	GetByID(ctx context.Context, id string) (*Cart, error)
}

//go:generate mockery --name=CartUsecase --output=../mocks
type CartUsecase interface {
	ListCarts(ctx context.Context, customerID string, req *req.ListCartReq) ([]*CartWithProduct, *paging.Pagination, error)
	Create(ctx context.Context, req *req.CartCreateReq) (*Cart, error)
	Update(ctx context.Context, id string, req *req.CartUpdateReq) (*Cart, error)
	Delete(ctx context.Context, id string, customerID string) (*Cart, error)
}

type CartWithCustomerAndProduct struct {
	ID         string              `json:"id"`
	CustomerID string              `json:"-"`
	Customer   CustomerRes         `json:"customer"`
	ProductID  string              `json:"-"`
	Product    ProductWithCategory `json:"product"`
	Quantity   int                 `json:"quantity"`
	IsActive   bool                `json:"is_active"`
	CreatedAt  time.Time           `json:"created_at"`
	UpdatedAt  time.Time           `json:"updated_at"`
}

func (CartWithCustomerAndProduct) TableName() string {
	return "carts"
}

type CartWithProduct struct {
	ID        string              `json:"id"`
	ProductID string              `json:"-"`
	Product   ProductWithCategory `json:"product"`
	Quantity  int                 `json:"quantity"`
	IsActive  bool                `json:"is_active"`
	CreatedAt time.Time           `json:"created_at"`
	UpdatedAt time.Time           `json:"updated_at"`
}

func (CartWithProduct) TableName() string {
	return "carts"
}
