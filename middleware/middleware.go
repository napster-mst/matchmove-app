package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"matchmove.com/demo/common/alias"
	"matchmove.com/demo/common/util"
	"matchmove.com/demo/ent"
)

const MAX_USER_REQUEST_LIMIT = 5

type middlewareSVC struct {
	db orm.DB
}

type Middleware interface {
	AdminMiddleware(next http.Handler) http.Handler
	UserMiddleware(next http.Handler) http.Handler
}

func NewMiddleware(db orm.DB) Middleware {
	return &middlewareSVC{db: db}
}

func (ms *middlewareSVC) AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == "/admin/login" {
			next.ServeHTTP(w, r)
			return
		}

		t := r.Header.Get(alias.AccessToken)
		if t == "" {
			util.RespondWithError(w, http.StatusBadGateway, "no token provided")
			return
		}

		var user ent.User
		var token ent.Token

		ms.db.Model(&token).Where("token = ? AND status_id = (SELECT id FROM status WHERE name = ?)", t, string(alias.TokenStatusActive)).OrderExpr("expiry DESC").Limit(1).Select()
		if token.Token == "" || token.Expiry.Before(time.Now()) {
			util.RespondWithError(w, http.StatusBadGateway, "invalid/expired token")
			return
		}

		ms.db.Model(&user).Where("id = ? AND role_id = (SELECT id FROM roles WHERE name = ?)", token.UserId, string(alias.UserRoleAdmin)).Limit(1).Select()

		if user.Username == "" {
			util.RespondWithError(w, http.StatusForbidden, "don't have required permisson")
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (ms *middlewareSVC) UserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := r.Header.Get(alias.AccessToken)
		if t == "" {
			util.RespondWithError(w, http.StatusBadGateway, "no token provided")
			return
		}

		var user ent.User
		var token ent.Token

		ms.db.Model(&token).Where("token = ? AND status_id = (SELECT id FROM status WHERE name = ?)", t, string(alias.TokenStatusActive)).
			OrderExpr("expiry DESC").Limit(1).Select()
		if token.Token == "" || token.Expiry.Before(time.Now()) {
			util.RespondWithError(w, http.StatusBadGateway, "invalid/expired token")
			return
		}

		ms.db.Model(&user).Where("id = ?", token.UserId).Limit(1).Select()
		if user.Username == "" {
			util.RespondWithError(w, http.StatusForbidden, "don't have required permisson")
			return
		}

		if r.RequestURI == "/user/token/validate" {
			requestCount := ent.RequestCounts{}
			err := ms.db.Model(&requestCount).Where("token_id = ?", token.ID).Select()
			if err == pg.ErrNoRows {
				requestCount.TokenID = token.ID
				requestCount.Count = 1
				ms.db.Model(&requestCount).Insert()
			} else {
				requestCount.Count += 1
				if requestCount.Count > MAX_USER_REQUEST_LIMIT {
					util.RespondWithError(w, http.StatusForbidden, "max limit exceeded for this request")
					return
				}
				ms.db.Model(&requestCount).WherePK().Update()
			}
		}

		session := alias.Session("session")
		ctx := context.WithValue(r.Context(), session, user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
