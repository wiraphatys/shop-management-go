package adminUsecases

import "github.com/wiraphatys/shop-management-go/database"

type AdminUsecase interface {
	GetAdminByEmail(email string) (*database.Admin, error)
	CreateAdmin(admin *database.Admin) (*database.Admin, error)
}
