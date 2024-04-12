package authUsecases

import (
	"fmt"
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

func (u *authUsecaseImpl) SignOut(token string) error {
	// clear token to make user log out
	// after this token will expired and cannot repeat again
	cfg := config.GetConfig()
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		// Verify the signing method
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(cfg.Jwt.Secret), nil
	})
	if err != nil {
		return err
	}
	return nil
}
