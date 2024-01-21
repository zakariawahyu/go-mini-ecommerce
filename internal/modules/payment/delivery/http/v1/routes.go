package v1

import (
	"github.com/gofiber/fiber/v2"
	"go-mini-ecommerce/config"
	"go-mini-ecommerce/internal/infrastructure/db"
	"go-mini-ecommerce/internal/middleware"
	orderRepository "go-mini-ecommerce/internal/modules/order/repository"
	paymentUsecase "go-mini-ecommerce/internal/modules/payment/usecase"
)

func Routes(r fiber.Router, cfg *config.Config, db db.MysqlDBInterface) {
	orderRepo := orderRepository.NewOrderRepository(db)
	paymentUc := paymentUsecase.NewPaymentUsecase(orderRepo, cfg)
	paymentHandler := NewPaymentHandler(paymentUc)

	authMiddleware := middleware.JWTAuth()
	paymentRoutes := r.Group("/payment", authMiddleware)
	paymentRoutes.Get("/:id", paymentHandler.Create)
}
