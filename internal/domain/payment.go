package domain

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Payment struct {
	ID           string    `gorm:"unique;not null;index;primary_key" json:"id"`
	OrderID      string    `gorm:"not null;size:255" json:"order_id"`
	PaymentToken string    `gorm:"not null" json:"payment_token"`
	PaymentURL   string    `gorm:"not null" json:"payment_url"`
	IsActive     bool      `json:"is_active"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (payment *Payment) BeforeCreate(tx *gorm.DB) error {
	payment.ID = uuid.New().String()
	payment.IsActive = true
	return nil
}

//go:generate mockery --name=PaymentRepository --output=../mocks
type PaymentRepository interface {
	Create(ctx context.Context, payment *Payment) error
}

//go:generate mockery --name=PaymentUsecase --output=../mocks
type PaymentUsecase interface {
	Create(ctx context.Context, orderID string) (*Payment, error)
}
