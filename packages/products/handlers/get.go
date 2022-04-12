package handlers

import (
	"github.com/BawnX/go-example-2/packages/shared/models"
	"github.com/BawnX/go-example-2/packages/shared/repositories"
	"github.com/BawnX/go-example-2/packages/shared/utils"
	"net/http"
)

// swagger:route GET /products products listProducts
// Return a list of products from the database
// responses:
//	200: productsResponse

// ListAll handles GET requests and returns all current products
func (p *ProductHandler) ListAll(rw http.ResponseWriter, r *http.Request) {
	p.Log.Debug("get all records")

	rw.Header().Add("Content-type", "application/json")
	prods := repositories.GetProducts()

	err := utils.ToJSON(prods, rw)
	if err != nil {
		// we should never be here but log the error just incase
		p.Log.Error("serializing product", err)
	}
}

// swagger:route GET /products/{id} products listSingleProduct
// Return a list of products from the database
// responses:
//	200: productResponse
//	404: errorResponse

// ListSingle handles GET requests
func (p *ProductHandler) ListSingle(rw http.ResponseWriter, r *http.Request) {
	id := getProductID(r)

	p.Log.Debug("get record id", id)

	prod, err := repositories.GetProductByID(id)

	switch err {
	case nil:

	case models.ErrProductNotFound:
		p.Log.Error("fetching product", err)

		rw.WriteHeader(http.StatusNotFound)
		utils.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	default:
		p.Log.Error("fetching product", err)

		rw.WriteHeader(http.StatusInternalServerError)
		utils.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	err = utils.ToJSON(prod, rw)
	if err != nil {
		// we should never be here but log the error just incase
		p.Log.Error("serializing product", err)
	}
}
