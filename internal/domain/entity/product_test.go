package entity

import (
	"testing"

	"github.com/caiojorge/fiap-challenge-ddd/internal/shared"
	"github.com/stretchr/testify/assert"
)

func TestProduct(t *testing.T) {

	product, err := ConvertProduct(shared.NewIDGenerator(), "Lanche XPTO", "Pão queijo e carne", "Lanche", 30.00)
	assert.Nil(t, err)
	assert.NotNil(t, product)

	assert.Equal(t, "Lanche XPTO", product.Name)

	assert.NotEmpty(t, product.GetID())

	product2, err := NewProduct("Lanche XPTO", "Pão queijo e carne", "Lanche", 30.00)
	assert.Nil(t, err)
	assert.NotNil(t, product2)

	assert.Equal(t, "Lanche XPTO", product.Name)

	assert.NotEmpty(t, product.GetID())

}
