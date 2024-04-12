package adminEntities

import "time"

type AdminData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AdminResponse struct {
	AID       string    `json:"aid"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
