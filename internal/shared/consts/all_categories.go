package sharedconsts

import "strings"

const (
	Snacks     = "LANCHES"
	Drinks     = "BEBIDAS"
	SideDishes = "ACOMPANHAMENTOS"
	Desserts   = "SOBREMESAS"
)

var AllCategories = []string{
	Snacks,
	Drinks,
	SideDishes,
	Desserts,
}

// GetCategoryIndex returns the index of the category in the AllCategories slice
func GetCategoryIndex(category string) int {
	for i, c := range AllCategories {
		if c == category {
			return i
		}
	}
	return -1
}

// IsCategoryValid checks if the category is valid
func IsCategoryValid(category string) bool {
	upperStr := strings.ToUpper(category)

	for _, c := range AllCategories {
		if c == upperStr {
			return true
		}
	}
	return false
}
