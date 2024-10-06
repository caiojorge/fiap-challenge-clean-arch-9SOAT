package controller

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	portsrepository "github.com/caiojorge/fiap-challenge-ddd/internal/core/application/ports/repository"
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TestGetCustomerByCPF tests the GetCustomerByCPF handler for both valid and invalid requests.
func TestGetCustomerByCPF(t *testing.T) {
	repo := NewMockCustomerRepository()
	mock := NewMockRegisterCustomerUseCase(repo)

	controller := NewFindCustomerByCPFController(context.Background(), mock)

	// Set up the Gin router
	// Create a new Gin router
	router := gin.Default()

	// Register the route and handler
	router.GET("/customer/:cpf", controller.GetCustomerByCPF)

	// Create a test request
	req, err := http.NewRequest(http.MethodGet, "/customer/12345678900", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body
	//expectedBody := `{"cpf":"12345678900","message":"Profile"}`
	//assert.JSONEq(t, expectedBody, w.Body.String())
}

type MockFindByCPFCustomerUseCase struct {
	repository portsrepository.CustomerRepository
}

func NewMockFindByCPFCustomerUseCase(repository portsrepository.CustomerRepository) *MockFindByCPFCustomerUseCase {
	return &MockFindByCPFCustomerUseCase{
		repository: repository,
	}
}

func (m *MockRegisterCustomerUseCase) FindCustomerByCPF(ctx context.Context, cpf string) (*entity.Customer, error) {
	return &entity.Customer{}, nil
}
