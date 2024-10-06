package usecase

import (
	"context"

	ports "github.com/caiojorge/fiap-challenge-ddd/internal/core/application/ports/repository"
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
)

type OrderFindAllUseCase struct {
	repository ports.OrderRepository
}

func NewOrderFindAll(repository ports.OrderRepository) *OrderFindAllUseCase {
	return &OrderFindAllUseCase{
		repository: repository,
	}
}

// FindAllOrder busca todas as ordens
func (cr *OrderFindAllUseCase) FindAllOrders(ctx context.Context) ([]*entity.Order, error) {

	orders, err := cr.repository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return orders, nil
}
