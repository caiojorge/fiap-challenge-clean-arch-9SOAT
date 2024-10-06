package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
	portsrepository "github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/repository"
)

type ProductRegisterUseCase struct {
	repository portsrepository.ProductRepository
}

func NewProductRegister(repository portsrepository.ProductRepository) *ProductRegisterUseCase {
	return &ProductRegisterUseCase{
		repository: repository,
	}
}

// RegisterProduct registra um novo cliente.
func (cr *ProductRegisterUseCase) RegisterProduct(ctx context.Context, product *entity.Product) error {

	fmt.Println("usecase: verifica se o produto existe: " + product.GetName())
	entityFound, err := cr.repository.FindByName(ctx, product.GetName())
	if err != nil && err.Error() != "product not found" {
		fmt.Println("usecase: err: " + err.Error())
		return err
	}

	if entityFound != nil {
		fmt.Println("usecase: producto já existe: " + product.GetName())
		return errors.New("product already exists")
	}

	fmt.Println("usecase: Criando produto: " + product.GetName())
	product.DefineID()
	err = cr.repository.Create(ctx, product)
	if err != nil {
		return err
	}

	return nil
}
