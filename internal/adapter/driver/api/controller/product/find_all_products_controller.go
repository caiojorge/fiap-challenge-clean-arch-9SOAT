package controller

import (
	"context"
	"net/http"

	"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driver/api/dto"
	portsusecase "github.com/caiojorge/fiap-challenge-ddd/internal/core/application/usecase/product"
	"github.com/gin-gonic/gin"
)

type FindAllProductController struct {
	usecase portsusecase.FindAllProductsUseCase
	ctx     context.Context
}

func NewFindAllProductController(ctx context.Context, usecase portsusecase.FindAllProductsUseCase) *FindAllProductController {
	return &FindAllProductController{
		usecase: usecase,
		ctx:     ctx,
	}
}

// GetAllProducts returns a list of all products
// @Summary Get all products
// @Description Get details of all products
// @Tags Products
// @Accept  json
// @Produce  json
// @Success 200 {array} dto.ProductDTO
// @Failure 400 {object} map[string]string "Invalida data"
// @Failure 404 {object} map[string]string "No products foundr"
// @Router /products [get]
func (cr *FindAllProductController) GetAllProducts(c *gin.Context) {

	entityProducts, err := cr.usecase.FindAllProducts(cr.ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	if len(entityProducts) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No products found"})
		return
	}

	var dtoproducts []dto.ProductDTO
	dto := dto.ProductDTO{}
	for _, entity := range entityProducts {
		dto.FromEntity(*entity)
		dtoproducts = append(dtoproducts, dto)
	}

	c.JSON(http.StatusOK, dtoproducts)
}
