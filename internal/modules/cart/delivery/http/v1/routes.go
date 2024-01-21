package v1

import (
	"github.com/gofiber/fiber/v2"
	"go-mini-ecommerce/config"
	"go-mini-ecommerce/internal/infrastructure/db"
	"go-mini-ecommerce/internal/middleware"
	"go-mini-ecommerce/internal/modules/cart/repository"
	"go-mini-ecommerce/internal/modules/cart/usecase"
	"time"
)

func Routes(r fiber.Router, cfg *config.Config, db db.MysqlDBInterface) {
	cartRepo := repository.NewCartRepository(db)
	cartUsecase := usecase.NewCartUsecase(cartRepo, cfg.App.Timeout*time.Second)
	cartHandler := NewCartHandler(cartUsecase)

	authMiddleware := middleware.JWTAuth()

	cartRoutes := r.Group("/cart", authMiddleware)
	cartRoutes.Get("/", cartHandler.ListCarts)
	cartRoutes.Post("/", cartHandler.Create)
	cartRoutes.Put("/:id", cartHandler.Update)
	cartRoutes.Delete("/:id", cartHandler.Delete)
}
