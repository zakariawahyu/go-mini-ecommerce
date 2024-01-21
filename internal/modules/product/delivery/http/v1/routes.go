package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"go-mini-ecommerce/config"
	"go-mini-ecommerce/internal/infrastructure/db"
	categoryRepository "go-mini-ecommerce/internal/modules/category/repository"
	productRepository "go-mini-ecommerce/internal/modules/product/repository"
	productUsecase "go-mini-ecommerce/internal/modules/product/usecase"
	"time"
)

func Routes(r fiber.Router, cfg *config.Config, db db.MysqlDBInterface, redis *redis.Client) {
	productRepo := productRepository.NewProductRepository(db)
	productRedisRepo := productRepository.NewProductRedisRepo(redis)
	categoryRepo := categoryRepository.NewCategoryRepository(db)
	productUc := productUsecase.NewProductUsecase(productRepo, productRedisRepo, categoryRepo, cfg.App.Timeout*time.Second)
	productHandler := NewProductHandler(productUc)

	productRoute := r.Group("/product")
	productRoute.Get("", productHandler.ListProducts)
	productRoute.Post("", productHandler.Create)
	productRoute.Put("/:id", productHandler.Update)
	productRoute.Get("/:slug", productHandler.GetBySlug)
}
