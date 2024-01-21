package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"go-mini-ecommerce/config"
	"go-mini-ecommerce/internal/infrastructure/db"
	categoryRepository "go-mini-ecommerce/internal/modules/category/repository"
	categoryUsecase "go-mini-ecommerce/internal/modules/category/usecase"
	"time"
)

func Routes(r fiber.Router, cfg *config.Config, db db.MysqlDBInterface, redis *redis.Client) {
	categoryRepo := categoryRepository.NewCategoryRepository(db)
	categoryRedisRepo := categoryRepository.NewCategoryRedisRepo(redis)
	categoryUc := categoryUsecase.NewCategoryUsecase(categoryRepo, categoryRedisRepo, cfg.App.Timeout*time.Second)
	categoryHandler := NewCategoryHandler(categoryUc)

	categoryRoute := r.Group("/category")
	categoryRoute.Get("", categoryHandler.ListCategories)
	categoryRoute.Post("", categoryHandler.Create)
	categoryRoute.Put("/:id", categoryHandler.Update)
	categoryRoute.Get("/:slug", categoryHandler.GetBySlug)
}
