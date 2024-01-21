package repository

import (
	"context"
	"go-mini-ecommerce/internal/domain"
	"go-mini-ecommerce/internal/infrastructure/db"
	"go-mini-ecommerce/internal/transport/req"
	"go-mini-ecommerce/utils/paging"
)

type cartRepository struct {
	db db.MysqlDBInterface
}

func NewCartRepository(db db.MysqlDBInterface) *cartRepository {
	return &cartRepository{
		db: db,
	}
}

func (r *cartRepository) ListCarts(ctx context.Context, customerID string, req *req.ListCartReq) ([]*domain.CartWithProduct, *paging.Pagination, error) {
	query := []db.Query{
		db.NewQuery("customer_id = ?", customerID),
		db.NewQuery("is_active = ?", true),
	}

	order := "created_at"
	if req.OrderBy != "" {
		order = req.OrderBy
		if req.OrderDesc {
			order += " DESC"
		}
	}

	var total int64
	if err := r.db.Count(ctx, &domain.CartWithProduct{}, &total, db.WithQuery(query...)); err != nil {
		return nil, nil, err
	}

	pagination := paging.New(req.Page, req.Limit, total)

	var products []*domain.CartWithProduct
	if err := r.db.Find(
		ctx,
		&products,
		db.WithQuery(query...),
		db.WithLimit(int(pagination.Limit)),
		db.WithOffset(int(pagination.Skip)),
		db.WithOrder(order),
		db.WithPreload([]string{"Product.Category"}),
	); err != nil {
		return nil, nil, err
	}

	return products, pagination, nil
}

func (r *cartRepository) Create(ctx context.Context, cart *domain.Cart) error {
	return r.db.Create(ctx, cart)
}

func (r *cartRepository) Update(ctx context.Context, cart *domain.Cart) error {
	return r.db.Update(ctx, cart)
}

func (r *cartRepository) GetByID(ctx context.Context, id string) (*domain.Cart, error) {
	var cart domain.Cart

	query := []db.Query{
		db.NewQuery("id = ?", id),
		db.NewQuery("is_active = ?", true),
	}

	if err := r.db.FindOne(ctx, &cart, db.WithQuery(query...)); err != nil {
		return nil, err
	}

	return &cart, nil
}

func (r *cartRepository) GetPreloadByID(ctx context.Context, id string) (*domain.CartWithProduct, error) {
	var cart domain.CartWithProduct

	query := []db.Query{
		db.NewQuery("id = ?", id),
		db.NewQuery("is_active = ?", true),
	}

	if err := r.db.FindOne(ctx, &cart, db.WithQuery(query...), db.WithPreload([]string{"Product"})); err != nil {
		return nil, err
	}

	return &cart, nil
}
