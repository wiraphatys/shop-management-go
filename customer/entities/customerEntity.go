package entities

type CustomerData struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
	City    string `json:"city"`
	Zip     string `json:"zip"`
}
