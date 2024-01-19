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

type productUsecase struct {
	productRepo  domain.ProductRepository
	categoryRepo domain.CategoryRepository
	ctxTimeout   time.Duration
}

func NewProductUsecase(productRepo domain.ProductRepository, categoryRepo domain.CategoryRepository, ctxTimeout time.Duration) *productUsecase {
	return &productUsecase{
		productRepo:  productRepo,
		categoryRepo: categoryRepo,
		ctxTimeout:   ctxTimeout,
	}
}

func (p *productUsecase) ListProducts(ctx context.Context, req *req.ListProductReq) ([]*domain.ProductWithCategory, *paging.Pagination, error) {
	products, pagination, err := p.productRepo.ListProducts(ctx, req)
	if err != nil {
		return nil, nil, err
	}

	return products, pagination, nil
}

func (u *productUsecase) Create(ctx context.Context, req *req.ProductCreateReq) (*domain.Product, error) {
	c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	_, err := u.categoryRepo.GetByID(c, req.CategoryID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.ErrNotFound
		}
		return nil, err
	}

	var product domain.Product
	helper.Copy(&product, &req)
	if err := u.productRepo.Create(c, &product); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, response.ErrRegisterConflict
		}
		return nil, err
	}

	return &product, nil
}

func (u *productUsecase) Update(ctx context.Context, id string, req *req.ProductUpdateReq) (*domain.Product, error) {
	c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	product, err := u.productRepo.GetByID(c, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.ErrNotFound
		}
		return nil, err
	}

	_, err = u.categoryRepo.GetByID(c, req.CategoryID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.ErrNotFound
		}
		return nil, err
	}

	helper.Copy(&product, &req)
	if err := u.productRepo.Update(c, product); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, response.ErrRegisterConflict
		}
		return nil, err
	}

	return product, nil
}

func (u *productUsecase) GetBySlug(ctx context.Context, slug string) (*domain.ProductWithCategory, error) {
	c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	product, err := u.productRepo.GetBySlug(c, slug)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.ErrNotFound
		}
		return nil, err
	}

	return product, nil
}
