package templatehelpers

import (
	"fmt"
	"html/template"

	"github.com/keto-granola/server/internal/product"
	"github.com/keto-granola/server/internal/webassets"
)

func FuncMap(assetsLoader *webassets.Loader) template.FuncMap {
	return template.FuncMap{
		"ingredientList": ingredientList,
		"centsToPrice":   centsToPrice,
		"asset":          assetsLoader.Asset,
		"assetCSS":       assetsLoader.AssetCSS,
	}
}

func ingredientList(ingredients []product.Ingredient) string {
	// TODO: implement
	// desired output: ing1 %80 (sub1, sub2, sub3), ing2 5%

	return ""
}

const centsPerUnit = 100

func centsToPrice(cents int32) string {
	return fmt.Sprintf("%.2f", float32(cents)/centsPerUnit)
}
