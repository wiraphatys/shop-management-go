package adminRepositories

import "github.com/wiraphatys/shop-management-go/database"

type AdminRepository interface {
	FindAdminByEmail(email string) (*database.Admin, error)
	FindAdminById(a_id string) (*database.Admin, error)
	InsertAdmin(admin *database.Admin) (*database.Admin, error)
}
