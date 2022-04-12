package handlers

import (
	"context"
	"github.com/BawnX/go-example-2/packages/shared/models"
	"github.com/BawnX/go-example-2/packages/shared/utils"
	"net/http"
)

// MiddlewareValidate validates the product in the request and calls next if ok
func (p *ProductHandler) MiddlewareValidate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := &models.ProductModel{}

		err := utils.FromJSON(prod, r.Body)
		if err != nil {
			p.Log.Error("deserializing product", err)

			rw.WriteHeader(http.StatusBadRequest)
			utils.ToJSON(&GenericError{Message: err.Error()}, rw)
			return
		}

		// validate the product
		errs := p.Validation.Validate(prod)
		if len(errs) != 0 {
			p.Log.Error("validating product", errs)

			// return the validation messages as an array
			rw.WriteHeader(http.StatusUnprocessableEntity)
			utils.ToJSON(&ValidationError{Messages: errs.Errors()}, rw)
			return
		}

		// add the product to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}
