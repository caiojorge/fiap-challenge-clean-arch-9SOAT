package controller

import (
	"context"
	"errors"
	"net/http"

	"github.com/caiojorge/fiap-challenge-ddd/internal/adapter/driver/api/dto"
	portsusecase "github.com/caiojorge/fiap-challenge-ddd/internal/core/application/ports/usecase/checkout"
	"github.com/caiojorge/fiap-challenge-ddd/internal/core/domain/entity"
	"github.com/caiojorge/fiap-challenge-ddd/internal/shared/formatter"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

var ErrAlreadyExists = errors.New("order already exists")

type CreateCheckoutController struct {
	ctx     context.Context
	usecase portsusecase.CreateCheckoutUseCase
}

func NewCreateCheckoutController(ctx context.Context,
	usecase portsusecase.CreateCheckoutUseCase) *CreateCheckoutController {
	return &CreateCheckoutController{
		usecase: usecase,
		ctx:     ctx,
	}
}

// PostCreateCheckout godoc
// @Summary Create Checkout
// @Schemes
// @Description Efetiva o pagamento do cliente, via fake checkout nesse momento, e libera o pedido para preparação. A ordem muda de status nesse momento, para em preparação.
// @Tags Checkouts
// @Accept json
// @Produce json
// @Param        request   body     dto.CreateCheckoutDTO  true  "cria novo Checkout"
// @Success 200 {object} dto.CheckoutDTO
// @Failure 400 {object} string "invalid data"
// @Failure 500 {object} string "internal server error"
// @Router /checkouts [post]
func (r *CreateCheckoutController) PostCreateCheckout(c *gin.Context) {
	var dto dto.CreateCheckoutDTO

	if err := c.BindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	dto.CustomerCPF = formatter.RemoveMaksFromCPF(dto.CustomerCPF)

	var entity entity.Checkout
	err := copier.Copy(&entity, &dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transactionID, err := r.usecase.CreateCheckout(r.ctx, &entity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if transactionID == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create transaction on gateway"})
		return
	}

	c.JSON(http.StatusOK, entity)
}
