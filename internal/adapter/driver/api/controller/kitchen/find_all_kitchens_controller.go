package controller

import (
	"context"
	"net/http"

	"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driver/api/dto"
	portsusecase "github.com/caiojorge/fiap-challenge-ddd/internal/core/application/ports/usecase/kitchen"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type FindKitchenAllController struct {
	usecase portsusecase.FindAllKitchenUseCase
	ctx     context.Context
}

func NewFindKitchenAllController(ctx context.Context, usecase portsusecase.FindAllKitchenUseCase) *FindKitchenAllController {
	return &FindKitchenAllController{
		usecase: usecase,
		ctx:     ctx,
	}
}

// GetAllOrders returns a list of all orders in the kitchen
// @Summary Get all orders in the kitchen
// @Description Retorna todos os pedidos (orders) que estão na cozinha para inicio de preparação. Se não houver pedidos, retorna um erro (404).
// @Tags Kitchens
// @Accept  json
// @Produce  json
// @Success 200 {array} dto.KitchenDTO
// @Failure 400 {object} string "Bad Request"
// @Failure 404 {object} string "Not Found"
// @Router /kitchens/orders [get]
func (r *FindKitchenAllController) GetAllOrdersInKitchen(c *gin.Context) {

	entities, err := r.usecase.FindAllKitchen(r.ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	if len(entities) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No orders found"})
		return
	}

	var dtos []dto.KitchenDTO

	copier.Copy(&dtos, &entities)

	c.JSON(http.StatusOK, dtos)
}
