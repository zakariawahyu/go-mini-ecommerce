package v1

import (
	"github.com/gofiber/fiber/v2"
	"go-mini-ecommerce/config"
	"go-mini-ecommerce/internal/infrastructure/db"
	"go-mini-ecommerce/internal/middleware"
	cartRepo "go-mini-ecommerce/internal/modules/cart/repository"
	orderRepository "go-mini-ecommerce/internal/modules/order/repository"
	orderUsecase "go-mini-ecommerce/internal/modules/order/usecase"
	productRepository "go-mini-ecommerce/internal/modules/product/repository"
	"time"
)

func Routes(r fiber.Router, cfg *config.Config, db db.MysqlDBInterface) {
	orderRepo := orderRepository.NewOrderRepository(db)
	productRepo := productRepository.NewProductRepository(db)
	cartRepo := cartRepo.NewCartRepository(db)
	orderUc := orderUsecase.NewOrderUsecase(orderRepo, productRepo, cartRepo, cfg.App.Timeout*time.Second)
	orderHandler := NewOrderHandler(orderUc)

	authMiddleware := middleware.JWTAuth()
	orderRoutes := r.Group("/order", authMiddleware)
	orderRoutes.Post("/", orderHandler.Create)
	orderRoutes.Get("/:id", orderHandler.GetByID)
}
