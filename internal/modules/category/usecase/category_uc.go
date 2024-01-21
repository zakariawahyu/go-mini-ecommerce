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

type categoryUsecase struct {
	categoryRepo domain.CategoryRepository
	ctxTimeout   time.Duration
}

func NewCategoryUsecase(categoryRepo domain.CategoryRepository, ctxTimeout time.Duration) *categoryUsecase {
	return &categoryUsecase{
		categoryRepo: categoryRepo,
		ctxTimeout:   ctxTimeout,
	}
}

func (p *categoryUsecase) ListCategories(ctx context.Context, req *req.ListCategoryReq) ([]*domain.Category, *paging.Pagination, error) {
	categories, pagination, err := p.categoryRepo.ListCategories(ctx, req)
	if err != nil {
		return nil, nil, err
	}

	return categories, pagination, nil
}

func (u *categoryUsecase) Create(ctx context.Context, req *req.CategoryCreateReq) (*domain.Category, error) {
	c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	var category domain.Category
	helper.Copy(&category, &req)
	if err := u.categoryRepo.Create(c, &category); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, response.ErrRegisterConflict
		}
		return nil, err
	}

	return &category, nil
}

func (u *categoryUsecase) Update(ctx context.Context, id string, req *req.CategoryUpdateReq) (*domain.Category, error) {
	c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	category, err := u.categoryRepo.GetByID(c, id)
	if err != nil {
		return nil, err
	}

	helper.Copy(&category, &req)
	if err := u.categoryRepo.Update(c, category); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, response.ErrRegisterConflict
		}
		return nil, err
	}

	return category, nil
}

func (u *categoryUsecase) GetBySlug(ctx context.Context, slug string) (*domain.Category, error) {
	c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	category, err := u.categoryRepo.GetBySlug(c, slug)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.ErrNotFound
		}
		return nil, err
	}

	return category, nil
}
