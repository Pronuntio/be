package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pronuntio/core/domain/user"
	"go.uber.org/zap"
)

type UserService struct {
	dao *user.Dao
	l   *zap.Logger
}

func (us *UserService) NewUserService(dao *user.Dao, l *zap.Logger) *UserService {
	return &UserService{
		dao: dao,
		l:   l,
	}
}

func (us *UserService) GetRoutes(router *mux.Router) *mux.Router {
	r := router.PathPrefix("/users").Subrouter()

	r.HandleFunc("/", us.ListUsers).Methods("GET")

	return r
}

func (us *UserService) ListUsers(w http.ResponseWriter, r *http.Request) {

}

func (us *UserService) GetUser(ID uint32) {

}
