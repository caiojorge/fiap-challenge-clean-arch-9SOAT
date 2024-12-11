package usecase

import (
	"context"

	ports "github.com/caiojorge/fiap-challenge-ddd/internal/domain/repository"
)

type ProductFindByCategoryUseCase struct {
	repository ports.ProductRepository
}

func NewProductFindByCategory(repository ports.ProductRepository) *ProductFindByCategoryUseCase {
	return &ProductFindByCategoryUseCase{
		repository: repository,
	}
}

func (cr *ProductFindByCategoryUseCase) FindProductByCategory(ctx context.Context, category string) ([]*FindProductByCategoryOutputDTO, error) {

	products, err := cr.repository.FindByCategory(ctx, category)
	if err != nil {
		return nil, err
	}

	var productsDTO []*FindProductByCategoryOutputDTO

	for _, product := range products {
		productDTO := &FindProductByCategoryOutputDTO{}
		productDTO.FromEntity(*product)

		productsDTO = append(productsDTO, productDTO)
	}

	return productsDTO, nil
}
