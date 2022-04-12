package handlers

import (
	"github.com/BawnX/go-example-2/packages/shared/models"
	"github.com/BawnX/go-example-2/packages/shared/repositories"
	"github.com/BawnX/go-example-2/packages/shared/utils"
	"net/http"
)

// swagger:route PUT /products products updateProduct
// Update a products details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  422: errorValidation

// Update handles PUT requests to update products
func (p *ProductHandler) Update(rw http.ResponseWriter, r *http.Request) {

	// fetch the product from the context
	prod := r.Context().Value(KeyProduct{}).(*models.ProductModel)
	p.Log.Debug("updating record id", prod.ID)

	err := repositories.UpdateProduct(prod)
	if err == models.ErrProductNotFound {
		p.Log.Error("product not found", err)

		rw.WriteHeader(http.StatusNotFound)
		utils.ToJSON(&GenericError{Message: "Product not found in database"}, rw)
		return
	}

	// write the no content success header
	rw.WriteHeader(http.StatusNoContent)
}
