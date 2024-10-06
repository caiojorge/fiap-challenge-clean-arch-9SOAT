package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driver/api/dto"
	portsrepository "github.com/caiojorge/fiap-challenge-ddd/internal/core/application/ports/repository"
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/valueobject"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TestGetCustomerByCPF tests the GetCustomerByCPF handler for both valid and invalid requests.
func TestGetAllCustomers(t *testing.T) {

	repo := NewMockCustomerRepository()
	mock := NewMockFindAllCustomersUseCase(repo)

	controller := NewFindAllCustomersController(context.Background(), mock)

	// Set up the Gin router
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// Register the handler
	r.GET("/customer", controller.GetAllCustomers)

	req, _ := http.NewRequest(http.MethodGet, "/customer", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)

	// Now parse the JSON body into the slice of Customer structs
	var customers []*dto.CustomerDTO
	err := json.Unmarshal(resp.Body.Bytes(), &customers)
	assert.NoError(t, err, "Should decode response without error")

	// Check the length of the returned slice to ensure you received two customers
	assert.Len(t, customers, 2, "Expected 2 customers in response")

	// Optionally check the actual data
	cpf1, _ := valueobject.NewCPF("400.228.165-50")
	cpf2, _ := valueobject.NewCPF("364.584.534-85")

	expectedCustomers := []*dto.CustomerDTO{
		{CPF: cpf1.GetValue(), Name: "John Doe", Email: "john@example.com"},
		{CPF: cpf2.GetValue(), Name: "Jane Doe", Email: "jane@example.com"},
	}
	assert.Equal(t, expectedCustomers, customers, "Expected returned customers to match mock data")

}

type MockFindAllCustomersUseCase struct {
	repository portsrepository.CustomerRepository
}

func NewMockFindAllCustomersUseCase(repository portsrepository.CustomerRepository) *MockFindAllCustomersUseCase {
	return &MockFindAllCustomersUseCase{
		repository: repository,
	}
}

func (m *MockFindAllCustomersUseCase) FindAllCustomers(ctx context.Context) ([]*entity.Customer, error) {

	cpf1, _ := valueobject.NewCPF("400.228.165-50")
	cpf2, _ := valueobject.NewCPF("364.584.534-85")

	expectedCustomers := []*entity.Customer{
		{CPF: *cpf1, Name: "John Doe", Email: "john@example.com"},
		{CPF: *cpf2, Name: "Jane Doe", Email: "jane@example.com"},
	}

	return expectedCustomers, nil
}
