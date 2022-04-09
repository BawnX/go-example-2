package handlers

import (
	"ms-2/packages/shared/models"
	"ms-2/packages/shared/repositories"
	"ms-2/packages/shared/utils"
	"net/http"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Update a products details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  501: errorResponse

// Delete handles DELETE requests and removes items from the database
func (p *ProductHandler) Delete(rw http.ResponseWriter, r *http.Request) {
	id := getProductID(r)

	p.Log.Debug("deleting record id", id)

	err := repositories.DeleteProduct(id)
	if err == models.ErrProductNotFound {
		p.Log.Error("deleting record id does not exist")

		rw.WriteHeader(http.StatusNotFound)
		utils.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	if err != nil {
		p.Log.Error("deleting record", err)

		rw.WriteHeader(http.StatusInternalServerError)
		utils.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}
