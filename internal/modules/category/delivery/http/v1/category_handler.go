package v1

import (
	"github.com/gofiber/fiber/v2"
	"go-mini-ecommerce/internal/domain"
	"go-mini-ecommerce/internal/transport/req"
	"go-mini-ecommerce/internal/transport/res"
	"go-mini-ecommerce/utils/helper"
	"go-mini-ecommerce/utils/response"
)

type categoryHandler struct {
	categoryUsecase domain.CategoryUsecase
}

func NewCategoryHandler(categoryUsecase domain.CategoryUsecase) *categoryHandler {
	return &categoryHandler{
		categoryUsecase: categoryUsecase,
	}
}

func (h *categoryHandler) ListCategories(c *fiber.Ctx) error {
	var request req.ListCategoryReq
	var ctx = c.Context()

	if err := c.QueryParser(&request); err != nil {
		return err
	}

	categories, pagination, err := h.categoryUsecase.ListCategories(ctx, &request)
	if err != nil {
		return err
	}

	var res res.ListCategories
	helper.Copy(&res.Categories, &categories)
	res.Pagination = pagination

	return c.JSON(response.NewSuccessResponse(fiber.StatusOK, res))
}

func (h *categoryHandler) Create(c *fiber.Ctx) error {
	var request req.CategoryCreateReq
	var ctx = c.Context()

	if err := c.BodyParser(&request); err != nil {
		return err
	}

	if err := request.Validate(); err != nil {
		return err
	}

	category, err := h.categoryUsecase.Create(ctx, &request)
	if err != nil {
		return err
	}

	return c.JSON(response.NewSuccessResponse(fiber.StatusOK, category))
}

func (h *categoryHandler) Update(c *fiber.Ctx) error {
	var request req.CategoryUpdateReq
	var ctx = c.Context()
	var id = c.Params("id")

	if err := c.BodyParser(&request); err != nil {
		return err
	}

	if err := request.Validate(); err != nil {
		return err
	}

	category, err := h.categoryUsecase.Update(ctx, id, &request)
	if err != nil {
		return err
	}

	return c.JSON(response.NewSuccessResponse(fiber.StatusOK, category))
}

func (h *categoryHandler) GetBySlug(c *fiber.Ctx) error {
	var ctx = c.Context()
	var slug = c.Params("slug")

	category, err := h.categoryUsecase.GetBySlug(ctx, slug)
	if err != nil {
		return err
	}

	return c.JSON(response.NewSuccessResponse(fiber.StatusOK, category))
}
