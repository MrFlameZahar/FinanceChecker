package router

import (
	"FinanceChecker/internal/ports/http/handlers/url/add"
	"FinanceChecker/internal/ports/http/handlers/url/get"
	"FinanceChecker/internal/repo"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRouter(log *slog.Logger, storage repo.Repository) http.Handler {
	transactionMux := mux.NewRouter()
	transactionMux.HandleFunc("/transaction", add.New(log, storage)).Methods("POST")
	transactionMux.HandleFunc("/transaction", get.New(log, storage)).Methods("GET")

	return transactionMux
}
