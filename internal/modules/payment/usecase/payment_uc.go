package usecase

import (
	"context"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"go-mini-ecommerce/config"
	"go-mini-ecommerce/internal/domain"
	"go-mini-ecommerce/utils/helper"
	"strings"
	"time"
)

type paymentUsecase struct {
	paymentRepo domain.PaymentRepository
	orderRepo   domain.OrderRepository
	productRepo domain.ProductRepository
	cfg         *config.Config
}

func NewPaymentUsecase(paymentRepo domain.PaymentRepository, orderRepo domain.OrderRepository, productRepo domain.ProductRepository, cfg *config.Config) *paymentUsecase {
	return &paymentUsecase{
		paymentRepo: paymentRepo,
		orderRepo:   orderRepo,
		productRepo: productRepo,
		cfg:         cfg,
	}
}

func (u *paymentUsecase) Create(ctx context.Context, orderID string) (*domain.Payment, error) {
	c, cancel := context.WithTimeout(ctx, u.cfg.App.Timeout*time.Second)
	defer cancel()

	midtrans.ServerKey = u.cfg.Midtrans.ServerKey
	if u.cfg.App.Environment != "production" {
		midtrans.Environment = midtrans.Sandbox
	}

	order, err := u.orderRepo.GetByID(c, orderID)
	if err != nil {
		return nil, err
	}

	var enabledPaymentTypes []snap.SnapPaymentType
	enabledPaymentTypes = append(enabledPaymentTypes, snap.AllSnapPaymentType...)

	itemDetail := make([]midtrans.ItemDetails, 0)
	for _, orderItem := range order.OrderItems {
		productName := orderItem.Product.Name
		categoryName := orderItem.Product.Category.Name
		if len(productName) > 50 {
			productName = strings.TrimSpace(orderItem.Product.Name[0:50])
		}

		if len(categoryName) > 50 {
			categoryName = strings.TrimSpace(orderItem.Product.Category.Name[0:50])
		}
		itemDetail = append(itemDetail, midtrans.ItemDetails{
			ID: orderItem.Product.ID,

			Name:     productName,
			Category: categoryName,
			Price:    int64(orderItem.Product.Price),
			Qty:      int32(orderItem.Quantity),
		})

		var product domain.Product
		helper.Copy(&product, &orderItem.Product)
		product.Stock = product.Stock - int64(orderItem.Quantity)

		if err := u.productRepo.Update(c, &product); err != nil {
			return nil, err
		}
	}

	snapRequest := snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  order.ID,
			GrossAmt: int64(order.TotalPrice),
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: order.Customer.FullName,
			Email: order.Customer.Email,
		},
		Items:           &itemDetail,
		EnabledPayments: enabledPaymentTypes,
	}

	snapResponse, errSnap := snap.CreateTransaction(&snapRequest)
	if errSnap != nil {
		return nil, errSnap.GetRawError()
	}

	var payment domain.Payment
	payment.OrderID = orderID
	payment.PaymentToken = snapResponse.Token
	payment.PaymentURL = snapResponse.RedirectURL

	if err = u.paymentRepo.Create(c, &payment); err != nil {
		return nil, err
	}

	return &payment, nil
}
