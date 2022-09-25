package models

type ValidCityTable struct {
	Id   int `gorm:"primaryKey;autoIncrement"`
	Name string
}

type ValidAirlineTable struct {
	Id   int `gorm:"primaryKey;autoIncrement"`
	Name string
}

type ValidAgencyTable struct {
	Id   int `gorm:"primaryKey;autoIncrement"`
	Name string
}

type ValidSupplierTable struct {
	Id   int `gorm:"primaryKey;autoIncrement"`
	Name string
}
