package repository

import (
	"context"
	"go-mini-ecommerce/internal/domain"
	"go-mini-ecommerce/internal/infrastructure/db"
	"go-mini-ecommerce/internal/transport/req"
	"go-mini-ecommerce/utils/paging"
)

type categoryRepository struct {
	db db.MysqlDBInterface
}

func NewCategoryRepository(db db.MysqlDBInterface) *categoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (r *categoryRepository) ListCategories(ctx context.Context, req *req.ListCategoryReq) ([]*domain.Category, *paging.Pagination, error) {
	query := make([]db.Query, 0)
	if req.Name != "" {
		query = append(query, db.NewQuery("name LIKE ?", "%"+req.Name+"%"))
	}

	order := "created_at"
	if req.OrderBy != "" {
		order = req.OrderBy
		if req.OrderDesc {
			order += " DESC"
		}
	}

	var total int64
	if err := r.db.Count(ctx, &domain.Category{}, &total, db.WithQuery(query...)); err != nil {
		return nil, nil, err
	}

	pagination := paging.New(req.Page, req.Limit, total)

	var categories []*domain.Category
	if err := r.db.Find(
		ctx,
		&categories,
		db.WithQuery(query...),
		db.WithLimit(int(pagination.Limit)),
		db.WithOffset(int(pagination.Skip)),
		db.WithOrder(order),
	); err != nil {
		return nil, nil, err
	}

	return categories, pagination, nil
}

func (r *categoryRepository) Create(ctx context.Context, category *domain.Category) error {
	return r.db.Create(ctx, category)
}

func (r *categoryRepository) Update(ctx context.Context, category *domain.Category) error {
	return r.db.Update(ctx, category)
}

func (r *categoryRepository) GetByID(ctx context.Context, id string) (*domain.Category, error) {
	var category domain.Category
	query := db.NewQuery("id = ?", id)
	if err := r.db.FindOne(ctx, &category, db.WithQuery(query)); err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *categoryRepository) GetBySlug(ctx context.Context, slug string) (*domain.Category, error) {
	var category domain.Category
	query := db.NewQuery("slug = ?", slug)
	if err := r.db.FindOne(ctx, &category, db.WithQuery(query)); err != nil {
		return nil, err
	}

	return &category, nil
}
