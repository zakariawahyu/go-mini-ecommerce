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

// ListCategories Get all category
// @Tags Category
// @Summary Get all category
// @Description Get all category
// @Produce json
// @Success 200 {object}  res.ListCategories
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /category [get]
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

// Create Create category
// @Tags Category
// @Summary Create new category
// @Description Create new category
// @Accept json
// @Produce json
// @Param body body req.CategoryCreateReq true "Create category"
// @Success 200 {object} domain.Category
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /category [post]
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

// Update Update category
// @Tags Category
// @Summary Update single category
// @Description Update single category by id
// @Accept json
// @Produce json
// @Param id path string true "category id"
// @Param body body req.CategoryUpdateReq true "Update category"
// @Success 200 {object} domain.Category
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /category/{id} [put]
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

// GetBySlug Get category by slug
// @Tags Category
// @Summary Get category by slug
// @Description Get single category by slug
// @Produce json
// @Param slug path string true "category slug"
// @Success 200 {object} domain.Category
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /category/{slug} [get]
func (h *categoryHandler) GetBySlug(c *fiber.Ctx) error {
	var ctx = c.Context()
	var slug = c.Params("slug")

	category, err := h.categoryUsecase.GetBySlug(ctx, slug)
	if err != nil {
		return err
	}

	return c.JSON(response.NewSuccessResponse(fiber.StatusOK, category))
}
