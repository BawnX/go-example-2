package config

import (
	goHandlers "github.com/gorilla/handlers"
	"net/http"
)

func NewCorsConfig() func(http.Handler) http.Handler {
	return goHandlers.CORS(goHandlers.AllowedOrigins([]string{"*"}))
}
