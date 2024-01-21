package v1

import (
	"github.com/gofiber/fiber/v2"
	"go-mini-ecommerce/internal/domain"
	"go-mini-ecommerce/internal/transport/req"
	"go-mini-ecommerce/internal/transport/res"
	"go-mini-ecommerce/utils/helper"
	"go-mini-ecommerce/utils/response"
)

type orderHandler struct {
	orderUsecase domain.OrderUsecase
}

func NewOrderHandler(orderUsecase domain.OrderUsecase) *orderHandler {
	return &orderHandler{
		orderUsecase: orderUsecase,
	}
}

func (h *orderHandler) Create(c *fiber.Ctx) error {
	var request req.OrderCreateReq
	var ctx = c.Context()
	var customerID = c.Locals("customerID").(string)

	if err := c.BodyParser(&request); err != nil {
		return err
	}

	request.CustomerID = customerID
	order, err := h.orderUsecase.Create(ctx, &request)
	if err != nil {
		return err
	}

	var result res.OrderRes
	helper.Copy(&result, &order)
	return c.JSON(response.NewSuccessResponse(fiber.StatusOK, result))
}

func (h *orderHandler) GetByID(c *fiber.Ctx) error {
	var ctx = c.Context()
	var id = c.Params("id")

	order, err := h.orderUsecase.GetByID(ctx, id)
	if err != nil {
		return err
	}

	return c.JSON(response.NewSuccessResponse(fiber.StatusOK, order))
}
