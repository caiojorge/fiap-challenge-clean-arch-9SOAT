package controller

import (
	"context"
	"errors"
	"net/http"

	portsusecase "github.com/caiojorge/fiap-challenge-ddd/internal/usecase/checkout/create"
	"github.com/gin-gonic/gin"
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
// @Param        request   body     usecase.CheckoutInputDTO  true  "cria novo Checkout"
// @Success 200 {object} usecase.CheckoutOutputDTO
// @Failure 400 {object} string "invalid data"
// @Failure 500 {object} string "internal server error"
// @Router /checkouts [post]
func (r *CreateCheckoutController) PostCreateCheckout(c *gin.Context) {
	var dto portsusecase.CheckoutInputDTO

	if err := c.BindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	output, err := r.usecase.CreateCheckout(r.ctx, &dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if output == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create transaction on gateway"})
		return
	}

	c.JSON(http.StatusOK, output.GatewayTransactionID)
}
