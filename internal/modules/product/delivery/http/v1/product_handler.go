package v1

import (
	"github.com/gofiber/fiber/v2"
	"go-mini-ecommerce/internal/domain"
	"go-mini-ecommerce/internal/transport/req"
	"go-mini-ecommerce/internal/transport/res"
	"go-mini-ecommerce/utils/helper"
	"go-mini-ecommerce/utils/response"
)

type productHandler struct {
	productUsecase domain.ProductUsecase
}

func NewProductHandler(productUsecase domain.ProductUsecase) *productHandler {
	return &productHandler{
		productUsecase: productUsecase,
	}
}

func (h *productHandler) ListProducts(c *fiber.Ctx) error {
	var request req.ListProductReq
	var ctx = c.Context()

	if err := c.QueryParser(&request); err != nil {
		return err
	}

	products, pagination, err := h.productUsecase.ListProducts(ctx, &request)
	if err != nil {
		return err
	}

	var res res.ListProducts
	helper.Copy(&res.Products, &products)
	res.Pagination = pagination

	return c.JSON(response.NewSuccessResponse(fiber.StatusOK, res))
}

func (h *productHandler) Create(c *fiber.Ctx) error {
	var request req.ProductCreateReq
	var ctx = c.Context()

	if err := c.BodyParser(&request); err != nil {
		return err
	}

	if err := request.Validate(); err != nil {
		return err
	}

	product, err := h.productUsecase.Create(ctx, &request)
	if err != nil {
		return err
	}

	return c.JSON(response.NewSuccessResponse(fiber.StatusOK, product))
}

func (h *productHandler) Update(c *fiber.Ctx) error {
	var request req.ProductUpdateReq
	var ctx = c.Context()
	var id = c.Params("id")

	if err := c.BodyParser(&request); err != nil {
		return err
	}

	if err := request.Validate(); err != nil {
		return err
	}

	product, err := h.productUsecase.Update(ctx, id, &request)
	if err != nil {
		return err
	}

	return c.JSON(response.NewSuccessResponse(fiber.StatusOK, product))
}

func (h *productHandler) GetBySlug(c *fiber.Ctx) error {
	var ctx = c.Context()
	var slug = c.Params("slug")

	product, err := h.productUsecase.GetBySlug(ctx, slug)
	if err != nil {
		return err
	}

	return c.JSON(response.NewSuccessResponse(fiber.StatusOK, product))
}
