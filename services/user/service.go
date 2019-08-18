package user

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pronuntio/core/domain/user"
	"go.uber.org/zap"
)

type UserService struct {
	dao user.Dao
	l   *zap.Logger
}

func NewUserService(dao user.Dao, l *zap.Logger) *UserService {
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
	users, err := us.dao.ListUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (us *UserService) GetUser(ID uint32) {

}
