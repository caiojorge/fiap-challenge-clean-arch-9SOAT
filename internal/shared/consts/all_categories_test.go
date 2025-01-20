package sharedconsts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCategories(t *testing.T) {

	category := "LANCHE"

	isOK := IsCategoryValid(category)
	assert.Equal(t, true, isOK)

	category = "LANCHES"
	isOK = IsCategoryValid(category)
	assert.Equal(t, true, isOK)

	index := GetCategoryIndex(category)
	assert.Greater(t, index, -1)

	category = "PASTA"
	isOK = IsCategoryValid(category)
	assert.Equal(t, false, isOK)

	index = GetCategoryIndex(category)
	assert.Equal(t, index, -1)

}
