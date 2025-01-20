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
	categoryWithS := ToPlural(category)
	upperStr := strings.ToUpper(categoryWithS)

	for i, c := range AllCategories {
		if c == upperStr {
			return i
		}
	}
	return -1
}

// IsCategoryValid checks if the category is valid
func IsCategoryValid(category string) bool {
	categoryWithS := ToPlural(category)
	upperStr := strings.ToUpper(categoryWithS)

	for _, c := range AllCategories {
		if c == upperStr {
			return true
		}
	}
	return false
}

func ToPlural(category string) string {

	upperStr := strings.ToUpper(category)
	if len(upperStr) > 0 && upperStr[len(upperStr)-1] != 'S' {
		upperStr += "S"
	}
	return upperStr
}
