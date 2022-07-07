package main

import (
	"fmt"
	"go-rest-api/config"
	"go-rest-api/datastore"
	"go-rest-api/svc"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
)

var conf *config.AppConfig

func main() {
	conf = config.BuildConfig()
	logger := log.NewJSONLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)
	datastore, err := datastore.NewDataStore(&conf.DbConfig, logger)
	if err != nil {
		panic(err)
	}
	defer datastore.Close()
	router := mux.NewRouter()
	service := svc.NewServices(datastore, conf, logger)
	server := svc.NewServer(&conf.HTTPServer, logger, router, service)

	shutdown := make(chan error, 1)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	go func() {
		logger.Log("msg", "Starting server..", "bindAddr", *conf.HTTPServer.ListenAddr)
		err := server.Start()
		shutdown <- err
	}()
	select {
	case signalKill := <-interrupt:
		logger.Log("msg", fmt.Sprintf("Stopping Server: %s", signalKill.String()))
	case err := <-shutdown:
		logger.Log("error", err)
	}

	err = server.Shutdown()
	if err != nil {
		logger.Log("error", err)
	}
}
