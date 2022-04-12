package products

import (
	"github.com/BawnX/go-example-2/packages/products/handlers"
	"github.com/BawnX/go-example-2/packages/shared/utils"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"net/http"
)

func NewProductsRouter(router *mux.Router, log hclog.Logger, validation *utils.Validation) *mux.Router {
	productsHandlers := handlers.NewProductsHandlers(log, validation)

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/products", productsHandlers.ListAll)
	getRouter.HandleFunc("/products/{id:[0-9]+}", productsHandlers.ListSingle)

	putRouter := router.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/products", productsHandlers.Update)
	putRouter.Use(productsHandlers.MiddlewareValidate)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/products", productsHandlers.Create)
	postRouter.Use(productsHandlers.MiddlewareValidate)

	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/products/{id:[0-9]+}", productsHandlers.Delete)

	return router
}
