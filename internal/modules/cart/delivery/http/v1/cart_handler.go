package v1

import (
	"github.com/gofiber/fiber/v2"
	"go-mini-ecommerce/internal/domain"
	"go-mini-ecommerce/internal/transport/req"
	"go-mini-ecommerce/internal/transport/res"
	"go-mini-ecommerce/utils/helper"
	"go-mini-ecommerce/utils/response"
)

type cartHandler struct {
	cartUsecase domain.CartUsecase
}

func NewCartHandler(cartUsecase domain.CartUsecase) *cartHandler {
	return &cartHandler{
		cartUsecase: cartUsecase,
	}
}

// ListCarts Get all carts
// @Tags Cart
// @Summary Get all carts
// @Description Get all carts
// @Produce json
// @Success 200 {object}  res.ListCarts
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /cart [get]
func (h *cartHandler) ListCarts(c *fiber.Ctx) error {
	var request req.ListCartReq
	var ctx = c.Context()
	var customerID = c.Locals("customerID").(string)

	if err := c.QueryParser(&request); err != nil {
		return err
	}

	carts, pagination, err := h.cartUsecase.ListCarts(ctx, customerID, &request)
	if err != nil {
		return err
	}

	var res res.ListCarts
	helper.Copy(&res.Carts, &carts)
	res.Pagination = pagination

	return c.JSON(response.NewSuccessResponse(fiber.StatusOK, res))
}

// Create Create cart
// @Tags Cart
// @Summary Create new cart
// @Description Create new cart
// @Accept json
// @Produce json
// @Param body body req.CartCreateReq true "Create Cart"
// @Success 200 {object} domain.Cart
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /cart [post]
func (h *cartHandler) Create(c *fiber.Ctx) error {
	var request req.CartCreateReq
	var ctx = c.Context()
	var customerID = c.Locals("customerID").(string)

	if err := c.BodyParser(&request); err != nil {
		return err
	}

	if err := request.Validate(); err != nil {
		return err
	}

	request.CustomerID = customerID
	cart, err := h.cartUsecase.Create(ctx, &request)
	if err != nil {
		return err
	}

	return c.JSON(response.NewSuccessResponse(fiber.StatusOK, cart))
}

// Update Update cart
// @Tags Cart
// @Summary Update single cart
// @Description Update single cart by id
// @Accept json
// @Produce json
// @Param id path string true "cart id"
// @Param body body req.CartUpdateReq true "Update Cart"
// @Success 200 {object} domain.Cart
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /cart/{id} [put]
func (h *cartHandler) Update(c *fiber.Ctx) error {
	var request req.CartUpdateReq
	var ctx = c.Context()
	var customerID = c.Locals("customerID").(string)
	var id = c.Params("id")

	if err := c.BodyParser(&request); err != nil {
		return err
	}

	if err := request.Validate(); err != nil {
		return err
	}

	request.CustomerID = customerID
	cart, err := h.cartUsecase.Update(ctx, id, &request)
	if err != nil {
		return err
	}
	return c.JSON(response.NewSuccessResponse(fiber.StatusOK, cart))
}

// Delete Delete cart by id
// @Tags Cart
// @Summary Delete cart by id
// @Description Delete cart by id
// @Produce json
// @Param id path string true "cart id"
// @Success 200 {object} response.SuccessResponse
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /cart/{id} [delete]
func (h *cartHandler) Delete(c *fiber.Ctx) error {
	var ctx = c.Context()
	var customerID = c.Locals("customerID").(string)
	var id = c.Params("id")

	_, err := h.cartUsecase.Delete(ctx, id, customerID)
	if err != nil {
		return err
	}
	return c.JSON(response.NewSuccessResponse(fiber.StatusOK, fiber.Map{
		"message": "Cart deleted",
	}))
}
