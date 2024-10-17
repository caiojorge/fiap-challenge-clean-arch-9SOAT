package controller

// import (
// 	"context"
// 	"errors"
// 	"sync"

// 	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
// )

// type MockProductRepository struct {
// 	mu       sync.Mutex
// 	products map[string]*entity.Product
// }

// // NewMockCustomerRepository cria uma nova instância de um MockCustomerRepository.
// func NewMockProductRepository() *MockProductRepository {
// 	return &MockProductRepository{
// 		products: make(map[string]*entity.Product),
// 	}
// }

// // Create simula a criação de um novo cliente no repositório.
// func (repo *MockProductRepository) Create(ctx context.Context, product *entity.Product) error {
// 	repo.mu.Lock()
// 	defer repo.mu.Unlock()

// 	if _, exists := repo.products[product.ID]; exists {
// 		return errors.New("customer already exists")
// 	}

// 	repo.products[product.ID] = product
// 	return nil
// }

// // Update simula a atualização de um cliente no repositório.
// func (repo *MockProductRepository) Update(ctx context.Context, product *entity.Product) error {
// 	repo.mu.Lock()
// 	defer repo.mu.Unlock()

// 	if _, exists := repo.products[product.ID]; !exists {
// 		return errors.New("customer not found")
// 	}

// 	repo.products[product.ID] = product
// 	return nil
// }

// // Find simula a recuperação de um cliente pelo ID.
// func (repo *MockProductRepository) Find(ctx context.Context, id string) (*entity.Product, error) {
// 	repo.mu.Lock()
// 	defer repo.mu.Unlock()

// 	customer, exists := repo.products[id]
// 	if !exists {
// 		return nil, errors.New("customer not found")
// 	}
// 	return customer, nil
// }

// // FindAll simula a recuperação de uma lista de clientes.
// func (repo *MockProductRepository) FindAll(ctx context.Context) ([]*entity.Product, error) {
// 	repo.mu.Lock()
// 	defer repo.mu.Unlock()

// 	var products []*entity.Product
// 	for _, product := range repo.products {
// 		products = append(products, product)
// 	}
// 	return products, nil
// }

// // delete
// func (repo *MockProductRepository) Delete(ctx context.Context, id string) error {
// 	repo.mu.Lock()
// 	defer repo.mu.Unlock()

// 	if _, exists := repo.products[id]; !exists {
// 		return errors.New("product not found")
// 	}

// 	delete(repo.products, id)
// 	return nil
// }

// // Find simula a recuperação de um cliente pelo ID.
// func (repo *MockProductRepository) FindByName(ctx context.Context, name string) (*entity.Product, error) {
// 	repo.mu.Lock()
// 	defer repo.mu.Unlock()

// 	for _, product := range repo.products {
// 		if product.Name == name {
// 			return product, nil
// 		}
// 	}

// 	return nil, errors.New("product not found")
// }

// func (repo *MockProductRepository) FindByCategory(ctx context.Context, category string) ([]*entity.Product, error) {
// 	repo.mu.Lock()
// 	defer repo.mu.Unlock()

// 	var products []*entity.Product

// 	for _, product := range repo.products {
// 		if product.Category == category {
// 			//return product, nil
// 			products = append(products, product)
// 		}
// 	}

// 	if len(products) > 0 {
// 		return products, nil
// 	}

// 	return nil, errors.New("product not found")
// }
