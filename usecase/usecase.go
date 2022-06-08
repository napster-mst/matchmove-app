package usecase

import "matchmove.com/demo/dtos"

type AdminInterf interface {
	AdminLogin(dtos.Login) (string, error)
	CreateTokenForUser(id string) (string, error)
	CheckTokenStausesAdmin() (dtos.TokenStatusResponse, error)
	AdminDisableTokens() (bool, error)
}

type UserInterf interface {
	ValidateToken(dtos.TokenValidator) (bool, error)
}
