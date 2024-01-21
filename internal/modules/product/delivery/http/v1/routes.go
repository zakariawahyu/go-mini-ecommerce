package v1

import (
	"github.com/gofiber/fiber/v2"
	"go-mini-ecommerce/config"
	"go-mini-ecommerce/internal/infrastructure/db"
	repoCategory "go-mini-ecommerce/internal/modules/category/repository"
	"go-mini-ecommerce/internal/modules/product/repository"
	"go-mini-ecommerce/internal/modules/product/usecase"
	"time"
)

func Routes(r fiber.Router, cfg *config.Config, db db.MysqlDBInterface) {
	productRepo := repository.NewProductRepository(db)
	categoryRepo := repoCategory.NewCategoryRepository(db)
	productUsecase := usecase.NewProductUsecase(productRepo, categoryRepo, cfg.App.Timeout*time.Second)
	productHandler := NewProductHandler(productUsecase)

	productRoute := r.Group("/product")
	productRoute.Get("", productHandler.ListProducts)
	productRoute.Post("", productHandler.Create)
	productRoute.Put("/:id", productHandler.Update)
	productRoute.Get("/:slug", productHandler.GetBySlug)
}
