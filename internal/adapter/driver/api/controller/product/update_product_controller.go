package controller

import (
	"context"
	"net/http"

	"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driver/api/dto"
	portsusecase "github.com/caiojorge/fiap-challenge-ddd/internal/core/application/usecase/product"
	"github.com/gin-gonic/gin"
)

type UpdateProductController struct {
	usecase portsusecase.UpdateProductUseCase
	ctx     context.Context
}

func NewUpdateProductController(ctx context.Context, usecase portsusecase.UpdateProductUseCase) *UpdateProductController {
	return &UpdateProductController{
		usecase: usecase,
		ctx:     ctx,
	}
}

// PutUpdateProduct updates a Product by id
// @Summary Update a Product
// @Description Update details of a Product by id
// @Tags Products
// @Accept  json
// @Produce  json
// @Param id path string true "Product id"
// @Param Product body dto.CreateProductDTO true "Product data"
// @Success 200 {object} dto.ProductDTO
// @Failure 400 {object} string "Invalid data"
// @Failure 404 {object} string "Product not found"
// @Router /products/{id} [put]
func (r *UpdateProductController) PutUpdateProduct(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	var dto dto.ProductDTO

	if err := c.BindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	if dto.ID == "" || dto.ID != id {
		dto.ID = id
	}

	product, err := dto.ToEntity()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	err = r.usecase.UpdateProduct(r.ctx, *product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto)
}
