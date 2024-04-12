package authUsecases

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/wiraphatys/shop-management-go/admin/adminEntities"
	"github.com/wiraphatys/shop-management-go/admin/adminRepositories"
	"github.com/wiraphatys/shop-management-go/config"
	"golang.org/x/crypto/bcrypt"
)

type authUsecaseImpl struct {
	adminRepository adminRepositories.AdminRepository
}

func NewAuthUsecase(adminRepository adminRepositories.AdminRepository) AuthUsecase {
	return &authUsecaseImpl{
		adminRepository: adminRepository,
	}
}

func (u *authUsecaseImpl) SignIn(adminData *adminEntities.AdminData) (string, error) {
	existedAdmin, err := u.adminRepository.FindAdminByEmail(adminData.Email)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existedAdmin.Password), []byte(adminData.Password)); err != nil {
		return "", err
	}

	// create jwt token
	cfg := config.GetConfig()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  existedAdmin.AID,
		"exp": time.Now().Add(time.Second * time.Duration(cfg.Jwt.Expiration)).Unix(),
	})

	tokenString, err := token.SignedString([]byte(cfg.Jwt.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
