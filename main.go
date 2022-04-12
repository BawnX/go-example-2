package main

import (
	"fmt"
	"github.com/BawnX/go-example-2/config"
	"github.com/BawnX/go-example-2/packages"
	"github.com/BawnX/go-example-2/packages/shared/utils"
	"github.com/gorilla/mux"
	"os"
)

func main() {
	config.NewEnvironment()

	logServer := config.NewLoggerStandard("Main", config.IsBuildMode())
	log := config.NewLogger("Main", config.IsBuildMode())

	validation := utils.NewValidation()

	cors := config.NewCorsConfig()
	serverRouter := mux.NewRouter()
	serverRouter = config.NewSwagger(serverRouter)
	serverRouter = packages.NewPackages(serverRouter, validation)

	server := config.NewHttpServerConfig(cors, serverRouter, logServer)
	go func() {
		log.Info(fmt.Sprint("Starting server on port ", *config.BindAddress))
		log.Info(fmt.Sprint("Trace level server ", *config.LogLevel))

		err := server.ListenAndServe()
		if err != nil {
			log.Error("Error starting server: \n", err)
			os.Exit(1)
		}
	}()

	config.CatchHttpServerError(log, &server)
}
