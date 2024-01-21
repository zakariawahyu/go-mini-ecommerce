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
	orderUsecase   domain.OrderUsecase
	paymentUsecase domain.PaymentUsecase
}

func NewOrderHandler(orderUsecase domain.OrderUsecase, paymentUsecase domain.PaymentUsecase) *orderHandler {
	return &orderHandler{
		orderUsecase:   orderUsecase,
		paymentUsecase: paymentUsecase,
	}
}

// Create Create order
// @Tags Order
// @Summary Create new order
// @Description Create new order
// @Accept json
// @Produce json
// @Param body body req.OrderCreateReq true "Create order"
// @Success 200 {object} res.OrderWithPaymentRes
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /order [post]
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

	payment, err := h.paymentUsecase.Create(ctx, order.ID)
	var result res.OrderWithPaymentRes
	helper.Copy(&result, &order)

	result.Payment = payment
	return c.JSON(response.NewSuccessResponse(fiber.StatusOK, result))
}

// GetByID Get order by id
// @Tags Order
// @Summary Get order by id
// @Description Get single order by id
// @Produce json
// @Param id path string true "order id"
// @Success 200 {object} domain.OrderRes
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /category/{slug} [get]
func (h *orderHandler) GetByID(c *fiber.Ctx) error {
	var ctx = c.Context()
	var id = c.Params("id")

	order, err := h.orderUsecase.GetByID(ctx, id)
	if err != nil {
		return err
	}

	return c.JSON(response.NewSuccessResponse(fiber.StatusOK, order))
}
