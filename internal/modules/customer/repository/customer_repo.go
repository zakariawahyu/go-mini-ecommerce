package repository

import (
	"context"
	"go-mini-ecommerce/internal/domain"
	"go-mini-ecommerce/internal/infrastructure/db"
)

type customerRepository struct {
	db db.MysqlDBInterface
}

func NewCustomerRepository(db db.MysqlDBInterface) *customerRepository {
	return &customerRepository{
		db: db,
	}
}

func (r *customerRepository) Create(ctx context.Context, customer *domain.Customer) error {
	return r.db.Create(ctx, customer)
}

func (r *customerRepository) Update(ctx context.Context, customer *domain.Customer) error {
	return r.db.Update(ctx, customer)
}

func (r *customerRepository) GetByID(ctx context.Context, id string) (*domain.Customer, error) {
	var customer domain.Customer
	query := db.NewQuery("id = ?", id)
	if err := r.db.FindOne(ctx, &customer, db.WithQuery(query)); err != nil {
		return nil, err
	}

	return &customer, nil
}

func (r *customerRepository) GetByEmail(ctx context.Context, email string) (*domain.Customer, error) {
	var customer domain.Customer
	query := db.NewQuery("email = ?", email)
	if err := r.db.FindOne(ctx, &customer, db.WithQuery(query)); err != nil {
		return nil, err
	}

	return &customer, nil
}
