package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/redis/go-redis/v9"
	"go-mini-ecommerce/config"
	_ "go-mini-ecommerce/docs"
	"go-mini-ecommerce/internal/infrastructure/db"
	cartRoutes "go-mini-ecommerce/internal/modules/cart/delivery/http/v1"
	categoryRoutes "go-mini-ecommerce/internal/modules/category/delivery/http/v1"
	customerRoutes "go-mini-ecommerce/internal/modules/customer/delivery/http/v1"
	orderRoutes "go-mini-ecommerce/internal/modules/order/delivery/http/v1"
	productRoutes "go-mini-ecommerce/internal/modules/product/delivery/http/v1"
	"go-mini-ecommerce/utils/response"
)

type server struct {
	fiber *fiber.App
	cfg   *config.Config
	db    db.MysqlDBInterface
	redis *redis.Client
}

func NewHttpServer(cfg *config.Config, db db.MysqlDBInterface, redis *redis.Client) *server {
	return &server{
		fiber: fiber.New(fiber.Config{
			ErrorHandler: response.ErrorHandler,
		}),
		cfg:   cfg,
		db:    db,
		redis: redis,
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

	s.fiber.Get("/swagger/*", swagger.HandlerDefault)

	v1 := s.fiber.Group("api/v1")
	customerRoutes.Routes(v1, s.cfg, s.db)
	categoryRoutes.Routes(v1, s.cfg, s.db, s.redis)
	productRoutes.Routes(v1, s.cfg, s.db, s.redis)
	cartRoutes.Routes(v1, s.cfg, s.db)
	orderRoutes.Routes(v1, s.cfg, s.db)
	if err := s.fiber.Listen(s.cfg.App.Port); err != nil {
		return err
	}

	return nil
}
