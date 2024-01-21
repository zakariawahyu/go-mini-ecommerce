package usecase

import (
	"context"
	"errors"
	"go-mini-ecommerce/internal/domain"
	"go-mini-ecommerce/internal/transport/req"
	"go-mini-ecommerce/utils/helper"
	"go-mini-ecommerce/utils/response"
	"gorm.io/gorm"
	"time"
)

type orderUsecase struct {
	orderRepo   domain.OrderRepository
	productRepo domain.ProductRepository
	cartRepo    domain.CartRepository
	ctxTimeout  time.Duration
}

func NewOrderUsecase(orderRepo domain.OrderRepository, productRepo domain.ProductRepository, cartRepo domain.CartRepository, ctxTimeout time.Duration) *orderUsecase {
	return &orderUsecase{
		orderRepo:   orderRepo,
		productRepo: productRepo,
		cartRepo:    cartRepo,
		ctxTimeout:  ctxTimeout,
	}
}

func (u *orderUsecase) Create(ctx context.Context, req *req.OrderCreateReq) (*domain.Order, error) {
	c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	var orderItems []*domain.OrderItem
	for _, carts := range req.Carts {
		cart, err := u.cartRepo.GetPreloadByID(c, carts.CartID)
		if err != nil {
			return nil, err
		}
		var orderItem domain.OrderItem
		orderItem.ProductID = cart.ProductID
		orderItem.Quantity = cart.Quantity
		orderItem.Amount = float64(cart.Product.Price) * float64(cart.Quantity)
		orderItems = append(orderItems, &orderItem)

		cart.IsActive = false
		var cartUpdate domain.Cart
		helper.Copy(&cartUpdate, &cart)
		if err = u.cartRepo.Update(c, &cartUpdate); err != nil {
			return nil, err
		}
	}

	order, err := u.orderRepo.Create(c, req, orderItems)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (u *orderUsecase) GetByID(ctx context.Context, id string) (*domain.OrderRes, error) {
	c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	order, err := u.orderRepo.GetByID(c, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.ErrNotFound
		}
		return nil, err
	}

	return order, nil
}
