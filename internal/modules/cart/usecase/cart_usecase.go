package usecase

import (
	"context"
	"errors"
	"go-mini-ecommerce/internal/domain"
	"go-mini-ecommerce/internal/transport/req"
	"go-mini-ecommerce/utils/helper"
	"go-mini-ecommerce/utils/paging"
	"go-mini-ecommerce/utils/response"
	"gorm.io/gorm"
	"time"
)

type cartUsecase struct {
	cartRepo   domain.CartRepository
	ctxTimeout time.Duration
}

func NewCartUsecase(cartRepo domain.CartRepository, ctxTimeout time.Duration) *cartUsecase {
	return &cartUsecase{
		cartRepo:   cartRepo,
		ctxTimeout: ctxTimeout,
	}
}

func (p *cartUsecase) ListCarts(ctx context.Context, customerID string, req *req.ListCartReq) ([]*domain.CartWithProduct, *paging.Pagination, error) {
	carts, pagination, err := p.cartRepo.ListCarts(ctx, customerID, req)
	if err != nil {
		return nil, nil, err
	}

	return carts, pagination, nil
}

func (u *cartUsecase) Create(ctx context.Context, req *req.CartCreateReq) (*domain.Cart, error) {
	c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	var cart domain.Cart
	helper.Copy(&cart, &req)
	if err := u.cartRepo.Create(c, &cart); err != nil {

		return nil, err
	}

	return &cart, nil
}

func (u *cartUsecase) Update(ctx context.Context, id string, req *req.CartUpdateReq) (*domain.Cart, error) {
	c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	cart, err := u.cartRepo.GetByID(c, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.ErrNotFound
		}
		return nil, err
	}

	if cart.CustomerID != req.CustomerID {
		return nil, response.ErrUnauthorized
	}

	helper.Copy(&cart, &req)
	if err := u.cartRepo.Update(c, cart); err != nil {

		return nil, err
	}

	return cart, nil
}

func (u *cartUsecase) Delete(ctx context.Context, id string, customerID string) (*domain.Cart, error) {
	c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	cart, err := u.cartRepo.GetByID(c, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.ErrNotFound
		}
		return nil, err
	}

	if cart.CustomerID != customerID {
		return nil, response.ErrUnauthorized
	}

	cart.IsActive = false
	if err := u.cartRepo.Update(c, cart); err != nil {

		return nil, err
	}

	return cart, nil
}
