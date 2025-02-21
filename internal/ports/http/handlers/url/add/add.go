package add

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
	UserID      int64
	Transaction transaction.Transaction
}

type Response struct {
	TransactionID int64
	response.Response
}

func New(log *slog.Logger, repo repo.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.add"

		log = log.With(slog.String("op", op))

		var request Request

		err := render.DecodeJSON(r.Body, &request)
		if err != nil {
			log.Error("failed to decode json", sl.Error(err))

			render.JSON(w, r, Response{0, response.Error("bad json")})

			return
		}

		transactionID, err := repo.Add(request.Transaction, request.UserID)
		if err != nil {
			log.Error("failed to add transaction", sl.Error(err))

			render.JSON(w, r, Response{0, response.Error("transaction error")})

			return
		}

		log.Info("transaction saved", slog.Attr{Key: "transactionid", Value: slog.Int64Value(transactionID)})

		render.JSON(w, r, response.Success())
	}
}
