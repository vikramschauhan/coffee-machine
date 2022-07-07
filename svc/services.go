package svc

import (
	"go-rest-api/config"
	"go-rest-api/datastore"

	"github.com/go-kit/kit/log"
)

type Services struct {
	dataStore *datastore.DataStore
	logger    log.Logger
}

func NewServices(dataStore *datastore.DataStore, config *config.AppConfig, logger log.Logger) *Services {
	services := &Services{
		dataStore: dataStore,
		logger:    logger,
	}
	return services
}

func (services *Services) HealthCheck() (string, error) {
	err := services.dataStore.Ping()
	if err != nil {
		return "", err
	}
	return "DB is up and running", nil
}
