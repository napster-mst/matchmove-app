package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"matchmove.com/demo/common/util"
	"matchmove.com/demo/dtos"
	"matchmove.com/demo/usecase"
)

type Controller struct {
	AdminCase usecase.AdminInterf
	UserCase  usecase.UserInterf
}

func NewController(adminCase usecase.AdminInterf, userCase usecase.UserInterf) Controller {
	return Controller{
		AdminCase: adminCase,
		UserCase:  userCase,
	}
}

func (c *Controller) loginHandler(w http.ResponseWriter, r *http.Request) {
	var req dtos.Login
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	token, err := c.AdminCase.AdminLogin(req)
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			util.RespondWithError(w, http.StatusForbidden, "invalid password")
			return
		}
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusAccepted, token)
}

func (c *Controller) createToken(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["user_id"]

	if userID == "" {
		util.RespondWithError(w, http.StatusNotFound, "user_id invalid")
		return
	}

	userToken, err := c.AdminCase.CreateTokenForUser(userID)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusOK, map[string]string{"token": userToken})
}

func (c *Controller) validateToken(w http.ResponseWriter, r *http.Request) {
	var req dtos.TokenValidator

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.Token == "" {
		util.RespondWithError(w, http.StatusBadRequest, "invalid token provided to check validation")
		return
	}

	if req.Token == "" {
		util.RespondWithError(w, http.StatusBadRequest, "no token provided to check")
		return
	}

	validate, err := c.UserCase.ValidateToken(req)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if !validate {
		util.RespondWithJSON(w, http.StatusOK, map[string]string{"validity": "token is invalid"})
		return
	}

	util.RespondWithJSON(w, http.StatusOK, map[string]string{"validity": "token is valid"})
}

func (c *Controller) checkUsersToken(w http.ResponseWriter, r *http.Request) {
	resp, err := c.AdminCase.CheckTokenStausesAdmin()
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(w, http.StatusOK, resp)
}

func (c *Controller) tokenRecalling(w http.ResponseWriter, r *http.Request) {
	success, err := c.AdminCase.AdminDisableTokens()
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if !success {
		util.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "unsuccess"})
		return
	}

	util.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "success"})
}
