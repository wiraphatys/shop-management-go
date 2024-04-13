package middlewareUsecases

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/wiraphatys/shop-management-go/admin/adminRepositories"
	"github.com/wiraphatys/shop-management-go/config"
	"github.com/wiraphatys/shop-management-go/database"
)

type middlewareUsecaseImpl struct {
	adminRepository adminRepositories.AdminRepository
}

func NewMiddlewareUsecase(adminRepository adminRepositories.AdminRepository) MiddlewareUsecase {
	return &middlewareUsecaseImpl{
		adminRepository: adminRepository,
	}
}

func (u *middlewareUsecaseImpl) VerifyToken(accessToken string) (*database.Admin, error) {
	cfg := config.GetConfig()
	claims := &jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(accessToken, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(cfg.Jwt.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	a_id, ok := (*claims)["id"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	admin, err := u.adminRepository.FindAdminById(a_id)
	if err != nil {
		return nil, err
	}

	return admin, nil
}
