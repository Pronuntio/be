package service

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pronuntio/core/pkg/word"
	"go.uber.org/zap"
)

type WordService struct {
	dao word.Dao
	l   *zap.Logger
}

func NewWordService(dao word.Dao, l *zap.Logger) *WordService {
	return &WordService{
		dao: dao,
		l:   l,
	}
}

func (ws *WordService) GetRoutes(router *mux.Router) *mux.Router {
	r := router.PathPrefix("/words").Subrouter()

	r.HandleFunc("/", ws.ListWords).Methods("GET")
	return r
}

func (ws *WordService) ListWords(w http.ResponseWriter, r *http.Request) {
	words, err := ws.dao.ListWords()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(words)
}
