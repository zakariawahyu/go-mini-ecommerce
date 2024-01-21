package v1

import (
	"github.com/gofiber/fiber/v2"
	"go-mini-ecommerce/config"
	"go-mini-ecommerce/internal/infrastructure/db"
	"go-mini-ecommerce/internal/modules/category/repository"
	"go-mini-ecommerce/internal/modules/category/usecase"
	"time"
)

func Routes(r fiber.Router, cfg *config.Config, db db.MysqlDBInterface) {
	categoryRepo := repository.NewCategoryRepository(db)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepo, cfg.App.Timeout*time.Second)
	categoryHandler := NewCategoryHandler(categoryUsecase)

	categoryRoute := r.Group("/category")
	categoryRoute.Get("", categoryHandler.ListCategories)
	categoryRoute.Post("", categoryHandler.Create)
	categoryRoute.Put("/:id", categoryHandler.Update)
	categoryRoute.Get("/:slug", categoryHandler.GetBySlug)
}
