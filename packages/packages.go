package packages

import (
	"github.com/BawnX/go-example-2/config"
	"github.com/BawnX/go-example-2/packages/products"
	"github.com/BawnX/go-example-2/packages/shared/database"
	"github.com/BawnX/go-example-2/packages/shared/utils"
	"github.com/gorilla/mux"
	"log"
)

func NewPackages(router *mux.Router, validation *utils.Validation) *mux.Router {
	db := database.NewDatabase(*config.ConnectionSql)

	var count int
	db.Raw("select count(*) from [emails].[Sender]").Scan(&count)
	log.Println(count)
	logProduct := config.NewLogger("Product", config.IsBuildMode())
	router = products.NewProductsRouter(router, logProduct, validation)
	return router
}
