package usecase

import (
	"context"

	ports "github.com/caiojorge/fiap-challenge-ddd/internal/core/application/ports/repository"
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
)

type OrderFindByIDUseCase struct {
	repository ports.OrderRepository
}

func NewOrderFindByID(repository ports.OrderRepository) *OrderFindByIDUseCase {
	return &OrderFindByIDUseCase{
		repository: repository,
	}
}

func (cr *OrderFindByIDUseCase) FindOrderByID(ctx context.Context, id string) (*entity.Order, error) {

	product, err := cr.repository.Find(ctx, id)
	if err != nil {
		return nil, err
	}

	return product, nil
}
