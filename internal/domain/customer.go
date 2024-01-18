package domain

import (
	"context"
	"github.com/google/uuid"
	"go-mini-ecommerce/internal/transport/req"
	"go-mini-ecommerce/utils/helper"
	"gorm.io/gorm"
	"time"
)

type Customer struct {
	ID        string    `gorm:"unique;not null;index;primary_key" json:"id"`
	FullName  string    `gorm:"not null" json:"full_name"`
	Email     string    `gorm:"unique;not null;index:idx_customer_email" json:"email"`
	Username  string    `gorm:"not null" json:"username"`
	Password  string    `gorm:"not null" json:"password"`
	IsActive  bool      `gorm:"not null" json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (customer *Customer) BeforeCreate(tx *gorm.DB) error {
	customer.ID = uuid.New().String()
	customer.IsActive = true
	customer.Password = helper.HashAndSalt([]byte(customer.Password))
	return nil
}

//go:generate mockery --name=CustomerRepository --output=../mocks
type CustomerRepository interface {
	Create(ctx context.Context, customer *Customer) error
	Update(ctx context.Context, customer *Customer) error
	GetByID(ctx context.Context, id string) (*Customer, error)
	GetByEmail(ctx context.Context, email string) (*Customer, error)
}

//go:generate mockery --name=CustomerUsecase --output=../mocks
type CustomerUsecase interface {
	Login(ctx context.Context, req *req.LoginReq) (*Customer, string, error)
	Register(ctx context.Context, req *req.RegisterReq) (*Customer, error)
}
