package authUsecases

import "github.com/wiraphatys/shop-management-go/admin/adminEntities"

type AuthUsecase interface {
	SignIn(adminData *adminEntities.AdminData) (string, error)
}
