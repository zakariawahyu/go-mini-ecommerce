package res

import (
	"go-mini-ecommerce/internal/domain"
	"time"
)

type OrderRes struct {
	ID         string             `json:"id"`
	CustomerID string             `json:"customer_id"`
	Status     domain.OrderStatus `gorm:"not null" json:"status"`
	TotalPrice float64            `json:"total_price"`
	OrderItems []domain.OrderItem `json:"order_items"`
	CreatedAt  time.Time          `json:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at"`
}
