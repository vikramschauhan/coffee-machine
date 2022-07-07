package datastore

type Employee struct {
	ID   int    `gorm:"primary_key" json:"id"`
	Name string `gorm:"type:varchar(50)" json:"name"`
}

func (dataStore *DataStore) GetAllEmployees() ([]Employee, error) {
	employees := []Employee{}
	err := dataStore.client.Find(&employees).Error
	if err != nil {
		return nil, err
	}
	return employees, err
}

func (Employee) TableName() string {
	return "employee"
}
