package domain

import (
	"context"
)

type PaymentUsecase interface {
	Create(ctx context.Context, orderID string) (string, error)
}
