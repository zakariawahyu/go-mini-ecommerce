package repository

import (
	"context"
	"go-mini-ecommerce/internal/domain"
	"go-mini-ecommerce/internal/infrastructure/db"
	"go-mini-ecommerce/internal/transport/req"
	"go-mini-ecommerce/utils/paging"
)

type productRepository struct {
	db db.MysqlDBInterface
}

func NewProductRepository(db db.MysqlDBInterface) *productRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) ListProducts(ctx context.Context, req *req.ListProductReq) ([]*domain.ProductWithCategory, *paging.Pagination, error) {
	query := make([]db.Query, 0)
	if req.Name != "" {
		query = append(query, db.NewQuery("name LIKE ?", "%"+req.Name+"%"))
	}

	if req.Category != "" {
		query = append(query, db.NewQuery("category_id = ?", req.Category))
	}

	order := "created_at"
	if req.OrderBy != "" {
		order = req.OrderBy
		if req.OrderDesc {
			order += " DESC"
		}
	}

	var total int64
	if err := r.db.Count(ctx, &domain.ProductWithCategory{}, &total, db.WithQuery(query...)); err != nil {
		return nil, nil, err
	}

	pagination := paging.New(req.Page, req.Limit, total)

	var products []*domain.ProductWithCategory
	if err := r.db.Find(
		ctx,
		&products,
		db.WithQuery(query...),
		db.WithLimit(int(pagination.Limit)),
		db.WithOffset(int(pagination.Skip)),
		db.WithOrder(order),
		db.WithPreload([]string{"Category"}),
	); err != nil {
		return nil, nil, err
	}

	return products, pagination, nil
}

func (r *productRepository) Create(ctx context.Context, product *domain.Product) error {
	return r.db.Create(ctx, product)
}

func (r *productRepository) Update(ctx context.Context, product *domain.Product) error {
	return r.db.Update(ctx, product)
}

func (r *productRepository) GetByID(ctx context.Context, id string) (*domain.Product, error) {
	var product domain.Product
	query := db.NewQuery("id = ?", id)
	if err := r.db.FindOne(ctx, &product, db.WithQuery(query)); err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *productRepository) GetBySlug(ctx context.Context, slug string) (*domain.ProductWithCategory, error) {
	var product domain.ProductWithCategory

	opts := []db.FindOption{
		db.WithPreload([]string{"Category"}),
		db.WithQuery(db.NewQuery("slug = ?", slug)),
	}

	if err := r.db.FindOne(ctx, &product, opts...); err != nil {
		return nil, err
	}

	return &product, nil
}
