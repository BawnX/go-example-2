package config

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func NewHttpServerConfig(cors func(http.Handler) http.Handler, serverMethods *mux.Router, log *log.Logger) http.Server {
	return http.Server{
		Addr:         fmt.Sprintf(":%v", *BindAddress), // configure the bind address
		Handler:      cors(serverMethods),              // set the default handler
		ErrorLog:     log,                              // set the logger for the server
		ReadTimeout:  5 * time.Second,                  // max time to read request from the sdk
		WriteTimeout: 10 * time.Second,                 // max time to write response to the sdk
		IdleTimeout:  120 * time.Second,                // max time for connections using TCP Keep-Alive
	}
}

func CatchHttpServerError(log hclog.Logger, server *http.Server) {
	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Error("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	_ = server.Shutdown(ctx)
}
