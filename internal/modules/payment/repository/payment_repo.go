package repository

import (
	"context"
	"go-mini-ecommerce/internal/domain"
	"go-mini-ecommerce/internal/infrastructure/db"
)

type paymentRepository struct {
	db db.MysqlDBInterface
}

func NewPaymentRepository(db db.MysqlDBInterface) *paymentRepository {
	return &paymentRepository{
		db: db,
	}
}

func (r *paymentRepository) Create(ctx context.Context, payment *domain.Payment) error {
	return r.db.Create(ctx, payment)
}
