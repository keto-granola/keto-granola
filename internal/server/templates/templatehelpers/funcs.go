package templatehelpers

import (
	"html/template"

	"github.com/keto-granola/server/internal/product"
)

func FuncMap() template.FuncMap {
	return template.FuncMap{
		"ingredientList": ingredientList,
	}
}

func ingredientList(ingredients []product.Ingredient) string {
	// TODO: implement
	// desired output: ing1 %80 (sub1, sub2, sub3), ing2 5%

	return ""
}
