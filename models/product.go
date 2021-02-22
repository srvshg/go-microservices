package models

//Product is a type
type Product struct {
	ID        int     `json:"id,omitempty"`
	Name      string  `json:"name"`
	UnitPrice float64 `json:"price"`
}
