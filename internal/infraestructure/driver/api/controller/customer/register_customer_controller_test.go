package controller

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
	mocksrepository "github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/repository/mocks"
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/valueobject"
	usecase "github.com/caiojorge/fiap-challenge-ddd/internal/core/usecase/customer/register"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestPostRegisterCustomer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocksrepository.NewMockCustomerRepository(ctrl)

	mockRepo.EXPECT().
		Create(gomock.Any(), gomock.AssignableToTypeOf(&entity.Customer{})).
		Return(nil)

	cpf, err := valueobject.NewCPF("123.456.789-09")
	assert.Nil(t, err)
	assert.NotNil(t, cpf)

	customer, err := entity.NewCustomer(*cpf, "John Doe", "email@email.com")
	assert.Nil(t, err)
	assert.NotNil(t, customer)

	mockRepo.EXPECT().
		Find(gomock.Any(), customer.CPF.Value).
		Return(nil, nil).
		Times(1)

	//repo := NewMockCustomerRepository()
	//mock := NewMockRegisterCustomerUseCase(mockRepo)

	mock := usecase.NewCustomerRegister(mockRepo)

	controller := NewRegisterCustomerController(context.Background(), mock)

	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Initialize the router
	r := gin.Default()
	r.POST("/register", controller.PostRegisterCustomer)

	// Create a JSON body
	requestBody := bytes.NewBuffer([]byte(`{"cpf":"123.456.789-09", "name":"John Doe","email":"email@email.com"}`))

	// Create the HTTP request with JSON body
	req, err := http.NewRequest("POST", "/register", requestBody)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check the response
	assert.Equal(t, http.StatusOK, w.Code, "Expected response code to be 200")
	//assert.Contains(t, w.Body.String(), "customer created John Doe", "Response body should contain correct customer name")
}

// type MockRegisterCustomerUseCase struct {
// 	repository portsrepository.CustomerRepository
// }

// func NewMockRegisterCustomerUseCase(repository portsrepository.CustomerRepository) *MockRegisterCustomerUseCase {
// 	return &MockRegisterCustomerUseCase{
// 		repository: repository,
// 	}
// }

// func (m *MockRegisterCustomerUseCase) RegisterCustomer(ctx context.Context, customer entity.Customer) error {
// 	return nil
// }

// type MockCustomerRepository struct {
// 	mu        sync.Mutex
// 	customers map[string]*entity.Customer
// }

// // NewMockCustomerRepository cria uma nova instância de um MockCustomerRepository.
// func NewMockCustomerRepository() *MockCustomerRepository {
// 	return &MockCustomerRepository{
// 		customers: make(map[string]*entity.Customer),
// 	}
// }

// // Create simula a criação de um novo cliente no repositório.
// func (repo *MockCustomerRepository) Create(ctx context.Context, customer *entity.Customer) error {
// 	repo.mu.Lock()
// 	defer repo.mu.Unlock()

// 	if _, exists := repo.customers[customer.GetCPF().Value]; exists {
// 		return errors.New("customer already exists")
// 	}

// 	repo.customers[customer.GetCPF().Value] = customer
// 	return nil
// }

// // Update simula a atualização de um cliente no repositório.
// func (repo *MockCustomerRepository) Update(ctx context.Context, customer *entity.Customer) error {
// 	repo.mu.Lock()
// 	defer repo.mu.Unlock()

// 	if _, exists := repo.customers[customer.GetCPF().Value]; !exists {
// 		return errors.New("customer not found")
// 	}

// 	repo.customers[customer.GetCPF().Value] = customer
// 	return nil
// }

// // Find simula a recuperação de um cliente pelo ID.
// func (repo *MockCustomerRepository) Find(ctx context.Context, id string) (*entity.Customer, error) {
// 	repo.mu.Lock()
// 	defer repo.mu.Unlock()

// 	customer, exists := repo.customers[id]
// 	if !exists {
// 		return nil, errors.New("customer not found")
// 	}
// 	return customer, nil
// }

// // FindAll simula a recuperação de uma lista de clientes.
// func (repo *MockCustomerRepository) FindAll(ctx context.Context) ([]*entity.Customer, error) {
// 	repo.mu.Lock()
// 	defer repo.mu.Unlock()

// 	var customers []*entity.Customer
// 	for _, customer := range repo.customers {
// 		customers = append(customers, customer)
// 	}
// 	return customers, nil
// }
