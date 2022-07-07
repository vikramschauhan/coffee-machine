package datastore

import (
	"github.com/go-kit/kit/log"

	"go-rest-api/config"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DataStore struct {
	client gorm.DB
	logger log.Logger
}

func NewDataStore(config *config.DbConfig, logger log.Logger) (*DataStore, error) {
	client, err := gorm.Open(*config.Engine, *config.ConnString)
	if err != nil {
		return nil, err
	}
	client.DB().SetMaxOpenConns(*config.MaxOpenConns)
	client.DB().SetMaxIdleConns(*config.MaxIdleConns)

	dataStore := DataStore{
		client: *client,
		logger: logger,
	}
	err = dataStore.Migrate()
	if err != nil {
		return nil, err
	}
	return &dataStore, err
}

func (dataStore *DataStore) Close() error {
	return dataStore.client.Close()
}

func (dataStore *DataStore) Ping() error {
	return dataStore.client.DB().Ping()
}

func (dataStore *DataStore) Migrate() error {
	err := dataStore.client.AutoMigrate(&Employee{}).Error
	return err
}
