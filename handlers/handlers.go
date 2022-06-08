package handlers

import (
	"net/http"

	"github.com/go-pg/pg/orm"
	"github.com/gorilla/mux"
	"matchmove.com/demo/middleware"
	"matchmove.com/demo/repository/repositoryimpl"
	"matchmove.com/demo/usecase/usecaseimpl"
)

type Router struct {
	*mux.Router
}

func getController(db orm.DB) Controller {
	repo := repositoryimpl.NewRepository(db)
	svcAdmin := usecaseimpl.NewServiceAdmin(repo)
	svcUser := usecaseimpl.NewServiceUser(repo)

	controller := NewController(svcAdmin, svcUser)

	return controller
}

func Routes(port string, db orm.DB) {
	c := getController(db)
	r := mux.NewRouter().StrictSlash(true)
	var admin, user Router

	admin = Router{r.PathPrefix("/admin").Subrouter()}
	user = Router{r.PathPrefix("/user").Subrouter()}
	middlewareSVC := middleware.NewMiddleware(db)

	admin.HandleFunc("/login", c.loginHandler).Methods(http.MethodPost)

	admin.Use(middlewareSVC.AdminMiddleware)
	user.Use(middlewareSVC.UserMiddleware)

	admin.HandleFunc("/token/{user_id}", c.createToken).Methods(http.MethodGet)
	admin.HandleFunc("/users", c.checkUsersToken).Methods(http.MethodGet)
	admin.HandleFunc("/tokens/disable", c.tokenRecalling).Methods(http.MethodPost)
	user.HandleFunc("/token/validate", c.validateToken).Methods(http.MethodPost)

	http.ListenAndServe(":"+port, r)
}
