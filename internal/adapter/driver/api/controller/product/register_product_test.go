package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
	productRepository "github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/repository"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRegisterProductController(t *testing.T) {
	repo := NewMockProductRepository()
	mock := NewMockRegisterProductUseCase(repo)
	controller := NewRegisterProductController(context.Background(), mock)

	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Initialize the router
	r := gin.Default()
	r.POST("/register", controller.PostRegisterProduct)

	// Create a JSON body
	requestBody := bytes.NewBuffer([]byte(`{"id": "1", "name":"Lanche XPTO","description":"Pão, carne, queijo e presunto","category":"Lanche","price": 100}`))

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
	assert.Equal(t, http.StatusOK, w.Code)

	// compare the response with the expected result
	assert.Equal(t, `{"id":"1","name":"Lanche XPTO","description":"Pão, carne, queijo e presunto","category":"Lanche","price":100}`, w.Body.String())
}

func TestFindAllProductsController(t *testing.T) {

	repo := NewMockProductRepository()
	mock := NewMockFindAllProductsUseCase(repo)
	controller := NewFindAllProductController(context.Background(), mock)

	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Initialize the router
	r := gin.Default()
	r.GET("/products", controller.GetAllProducts)

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/products", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check the response
	assert.Equal(t, http.StatusOK, w.Code)

	// compare the response with the expected result
	assert.Equal(t, `[{"id":"1","name":"Lanche XPTO","description":"Pão, carne, queijo e presunto","category":"Lanche","price":100}]`, w.Body.String())
}

func TestFindProductByIDController(t *testing.T) {

	repo := NewMockProductRepository()
	mock := NewMockFindProductByIDUseCase(repo)
	controller := NewFindProductByIDController(context.Background(), mock)

	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Initialize the router
	r := gin.Default()
	r.GET("/products/:id", controller.GetProductByID)

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/products/1", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check the response
	assert.Equal(t, http.StatusOK, w.Code)

	var result entity.Product
	err = json.Unmarshal(w.Body.Bytes(), &result)
	assert.Nil(t, err)

	expected := entity.Product{
		ID:          "1",
		Name:        "Lanche XPTO",
		Description: "Pão, carne, queijo e presunto",
		Category:    "Lanche",
		Price:       100,
	}

	assert.Equal(t, result, expected)

	// compare the response with the expected result
	//assert.Equal(t, `[{"id":"1","name":"Lanche XPTO","description":"Pão, carne, queijo e presunto","category":"Lanche","price":100}]`, w.Body.String())
}

type MockFindProductByIDUseCase struct {
	repository productRepository.ProductRepository
}

func NewMockFindProductByIDUseCase(repository productRepository.ProductRepository) *MockFindProductByIDUseCase {
	return &MockFindProductByIDUseCase{
		repository: repository,
	}
}

func (m *MockFindProductByIDUseCase) FindProductByID(ctx context.Context, id string) (*entity.Product, error) {
	product := entity.Product{
		ID:          id,
		Name:        "Lanche XPTO",
		Description: "Pão, carne, queijo e presunto",
		Category:    "Lanche",
		Price:       100,
	}

	return &product, nil
}

type MockFindAllProductsUseCase struct {
	repository productRepository.ProductRepository
}

func NewMockFindAllProductsUseCase(repository productRepository.ProductRepository) *MockFindAllProductsUseCase {
	return &MockFindAllProductsUseCase{
		repository: repository,
	}
}

func (m *MockFindAllProductsUseCase) FindAllProducts(ctx context.Context) ([]*entity.Product, error) {
	product := entity.Product{
		ID:          "1",
		Name:        "Lanche XPTO",
		Description: "Pão, carne, queijo e presunto",
		Category:    "Lanche",
		Price:       100,
	}

	var products []*entity.Product
	products = append(products, &product)

	return products, nil
}

type MockRegisterProductUseCase struct {
	repository productRepository.ProductRepository
}

func NewMockRegisterProductUseCase(repository productRepository.ProductRepository) *MockRegisterProductUseCase {
	return &MockRegisterProductUseCase{
		repository: repository,
	}
}

func (m *MockRegisterProductUseCase) RegisterProduct(ctx context.Context, customer *entity.Product) error {
	return nil
}

type MockProductRepository struct {
	mu       sync.Mutex
	products map[string]*entity.Product
}

// NewMockCustomerRepository cria uma nova instância de um MockCustomerRepository.
func NewMockProductRepository() *MockProductRepository {
	return &MockProductRepository{
		products: make(map[string]*entity.Product),
	}
}

// Create simula a criação de um novo cliente no repositório.
func (repo *MockProductRepository) Create(ctx context.Context, product *entity.Product) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	if _, exists := repo.products[product.ID]; exists {
		return errors.New("customer already exists")
	}

	repo.products[product.ID] = product
	return nil
}

// Update simula a atualização de um cliente no repositório.
func (repo *MockProductRepository) Update(ctx context.Context, product *entity.Product) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	if _, exists := repo.products[product.ID]; !exists {
		return errors.New("customer not found")
	}

	repo.products[product.ID] = product
	return nil
}

// Find simula a recuperação de um cliente pelo ID.
func (repo *MockProductRepository) Find(ctx context.Context, id string) (*entity.Product, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	customer, exists := repo.products[id]
	if !exists {
		return nil, errors.New("customer not found")
	}
	return customer, nil
}

// FindAll simula a recuperação de uma lista de clientes.
func (repo *MockProductRepository) FindAll(ctx context.Context) ([]*entity.Product, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	var products []*entity.Product
	for _, product := range repo.products {
		products = append(products, product)
	}
	return products, nil
}

// delete
func (repo *MockProductRepository) Delete(ctx context.Context, id string) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	if _, exists := repo.products[id]; !exists {
		return errors.New("product not found")
	}

	delete(repo.products, id)
	return nil
}

// Find simula a recuperação de um cliente pelo ID.
func (repo *MockProductRepository) FindByName(ctx context.Context, name string) (*entity.Product, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	for _, product := range repo.products {
		if product.Name == name {
			return product, nil
		}
	}

	return nil, errors.New("product not found")
}

func (repo *MockProductRepository) FindByCategory(ctx context.Context, category string) ([]*entity.Product, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	var products []*entity.Product

	for _, product := range repo.products {
		if product.Category == category {
			//return product, nil
			products = append(products, product)
		}
	}

	if len(products) > 0 {
		return products, nil
	}

	return nil, errors.New("product not found")
}
