package svc

import (
	"go-rest-api/datastore"
)

func (service *Services) GetAllEmployees() ([]datastore.Employee, error) {
	employees, err := service.dataStore.GetAllEmployees()

	if err != nil {
		return nil, err
	}
	return employees, err
}
