package controller

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driver/api/dto"
	portsusecase "github.com/caiojorge/fiap-challenge-ddd/internal/core/application/ports/usecase/customer"
	"github.com/gin-gonic/gin"
)

var ErrCustomerAlreadyExists = errors.New("customer already exists")

type RegisterCustomerController struct {
	usecase portsusecase.RegisterCustomerUseCase
	ctx     context.Context
}

func NewRegisterCustomerController(ctx context.Context, usecase portsusecase.RegisterCustomerUseCase) *RegisterCustomerController {
	return &RegisterCustomerController{
		usecase: usecase,
		ctx:     ctx,
	}
}

// PostRegisterCustomer godoc
// @Summary Create Customer
// @Schemes
// @Description Create Customer in DB
// @Tags Customers
// @Accept json
// @Produce json
// @Param        request   body     dto.CustomerDTO  true  "cria novo cliente"
// @Success 200 {object} dto.CustomerDTO
// @Failure 400 {object} map[string]string "invalid data"
// @Failure 409 {object} map[string]string "customer already exists"
// @Failure 500 {object} map[string]string "internal server error"
// @Router /customers [post]
func (r *RegisterCustomerController) PostRegisterCustomer(c *gin.Context) {
	var dto dto.CustomerDTO

	if err := c.BindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	entity, err := dto.ToEntity()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	fmt.Println("controller: Criando cliente: " + dto.CPF)
	err = r.usecase.RegisterCustomer(r.ctx, *entity)
	if err != nil {
		if err == ErrCustomerAlreadyExists {
			c.JSON(http.StatusConflict, gin.H{"error": "Customer already exists"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Use the user object, e.g., save to database, etc.
	// gin.H{"status": "customer created " + dto.Name}
	c.JSON(http.StatusOK, dto)
}
