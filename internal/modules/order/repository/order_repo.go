package repository

import (
	"context"
	"go-mini-ecommerce/internal/domain"
	"go-mini-ecommerce/internal/infrastructure/db"
	"go-mini-ecommerce/internal/transport/req"
	"go-mini-ecommerce/utils/helper"
)

type orderRepository struct {
	db db.MysqlDBInterface
}

func NewOrderRepository(db db.MysqlDBInterface) *orderRepository {
	return &orderRepository{
		db: db,
	}
}

func (r *orderRepository) Create(ctx context.Context, req *req.OrderCreateReq, orderItems []*domain.OrderItem) (*domain.Order, error) {
	var order domain.Order

	var totalPrice float64
	for _, orderItem := range orderItems {
		totalPrice += orderItem.Amount
	}

	order.TotalPrice = totalPrice
	order.CustomerID = req.CustomerID

	handler := func() error {
		return r.createOrder(ctx, &order, orderItems)
	}

	err := r.db.WithTransaction(handler)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (r *orderRepository) createOrder(ctx context.Context, order *domain.Order, orderItems []*domain.OrderItem) error {
	if err := r.db.Create(ctx, order); err != nil {
		return err
	}

	for _, orderItem := range orderItems {
		orderItem.OrderID = order.ID
	}

	if err := r.db.CreateInBatches(ctx, &orderItems, len(orderItems)); err != nil {
		return err
	}

	helper.Copy(&order.OrderItems, &orderItems)
	return nil
}

func (r *orderRepository) GetByID(ctx context.Context, id string) (*domain.OrderRes, error) {
	var order domain.OrderRes

	opts := []db.FindOption{
		db.WithPreload([]string{"Customer", "OrderItems.Product.Category"}),
		db.WithQuery(db.NewQuery("id = ?", id)),
	}

	if err := r.db.FindOne(ctx, &order, opts...); err != nil {
		return nil, err
	}

	return &order, nil
}
