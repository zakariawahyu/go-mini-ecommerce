package v1

import (
	"github.com/gofiber/fiber/v2"
	"go-mini-ecommerce/internal/domain"
	"go-mini-ecommerce/internal/transport/req"
	res2 "go-mini-ecommerce/internal/transport/res"
	"go-mini-ecommerce/utils/helper"
	"go-mini-ecommerce/utils/response"
)

type customerHandler struct {
	customerUsecase domain.CustomerUsecase
}

func NewCustomerHandler(customerUsecase domain.CustomerUsecase) *customerHandler {
	return &customerHandler{
		customerUsecase: customerUsecase,
	}
}

// Login Auth Login
// @Tags Auth
// @Summary Login customer
// @Description Login customer
// @Accept json
// @Produce json
// @Param body body req.LoginReq true "Login customer"
// @Success 200 {object} domain.Customer
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /login [post]
func (h *customerHandler) Login(c *fiber.Ctx) error {
	var request req.LoginReq
	var ctx = c.Context()

	if err := c.BodyParser(&request); err != nil {
		return err
	}

	if err := request.Validate(); err != nil {
		return err
	}

	_, token, err := h.customerUsecase.Login(ctx, &request)
	if err != nil {
		return err
	}

	return c.JSON(response.NewSuccessResponse(fiber.StatusOK, fiber.Map{
		"token": token,
	}))
}

// Register Auth Register
// @Tags Auth
// @Summary Register customer
// @Description Register customer
// @Accept json
// @Produce json
// @Param body body req.RegisterReq true "Register customer"
// @Success 200 {object} domain.Customer
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /register [post]
func (h *customerHandler) Register(c *fiber.Ctx) error {
	var request req.RegisterReq
	var ctx = c.Context()

	if err := c.BodyParser(&request); err != nil {
		return err
	}

	if err := request.Validate(); err != nil {
		return err
	}

	customer, err := h.customerUsecase.Register(ctx, &request)
	if err != nil {
		return err
	}

	var results res2.RegisterRes
	helper.Copy(&results, &customer)
	return c.JSON(response.NewSuccessResponse(fiber.StatusOK, results))
}
