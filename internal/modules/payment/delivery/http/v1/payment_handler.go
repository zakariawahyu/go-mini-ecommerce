package v1

import (
	"github.com/gofiber/fiber/v2"
	"go-mini-ecommerce/internal/domain"
	"go-mini-ecommerce/utils/response"
)

type paymentHandler struct {
	paymentUsecase domain.PaymentUsecase
}

func NewPaymentHandler(paymentUsecase domain.PaymentUsecase) *paymentHandler {
	return &paymentHandler{
		paymentUsecase: paymentUsecase,
	}
}

func (h *paymentHandler) Create(c *fiber.Ctx) error {
	var ctx = c.Context()
	var id = c.Params("id")

	_, err := h.paymentUsecase.Create(ctx, id)
	if err != nil {
		return err
	}

	return c.JSON(response.NewSuccessResponse(fiber.StatusOK, "ok"))
}
