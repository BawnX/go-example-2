package handlers

import (
	"fmt"
	"github.com/BawnX/go-example-2/packages/shared/utils"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"net/http"
	"strconv"
)

// KeyProduct is a key used for the Product object in the context
type KeyProduct struct{}

// ProductHandler for getting and updating products
type ProductHandler struct {
	utils.Handlers
}

// NewProductsHandlers returns a new products handler with the given logger
func NewProductsHandlers(l hclog.Logger, v *utils.Validation) *ProductHandler {
	return &ProductHandler{
		Handlers: utils.Handlers{
			Log:        l,
			Validation: v,
		},
	}
}

// ErrInvalidProductPath is an error message when the product path is not valid
var ErrInvalidProductPath = fmt.Errorf("Invalid Path, path should be /products/[id]")

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

// getProductID returns the product ID from the URL
// Panics if cannot convert the id into an integer
// this should never happen as the router ensures that
// this is a valid number
func getProductID(r *http.Request) int {
	// parse the product id from the url
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// should never happen
		panic(err)
	}

	return id
}
