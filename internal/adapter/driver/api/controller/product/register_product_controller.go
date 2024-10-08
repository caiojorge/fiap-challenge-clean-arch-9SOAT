package controller

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	portsusecase "github.com/caiojorge/fiap-challenge-ddd/internal/core/application/usecase/product/register"
	"github.com/gin-gonic/gin"
)

var ErrAlreadyExists = errors.New("product already exists")

type RegisterProductController struct {
	usecase portsusecase.RegisterProductUseCase
	ctx     context.Context
}

func NewRegisterProductController(ctx context.Context, usecase portsusecase.RegisterProductUseCase) *RegisterProductController {
	return &RegisterProductController{
		usecase: usecase,
		ctx:     ctx,
	}
}

// PostRegisterProduct godoc
// @Summary Create Product
// @Schemes
// @Description Create Product in DB
// @Tags Products
// @Accept json
// @Produce json
// @Param        request   body     portsusecase.RegisterProductInputDTO  true  "cria novo produto"
// @Success 200 {object} usecase.RegisterProductOutputDTO
// @Failure 400 {object} string "invalid data"
// @Failure 409 {object} string "product already exists"
// @Failure 500 {object} string "internal server error"
// @Router /products [post]
func (r *RegisterProductController) PostRegisterProduct(c *gin.Context) {
	var input portsusecase.RegisterProductInputDTO

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// Nesse cenário, o ID informado será ignorado e um novo ID será gerado
	fmt.Println("controller: Criando product: " + input.Name)
	output, err := r.usecase.RegisterProduct(r.ctx, &input)
	if err != nil {
		if err == ErrAlreadyExists {
			c.JSON(http.StatusConflict, gin.H{"error": "product already exists"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	log.Println("Product created: ", output)

	c.JSON(http.StatusOK, output)
}
