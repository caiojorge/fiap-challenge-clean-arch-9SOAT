package usecase

import "context"

type FetchPaidOrdersInputDTO struct{}
type FetchPaidOrdersOutputDTO struct {
	OrderID     string
	CustomerID  string
	TotalAmount float64
	Status      string
}

type FetchPaidOrdersUseCase interface {
	FetchPaidOrders(ctx context.Context) ([]*FetchPaidOrdersOutputDTO, error)
}
