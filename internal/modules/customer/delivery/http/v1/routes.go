package v1

import (
	"github.com/gofiber/fiber/v2"
	"go-mini-ecommerce/config"
	"go-mini-ecommerce/internal/infrastructure/db"
	"go-mini-ecommerce/internal/modules/customer/repository"
	"go-mini-ecommerce/internal/modules/customer/usecase"
	"time"
)

func Routes(r fiber.Router, cfg *config.Config, db db.MysqlDBInterface) {
	customerRepo := repository.NewCustomerRepository(db)
	customerUsecase := usecase.NewCustomerUsecase(customerRepo, cfg.App.Timeout*time.Second)
	customerHandler := NewCustomerHandler(customerUsecase)

	authRoute := r.Group("/auth")
	authRoute.Post("/register", customerHandler.Register)
	authRoute.Post("/login", customerHandler.Login)
}
