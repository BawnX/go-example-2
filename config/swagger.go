package config

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

func NewSwagger(router *mux.Router) *mux.Router {
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	getRouter := router.Methods(http.MethodGet).Subrouter()

	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	return router
}
