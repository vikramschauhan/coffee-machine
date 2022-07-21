package datastore

type RawMaterial struct {
	Name     string `gorm:"primary_key" json:"name"`
	Quantity string `gorm:"type:varchar(50)" json:"quantity"`
}

func (dataStore *DataStore) GetRawMaterials() ([]RawMaterial, error) {
	rawMaterials := []RawMaterial{}
	err := dataStore.client.Find(&rawMaterials).Error
	if err != nil {
		return nil, err
	}
	return rawMaterials, err
}

func (dataStore *DataStore) GetCoffeeTypes() ([]string, error) {
	var coffeeTypes []string
	err := dataStore.client.Table("ingredients").Pluck("coffeetype", &coffeeTypes).Error
	if err != nil {
		return nil, err
	}
	return coffeeTypes, err
}

func (RawMaterial) TableName() string {
	return "raw_material"
}
