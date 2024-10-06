package controller

import (
	"context"
	"net/http"

	"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driver/api/dto"
	portsusecase "github.com/caiojorge/fiap-challenge-ddd/internal/core/application/ports/usecase/customer"
	"github.com/gin-gonic/gin"
)

type FindAllCustomersController struct {
	usecase portsusecase.FindAllCustomersUseCase
	ctx     context.Context
}

func NewFindAllCustomersController(ctx context.Context, usecase portsusecase.FindAllCustomersUseCase) *FindAllCustomersController {
	return &FindAllCustomersController{
		usecase: usecase,
		ctx:     ctx,
	}
}

// GetAllCustomers returns a list of all customers
// @Summary Get all customers
// @Description Get details of all customers
// @Tags Customers
// @Accept  json
// @Produce  json
// @Success 200 {array} dto.CustomerDTO
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /customers [get]
func (cr *FindAllCustomersController) GetAllCustomers(c *gin.Context) {

	customers, err := cr.usecase.FindAllCustomers(cr.ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	if len(customers) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No customers found"})
		return
	}

	var dtoCustomers []dto.CustomerDTO
	dto := dto.CustomerDTO{}
	for _, customer := range customers {
		dto.FromEntity(*customer)
		dtoCustomers = append(dtoCustomers, dto)
	}

	c.JSON(http.StatusOK, dtoCustomers)
}
