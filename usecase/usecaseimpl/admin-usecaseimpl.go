package usecaseimpl

import (
	"errors"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
	"matchmove.com/demo/common/alias"
	"matchmove.com/demo/common/util"
	"matchmove.com/demo/dtos"
	"matchmove.com/demo/ent"
	"matchmove.com/demo/repository"
	"matchmove.com/demo/usecase"
)

type adminService struct {
	repo repository.Repository
}

func NewServiceAdmin(repo repository.Repository) usecase.AdminInterf {
	return &adminService{
		repo: repo,
	}
}

func (svc *adminService) AdminLogin(login dtos.Login) (string, error) {
	user, err := svc.repo.GetAdminCredentials(login.Username)
	if err != nil {
		return "", err
	}

	if user.Password == "" {
		return "", errors.New("invalid username/password")
	}

	// Comparing the password with the hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
		return "", err
	}

	token := util.TokenGenerator()
	expiry := time.Now().Add(time.Minute * 60)
	tokenEnt := ent.Token{
		Token:    token,
		StatusId: svc.repo.GetStatus(alias.TokenStatusActive),
		UserId:   user.ID,
		Expiry:   expiry,
	}

	err = svc.repo.InsertToken(&tokenEnt)
	if err != nil {
		return "", err
	}

	return tokenEnt.Token, nil
}

func (svc *adminService) CreateTokenForUser(id string) (token string, err error) {
	ok, err := svc.repo.CheckUserExist(id)
	if err != nil {
		return
	}
	if !ok {
		err = errors.New("user not exists")
		return
	}

	userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return
	}

	token = util.TokenGenerator()
	expiry := time.Now().Add(time.Minute * 60)
	tokenEnt := ent.Token{
		Token:    token,
		StatusId: svc.repo.GetStatus(alias.TokenStatusActive),
		UserId:   int(userID),
		Expiry:   expiry,
	}

	err = svc.repo.InsertToken(&tokenEnt)

	return
}

func (svc *adminService) CheckTokenStausesAdmin() (tsr dtos.TokenStatusResponse, err error) {
	tsr, err = svc.repo.CountTokenStatus()
	return
}

func (svc *adminService) AdminDisableTokens() (bool, error) {
	err := svc.repo.DisableAllUserTokens()
	if err != nil {
		return false, err
	}

	return true, nil
}
