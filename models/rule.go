package models

type Route struct {
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
}

type Rule struct {
	Routes      []Route  `json:"routes"`
	Airlines    []string `json:"airlines"`
	Agencies    []string `json:"agencies"`
	Suppliers   []string `json:"suppliers"`
	AmountType  string   `json:"amountType"`
	AmountValue float64  `json:"amountValue"`
}

type JSON []byte

type Report struct {
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
}

type TempRule struct {
	Id   int  `json:"id" gorm:"primaryKey;autoIncrement"`
	File JSON `json:"file" gorm:"type:JSON;not null"`
}
