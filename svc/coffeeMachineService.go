package svc

import (
	"go-rest-api/datastore"
)

func (service *Services) GetRawMaterials() ([]datastore.RawMaterial, error) {
	rawMaterials, err := service.dataStore.GetRawMaterials()

	if err != nil {
		return nil, err
	}
	return rawMaterials, err
}

func (service *Services) GetCoffeeTypes() ([]string, error) {
	coffeeTypes, err := service.dataStore.GetCoffeeTypes()

	if err != nil {
		return nil, err
	}
	return coffeeTypes, err
}
