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

// ListProducts Get all product
// @Tags Product
// @Summary Get all product
// @Description Get all product
// @Produce json
// @Success 200 {object}  res.ListProducts
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /product [get]
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

// Create Create product
// @Tags Product
// @Summary Create new product
// @Description Create new product
// @Accept json
// @Produce json
// @Param body body req.ProductCreateReq true "Create product"
// @Success 200 {object} domain.Product
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /product [post]
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

// Update Update product
// @Tags Product
// @Summary Update single product
// @Description Update single product by id
// @Accept json
// @Produce json
// @Param id path string true "product id"
// @Param body body req.ProductUpdateReq true "Update product"
// @Success 200 {object} domain.Product
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /product/{id} [put]
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

// GetBySlug Get product by slug
// @Tags Product
// @Summary Get product by slug
// @Description Get single product by slug
// @Produce json
// @Param slug path string true "product slug"
// @Success 200 {object} domain.ProductWithCategory
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /product/{slug} [get]
func (h *productHandler) GetBySlug(c *fiber.Ctx) error {
	var ctx = c.Context()
	var slug = c.Params("slug")

	product, err := h.productUsecase.GetBySlug(ctx, slug)
	if err != nil {
		return err
	}

	return c.JSON(response.NewSuccessResponse(fiber.StatusOK, product))
}
