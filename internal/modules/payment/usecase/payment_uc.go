package usecase

import (
	"context"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"go-mini-ecommerce/config"
	"go-mini-ecommerce/internal/domain"
	"time"
)

type paymentUsecase struct {
	orderRepo domain.OrderRepository
	cfg       *config.Config
}

func NewPaymentUsecase(orderRepo domain.OrderRepository, cfg *config.Config) *paymentUsecase {
	return &paymentUsecase{
		orderRepo: orderRepo,
		cfg:       cfg,
	}
}

func (u *paymentUsecase) Create(ctx context.Context, orderID string) (string, error) {
	c, cancel := context.WithTimeout(ctx, u.cfg.App.Timeout*time.Second)
	defer cancel()

	midtrans.ServerKey = u.cfg.Midtrans.ServerKey
	if u.cfg.App.Environment != "production" {
		midtrans.Environment = midtrans.Sandbox
	}

	order, err := u.orderRepo.GetByID(c, orderID)
	if err != nil {
		return "", err
	}

	var enabledPaymentTypes []snap.SnapPaymentType
	enabledPaymentTypes = append(enabledPaymentTypes, snap.AllSnapPaymentType...)

	itemDetail := make([]midtrans.ItemDetails, 0)
	for _, orderItem := range order.OrderItems {
		itemDetail = append(itemDetail, midtrans.ItemDetails{
			ID:    orderItem.Product.ID,
			Name:  orderItem.Product.ID,
			Price: int64(orderItem.Product.Price),
			Qty:   int32(orderItem.Quantity),
		})
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

	snapResponse, err := snap.CreateTransactionUrl(&snapRequest)
	if err != nil {
		return "", err
	}

	return snapResponse, nil
}
