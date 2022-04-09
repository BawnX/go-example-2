package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"ms-2/config"
	"ms-2/packages"
	"ms-2/packages/shared/utils"
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
