package adminRepositories

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/wiraphatys/shop-management-go/database"
	"gorm.io/gorm"
)

type adminRepositoryImpl struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &adminRepositoryImpl{
		db: db,
	}
}

func (r *adminRepositoryImpl) FindAdminByEmail(email string) (*database.Admin, error) {
	var admin database.Admin
	result := r.db.First(&admin, "email = ?", email)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, result.Error
		}
		log.Errorf("FindAdminByEmail: %v", result.Error)
		return nil, result.Error
	}
	return &admin, nil
}

func (r *adminRepositoryImpl) FindAdminById(a_id string) (*database.Admin, error) {
	var admin database.Admin
	result := r.db.First(&admin, "a_id = ?", a_id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, result.Error
		}
		log.Errorf("FindAdminById: %v", result.Error)
		return nil, result.Error
	}
	return &admin, nil
}

func (r *adminRepositoryImpl) InsertAdmin(admin *database.Admin) (*database.Admin, error) {
	// create new admin
	admin.AID = "1"
	result := r.db.Create(admin)
	if result.Error != nil {
		log.Errorf("InsertAdmin: %v", result.Error)
		return nil, result.Error
	}

	// returning created admin
	createdAdmin, err := r.FindAdminByEmail(admin.Email)
	if err != nil {
		return nil, err
	}

	return createdAdmin, nil
}
