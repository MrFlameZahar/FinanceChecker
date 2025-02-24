package get

import (
	sl "FinanceChecker/internal/lib/log"
	response "FinanceChecker/internal/models"
	"FinanceChecker/internal/models/transaction"
	"FinanceChecker/internal/repo"
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
)

type Request struct {
	UserID          int64
	TransactionType string
}

type Response struct {
	Transactions []transaction.Transaction
	response.Response
}

func New(log *slog.Logger, storage repo.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.Get"

		log.With(slog.String("op", op))

		var request Request

		err := render.DecodeJSON(r.Body, &request)
		if err != nil {
			log.Error("cant decode json", sl.Error(err))

			render.JSON(w, r, Response{nil, response.Error("something wrong with json")})

			return
		}

		res, err := storage.Get(request.UserID, request.TransactionType)
		if err != nil {
			log.Error("cant get transaction data", sl.Error(err))

			render.JSON(w, r, Response{nil, response.Error("cant get transaction data")})

			return
		}

		log.Info("get transaction info success")

		render.JSON(w, r, Response{res, response.Success()})
	}
}
