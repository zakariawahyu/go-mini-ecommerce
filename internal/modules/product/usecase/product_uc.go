package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go-mini-ecommerce/internal/domain"
	"go-mini-ecommerce/internal/transport/req"
	"go-mini-ecommerce/utils/helper"
	"go-mini-ecommerce/utils/paging"
	"go-mini-ecommerce/utils/response"
	"gorm.io/gorm"
	"time"
)

type productUsecase struct {
	productRepo      domain.ProductRepository
	productRedisRepo domain.ProductRedisRepo
	categoryRepo     domain.CategoryRepository
	ctxTimeout       time.Duration
}

func NewProductUsecase(productRepo domain.ProductRepository, productRedisRepo domain.ProductRedisRepo, categoryRepo domain.CategoryRepository, ctxTimeout time.Duration) *productUsecase {
	return &productUsecase{
		productRepo:      productRepo,
		productRedisRepo: productRedisRepo,
		categoryRepo:     categoryRepo,
		ctxTimeout:       ctxTimeout,
	}
}

func (u *productUsecase) ListProducts(ctx context.Context, req *req.ListProductReq) ([]*domain.ProductWithCategory, *paging.Pagination, error) {
	c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	cacheProduct := []*domain.ProductWithCategory{}
	productCached, _ := u.productRedisRepo.Get(c, fmt.Sprintf("products:%d-%d", req.Limit, req.Page))
	if err := json.Unmarshal([]byte(productCached), &cacheProduct); err == nil {
		return cacheProduct, nil, err
	}

	products, pagination, err := u.productRepo.ListProducts(ctx, req)
	if err != nil {
		return nil, nil, err
	}

	productString, _ := json.Marshal(&products)
	if err := u.productRedisRepo.Set(c, fmt.Sprintf("products:%d-%d", req.Limit, req.Page), productString, 60*time.Second); err != nil {
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

	productString, _ := json.Marshal(&product)
	if err := u.productRedisRepo.Set(c, fmt.Sprintf("product:%s", product.Slug), productString, 60*time.Second); err != nil {
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

	if err := u.productRedisRepo.Delete(c, fmt.Sprintf("product:%s", product.Slug)); err != nil {
		return nil, err
	}

	return product, nil
}

func (u *productUsecase) GetBySlug(ctx context.Context, slug string) (*domain.ProductWithCategory, error) {
	c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	cacheProduct := &domain.ProductWithCategory{}
	productCached, _ := u.productRedisRepo.Get(c, fmt.Sprintf("product:%s", slug))
	if err := json.Unmarshal([]byte(productCached), &cacheProduct); err == nil {
		return cacheProduct, err
	}

	product, err := u.productRepo.GetBySlug(c, slug)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.ErrNotFound
		}
		return nil, err
	}

	productString, _ := json.Marshal(&product)
	if err := u.productRedisRepo.Set(c, fmt.Sprintf("product:%s", product.Slug), productString, 60*time.Second); err != nil {
		return nil, err
	}
	return product, nil
}
