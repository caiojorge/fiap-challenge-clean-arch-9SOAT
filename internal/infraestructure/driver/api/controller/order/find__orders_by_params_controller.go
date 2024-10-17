package controller

import (
	"context"
	"net/http"

	portsusecase "github.com/caiojorge/fiap-challenge-ddd/internal/core/usecase/order/findbyparam"
	"github.com/gin-gonic/gin"
)

type FindByParamsController struct {
	usecase portsusecase.FindOrderByParamsUseCase
	ctx     context.Context
}

func NewFindByParamsController(ctx context.Context, usecase portsusecase.FindOrderByParamsUseCase) *FindByParamsController {
	return &FindByParamsController{
		usecase: usecase,
		ctx:     ctx,
	}
}

// GetByParamsOrders returns a list of all paid orders
// @Summary Get all paid orders
// @Description Retorna todos os pedidos (orders) registrados no sistema. Se não houver pedidos, retorna um erro (404).
// @Tags Orders
// @Accept  json
// @Produce  json
// @Success 200 {array} usecase.OrderFindByParamOutputDTO
// @Failure 400 {object} string "Bad Request"
// @Failure 404 {object} string "Not Found"
// @Router /orders/paid [get]
func (r *FindByParamsController) GetByParamsOrders(c *gin.Context) {

	orders, err := r.usecase.FindOrdersByParams(r.ctx, map[string]interface{}{"status": "paid"})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	if len(orders) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No orders found"})
		return
	}

	c.JSON(http.StatusOK, orders)
}
