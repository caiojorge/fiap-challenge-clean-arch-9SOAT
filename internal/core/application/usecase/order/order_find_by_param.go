package usecase

import (
	"context"

	ports "github.com/caiojorge/fiap-challenge-ddd/internal/core/application/ports/repository"
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
)

type OrderFindByParamsUseCase struct {
	repository ports.OrderRepository
}

func NewOrderFindByParams(repository ports.OrderRepository) *OrderFindByParamsUseCase {
	return &OrderFindByParamsUseCase{
		repository: repository,
	}
}

// FindAllOrder busca todas as ordens
func (cr *OrderFindByParamsUseCase) FindOrdersByParams(ctx context.Context, params map[string]interface{}) ([]*entity.Order, error) {

	orders, err := cr.repository.FindByParams(ctx, params)
	if err != nil {
		return nil, err
	}

	return orders, nil
}
