package usecase

import (
	"context"
	"errors"
	"go-mini-ecommerce/internal/domain"
	"go-mini-ecommerce/internal/transport/req"
	"go-mini-ecommerce/utils/helper"
	"go-mini-ecommerce/utils/jwt"
	"go-mini-ecommerce/utils/response"
	"gorm.io/gorm"
	"time"
)

type customerUsecase struct {
	customerRepo domain.CustomerRepository
	ctxTimeout   time.Duration
}

func NewCustomerUsecase(customerRepo domain.CustomerRepository, ctxTimeout time.Duration) *customerUsecase {
	return &customerUsecase{
		customerRepo: customerRepo,
		ctxTimeout:   ctxTimeout,
	}
}

func (u *customerUsecase) Login(ctx context.Context, req *req.LoginReq) (*domain.Customer, string, error) {
	c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	user, err := u.customerRepo.GetByEmail(c, req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, "", response.ErrLoginNotFound
		}
		return nil, "", err
	}

	if err = helper.ComparePassword(user.Password, req.Password); err != nil {
		return nil, "", response.ErrPassword
	}

	tokenData := map[string]interface{}{
		"id":    user.ID,
		"email": user.Email,
	}

	token, err := jwt.GenerateToken(tokenData)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (u *customerUsecase) Register(ctx context.Context, req *req.RegisterReq) (*domain.Customer, error) {
	c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	var customer domain.Customer
	helper.Copy(&customer, &req)
	if err := u.customerRepo.Create(c, &customer); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, response.ErrRegisterConflict
		}
		return nil, err
	}

	return &customer, nil
}
