package server

import (
	"github.com/gofiber/fiber/v2"
	"go-mini-ecommerce/config"
	"go-mini-ecommerce/internal/infrastructure/db"
	cartRoutes "go-mini-ecommerce/internal/modules/cart/delivery/http/v1"
	categoryRoutes "go-mini-ecommerce/internal/modules/category/delivery/http/v1"
	customerRoutes "go-mini-ecommerce/internal/modules/customer/delivery/http/v1"
	orderRoutes "go-mini-ecommerce/internal/modules/order/delivery/http/v1"
	paymentRoutes "go-mini-ecommerce/internal/modules/payment/delivery/http/v1"
	productRoutes "go-mini-ecommerce/internal/modules/product/delivery/http/v1"
	"go-mini-ecommerce/utils/response"
)

type server struct {
	fiber *fiber.App
	cfg   *config.Config
	db    db.MysqlDBInterface
}

func NewHttpServer(cfg *config.Config, db db.MysqlDBInterface) *server {
	return &server{
		fiber: fiber.New(fiber.Config{
			ErrorHandler: response.ErrorHandler,
		}),
		cfg: cfg,
		db:  db,
	}
}

func (s *server) Run() error {
	s.fiber.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"app-name":    s.cfg.App.Name,
			"app-version": s.cfg.App.Version,
			"environment": s.cfg.App.Environment,
			"app-timeout": s.cfg.App.Timeout,
		})
	})

	v1 := s.fiber.Group("api/v1")
	customerRoutes.Routes(v1, s.cfg, s.db)
	categoryRoutes.Routes(v1, s.cfg, s.db)
	productRoutes.Routes(v1, s.cfg, s.db)
	cartRoutes.Routes(v1, s.cfg, s.db)
	orderRoutes.Routes(v1, s.cfg, s.db)
	paymentRoutes.Routes(v1, s.cfg, s.db)
	if err := s.fiber.Listen(s.cfg.App.Port); err != nil {
		return err
	}

	return nil
}
