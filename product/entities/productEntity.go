package entities

type Product struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	UnitPrice   float64 `json:"unit_price"`
}
