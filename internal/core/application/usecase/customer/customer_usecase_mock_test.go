package usecase

import (
	"context"
	"errors"
	"sync"
	"testing"

	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/valueobject"
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/assert"
)

func TestCustomerFindByCPF(t *testing.T) {
	cpf, err := valueobject.NewCPF("123.456.789-09")
	assert.Nil(t, err)
	assert.NotNil(t, cpf)

	customer, err := entity.NewCustomer(*cpf, "John Doe", "email@email.com")
	assert.Nil(t, err)
	assert.NotNil(t, customer)

	repo := NewMockCustomerRepository()
	assert.NotNil(t, repo)

	finder := NewCustomerFindByCPF(repo)
	assert.NotNil(t, finder)

	err = repo.Create(context.Background(), customer)
	assert.Nil(t, err)

	customer2, err := finder.FindCustomerByCPF(context.Background(), "123.456.789-09")
	assert.Nil(t, err)
	assert.NotNil(t, customer2)
	assert.Equal(t, customer, customer2)

	customer3, err := finder.FindCustomerByCPF(context.Background(), "123.456.789-10")
	assert.NotNil(t, err)
	assert.Nil(t, customer3)

}

func TestCustomerRegister(t *testing.T) {

	cpf, err := valueobject.NewCPF("123.456.789-09")
	assert.Nil(t, err)
	assert.NotNil(t, cpf)

	customer, err := entity.NewCustomer(*cpf, "John Doe", "email@email.com")
	assert.Nil(t, err)
	assert.NotNil(t, customer)

	repo := NewMockCustomerRepository()
	assert.NotNil(t, repo)

	register := NewCustomerRegister(repo)
	assert.NotNil(t, register)

	err = register.RegisterCustomer(context.Background(), *customer)
	assert.Nil(t, err)

	customer2, err := repo.Find(context.Background(), "123.456.789-09")
	assert.Nil(t, err)
	assert.NotNil(t, customer2)
	assert.Equal(t, customer, customer2)

	customers, err := repo.FindAll(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, customers)
	assert.Len(t, customers, 1)

	all := NewCustomerFindAll(repo)
	assert.NotNil(t, all)

	customers2, err := all.FindAllCustomers(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, customers2)
	assert.Len(t, customers2, 1)

	updater := NewCustomerUpdate(repo)
	assert.NotNil(t, updater)

	customer3, err := entity.NewCustomer(*cpf, "John Doe update", "email@update.com")
	assert.Nil(t, err)
	assert.NotNil(t, customer3)

	err = updater.UpdateCustomer(context.Background(), *customer3)
	assert.Nil(t, err)

	customer4, err := repo.Find(context.Background(), "123.456.789-09")
	assert.Nil(t, err)
	assert.NotNil(t, customer4)
	assert.Equal(t, customer3, customer4)

}

type MockCustomerRepository struct {
	mu        sync.Mutex
	customers map[string]*entity.Customer
}

// NewMockCustomerRepository cria uma nova instância de um MockCustomerRepository.
func NewMockCustomerRepository() *MockCustomerRepository {
	return &MockCustomerRepository{
		customers: make(map[string]*entity.Customer),
	}
}

// Create simula a criação de um novo cliente no repositório.
func (repo *MockCustomerRepository) Create(ctx context.Context, customer *entity.Customer) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	if _, exists := repo.customers[customer.GetCPF().Value]; exists {
		return errors.New("customer already exists")
	}

	repo.customers[customer.GetCPF().Value] = customer
	return nil
}

// Update simula a atualização de um cliente no repositório.
func (repo *MockCustomerRepository) Update(ctx context.Context, customer *entity.Customer) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	if _, exists := repo.customers[customer.GetCPF().Value]; !exists {
		return errors.New("customer not found")
	}

	repo.customers[customer.GetCPF().Value] = customer
	return nil
}

// Find simula a recuperação de um cliente pelo ID.
func (repo *MockCustomerRepository) Find(ctx context.Context, id string) (*entity.Customer, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	customer, exists := repo.customers[id]
	if !exists {
		return nil, errors.New("customer not found")
	}
	return customer, nil
}

// FindAll simula a recuperação de uma lista de clientes.
func (repo *MockCustomerRepository) FindAll(ctx context.Context) ([]*entity.Customer, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	var customers []*entity.Customer
	for _, customer := range repo.customers {
		customers = append(customers, customer)
	}
	return customers, nil
}
