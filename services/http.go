package services

import "github.com/gorilla/mux"

type IHTTPController interface {
	GetRoutes(*mux.Router) *mux.Router
}
