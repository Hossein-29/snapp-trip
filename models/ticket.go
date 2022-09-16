package models

type Ticket struct {
	Origin       string  `json:"origin"`
	Destination  string  `json:"destination"`
	Airline      string  `json:"airline"`
	Agency       string  `json:"agency"`
	Supplier     string  `json:"supplier"`
	BasePrice    float64 `json:"baseprice"`
	Markup       float64 `json:"markup"`
	PayablePrice float64 `json:"payableprice"`
}
