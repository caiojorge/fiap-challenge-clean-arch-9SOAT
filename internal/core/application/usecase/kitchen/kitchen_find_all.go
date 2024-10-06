package usecase

import (
	"context"
	"errors"

	ports "github.com/caiojorge/fiap-challenge-ddd/internal/core/application/ports/repository"
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
	"github.com/jinzhu/copier"
)

type KitchenFindAllUseCase struct {
	repository ports.KitchenRepository
}

func NewKitchenFindAll(repository ports.KitchenRepository) *KitchenFindAllUseCase {
	return &KitchenFindAllUseCase{
		repository: repository,
	}
}

// FindAllOrder busca todas as ordens
func (cr *KitchenFindAllUseCase) FindAllKitchen(ctx context.Context) ([]*entity.Kitchen, error) {

	models, err := cr.repository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	if len(models) == 0 {
		return nil, errors.New("kitchens not found")
	}

	var entities []*entity.Kitchen

	copier.Copy(&entities, &models)

	return entities, nil
}
