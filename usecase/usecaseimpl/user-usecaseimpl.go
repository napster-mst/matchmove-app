package usecaseimpl

import (
	"errors"
	"time"

	"github.com/go-pg/pg"
	"matchmove.com/demo/common/alias"
	"matchmove.com/demo/dtos"
	"matchmove.com/demo/repository"
	"matchmove.com/demo/usecase"
)

type userService struct {
	repo repository.Repository
}

func NewServiceUser(repo repository.Repository) usecase.UserInterf {
	return &userService{
		repo: repo,
	}
}

func (svc *userService) ValidateToken(token dtos.TokenValidator) (validate bool, err error) {
	t, err := svc.repo.GetToken(token.Token)
	if err != nil {
		if err == pg.ErrNoRows {
			err = errors.New("invalid token")
		}
		return
	}

	if t.Expiry.Before(time.Now()) || t.StatusId == svc.repo.GetStatus(alias.TokenStatusInactive) {
		validate = false
		return
	}

	validate = true
	return
}
