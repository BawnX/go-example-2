package packages

import (
	"github.com/gorilla/mux"
	"ms-2/config"
	"ms-2/packages/products"
	"ms-2/packages/shared/utils"
)

func NewPackages(router *mux.Router, validation *utils.Validation) *mux.Router {
	logProduct := config.NewLogger("Product", config.IsBuildMode())
	router = products.NewProductsRouter(router, logProduct, validation)
	return router
}
