package middlewareUsecases

import "github.com/wiraphatys/shop-management-go/database"

type MiddlewareUsecase interface {
	VerifyToken(accessToken string) (*database.Admin, error)
}
