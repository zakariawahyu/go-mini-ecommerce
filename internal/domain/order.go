package domain

import (
	"context"
	"github.com/google/uuid"
	"go-mini-ecommerce/internal/transport/req"
	"gorm.io/gorm"
	"time"
)

type OrderStatus string

const (
	OrderStatusInProgress OrderStatus = "in-progress"
	OrderStatusDone       OrderStatus = "done"
)

type Order struct {
	ID         string      `gorm:"unique;not null;index;primary_key" json:"id"`
	CustomerID string      `gorm:"not null" json:"customer_id"`
	Status     OrderStatus `gorm:"not null" json:"status"`
	TotalPrice float64     `gorm:"not null" json:"total_price"`
	OrderItems []OrderItem `json:"order_items"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}

func (order *Order) BeforeCreate(tx *gorm.DB) error {
	order.ID = uuid.New().String()
	order.Status = OrderStatusInProgress
	return nil
}

type OrderItem struct {
	ID        string    `gorm:"unique;not null;index;primary_key" json:"id"`
	OrderID   string    `gorm:"not null;size:255" json:"order_id"`
	ProductID string    `gorm:"not null;size:255" json:"product_id"`
	Product   *Product  `json:"product"`
	Quantity  int       `gorm:"not null" json:"quantity"`
	Amount    float64   `gorm:"not null" json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (orderitem *OrderItem) BeforeCreate(tx *gorm.DB) error {
	orderitem.ID = uuid.New().String()
	return nil
}

//go:generate mockery --name=OrderRepository --output=../mocks
type OrderRepository interface {
	Create(ctx context.Context, req *req.OrderCreateReq, orderItems []*OrderItem) (*Order, error)
	GetByID(ctx context.Context, id string) (*OrderRes, error)
}

//go:generate mockery --name=OrderUsecase --output=../mocks
type OrderUsecase interface {
	Create(ctx context.Context, req *req.OrderCreateReq) (*Order, error)
	GetByID(ctx context.Context, id string) (*OrderRes, error)
}

type OrderRes struct {
	ID         string         `json:"id"`
	CustomerID string         `json:"customer_id"`
	Customer   CustomerRes    `json:"customer"`
	Status     OrderStatus    `json:"status"`
	TotalPrice float64        `json:"total_price"`
	OrderItems []OrderItemRes `gorm:"foreignKey:OrderID" json:"order_items"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
}

func (OrderRes) TableName() string {
	return "orders"
}

type OrderItemRes struct {
	ID        string               `gorm:"unique;not null;index;primary_key" json:"id"`
	OrderID   string               `gorm:"not null;size:255" json:"order_id"`
	ProductID string               `gorm:"not null;size:255" json:"product_id"`
	Product   *ProductWithCategory `json:"product"`
	Quantity  int                  `gorm:"not null" json:"quantity"`
	Amount    float64              `gorm:"not null" json:"amount"`
	CreatedAt time.Time            `json:"created_at"`
	UpdatedAt time.Time            `json:"updated_at"`
}

func (OrderItemRes) TableName() string {
	return "order_items"
}
