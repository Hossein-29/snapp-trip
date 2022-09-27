package models

type Route struct {
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
}

// type RuleDb struct {
// 	Id          int      `json:"id"`
// 	Routes      []Route  `json:"routes"`
// 	Airlines    []string `json:"airlines"`
// 	Agencies    []string `json:"agencies"`
// 	Suppliers   []string `json:"suppliers"`
// 	AmountType  string   `json:"amountType"`
// 	AmountValue float64  `json:"amountValue"`
// }

type Rule struct {
	Routes      []Route  `json:"routes"`
	Airlines    []string `json:"airlines"`
	Agencies    []string `json:"agencies"`
	Suppliers   []string `json:"suppliers"`
	AmountType  string   `json:"amountType"`
	AmountValue float64  `json:"amountValue"`
}

// type JSON []byte

type RuleResponse struct {
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
}

// type TempRule struct {
// 	Id   int  `json:"id" gorm:"primaryKey;autoIncrement"`
// 	File JSON `json:"file" gorm:"type:JSON;not null"`
// }

type RulesTable struct {
	Id          int `gorm:"primaryKey;autoIncrement"`
	AmountType  string
	AmountValue float64
}

type RoutesTable struct {
	Id     int `gorm:"primaryKey;autoIncrement"`
	Route  string
	RuleId int
	Rule   RulesTable `gorm:"foreignKey:RuleId;"`
}

type AirlinesTable struct {
	Id      int `gorm:"primaryKey;autoIncrement"`
	Airline string
	RuleId  int
	Rule    RulesTable `gorm:"foreignKey:RuleId;"`
}

type AgenciesTable struct {
	Id     int `gorm:"primaryKey;autoIncrement"`
	Agency string
	RuleId int
	Rule   RulesTable `gorm:"foreignKey:RuleId;"`
}

type SuppliersTable struct {
	Id       int `gorm:"primaryKey;autoIncrement"`
	Supplier string
	RuleId   int
	Rule     RulesTable `gorm:"foreignKey:RuleId;"`
}
