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

type categoryUsecase struct {
	categoryRepo      domain.CategoryRepository
	categoryRedisRepo domain.CategoryRedisRepo
	ctxTimeout        time.Duration
}

func NewCategoryUsecase(categoryRepo domain.CategoryRepository, categoryRedisRepo domain.CategoryRedisRepo, ctxTimeout time.Duration) *categoryUsecase {
	return &categoryUsecase{
		categoryRepo:      categoryRepo,
		categoryRedisRepo: categoryRedisRepo,
		ctxTimeout:        ctxTimeout,
	}
}

func (u *categoryUsecase) ListCategories(ctx context.Context, req *req.ListCategoryReq) ([]*domain.Category, *paging.Pagination, error) {
	c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	cacheCategory := []*domain.Category{}
	categoryCached, _ := u.categoryRedisRepo.Get(c, fmt.Sprintf("categories:%d-%d", req.Limit, req.Page))
	if err := json.Unmarshal([]byte(categoryCached), &cacheCategory); err == nil {
		return cacheCategory, nil, err
	}

	categories, pagination, err := u.categoryRepo.ListCategories(c, req)
	if err != nil {
		return nil, nil, err
	}

	categoryString, _ := json.Marshal(&categories)
	if err = u.categoryRedisRepo.Set(c, fmt.Sprintf("categories:%d-%d", req.Limit, req.Page), categoryString, 60*time.Second); err != nil {
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

	categoryString, _ := json.Marshal(&category)
	if err := u.categoryRedisRepo.Set(ctx, fmt.Sprintf("category:%s", category.Slug), categoryString, 60*time.Second); err != nil {
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

	if err := u.categoryRedisRepo.Delete(c, fmt.Sprintf("category:%s", category.Slug)); err != nil {
		return nil, err
	}

	return category, nil
}

func (u *categoryUsecase) GetBySlug(ctx context.Context, slug string) (*domain.Category, error) {
	c, cancel := context.WithTimeout(ctx, u.ctxTimeout)
	defer cancel()

	category := &domain.Category{}
	categoryCached, _ := u.categoryRedisRepo.Get(ctx, fmt.Sprintf("category:%s", slug))
	if err := json.Unmarshal([]byte(categoryCached), category); err == nil {
		return category, err
	}

	categoryRes, err := u.categoryRepo.GetBySlug(c, slug)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.ErrNotFound
		}
		return nil, err
	}

	categoryString, _ := json.Marshal(&categoryRes)
	if err := u.categoryRedisRepo.Set(ctx, fmt.Sprintf("category:%s", categoryRes.Slug), categoryString, 60*time.Second); err != nil {
		return nil, err
	}

	return categoryRes, nil
}
