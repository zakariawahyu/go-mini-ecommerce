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
