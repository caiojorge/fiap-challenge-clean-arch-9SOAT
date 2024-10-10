package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductWithError1(t *testing.T) {
	// nome inválido
	model := Product{ID: "", Name: "J", Description: "Pão queijo e carne", Price: 30.00, Category: "Lanche"}
	assert.NotNil(t, model.Validate())
}

func TestProductWithError2(t *testing.T) {
	// email inválido
	model := Product{ID: "1", Name: "", Description: "", Price: 0, Category: "Lanche"}

	assert.NotNil(t, model.Validate())
}

func TestProductWithNoError(t *testing.T) {
	model := Product{ID: "1", Name: "J", Description: "Pão queijo e carne", Price: 30.00, Category: "Lanche"}
	assert.Nil(t, model.Validate())
}
