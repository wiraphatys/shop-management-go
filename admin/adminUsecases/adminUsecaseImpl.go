package adminUsecases

import (
	"github.com/wiraphatys/shop-management-go/admin/adminRepositories"
	"github.com/wiraphatys/shop-management-go/database"
	"golang.org/x/crypto/bcrypt"
)

type adminUsecaseImpl struct {
	adminRepository adminRepositories.AdminRepository
}

func NewAdminUsecase(adminRepository adminRepositories.AdminRepository) AdminUsecase {
	return &adminUsecaseImpl{
		adminRepository: adminRepository,
	}
}

func (u *adminUsecaseImpl) GetAdminByEmail(email string) (*database.Admin, error) {
	result, err := u.adminRepository.FindAdminByEmail(email)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *adminUsecaseImpl) CreateAdmin(admin *database.Admin) (*database.Admin, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	admin.Password = string(hashedPassword)

	result, err := u.adminRepository.InsertAdmin(admin)
	if err != nil {
		return nil, err
	}

	return result, err
}
