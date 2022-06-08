package repository

import (
	"matchmove.com/demo/common/alias"
	"matchmove.com/demo/dtos"
	"matchmove.com/demo/ent"
)

type Repository interface {
	InsertToken(token *ent.Token) error
	DisableAllUserTokens() error
	GetAdminCredentials(username string) (ent.User, error)
	GetStatus(status alias.TokenStatus) int
	CheckUserExist(id string) (bool, error)
	GetToken(token string) (t ent.Token, err error)
	CountTokenStatus() (tsr dtos.TokenStatusResponse, err error)
}
