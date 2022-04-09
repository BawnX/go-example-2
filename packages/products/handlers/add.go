package handlers

import (
	"ms-2/packages/shared/models"
	"ms-2/packages/shared/repositories"
	"ms-2/packages/shared/utils"
	"net/http"
)

// swagger:route POST /products products createProduct
// Create a new product
//
// responses:
//	200: productResponse
//  422: errorValidation
//  501: errorResponse

// Create handles POST requests to add new products
func (p *ProductHandler) Create(rw http.ResponseWriter, r *http.Request) {
	// fetch the product from the context
	prod := r.Context().Value(KeyProduct{}).(*models.ProductModel)

	p.Log.Debug("Inserting product", prod)
	returnProd := repositories.AddProduct(prod)
	utils.ToJSON(returnProd, rw)
}
