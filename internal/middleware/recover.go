package middleware

import (
	"log/slog"
	"net/http"

	"movies-api/internal/dto"
	"movies-api/internal/response"
)

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				slog.ErrorContext(r.Context(), "panic recovered", "panic", err)
				response.WriteJSON(w, http.StatusInternalServerError, dto.InternalServerError("internal server error"))
			}
		}()

		next.ServeHTTP(w, r)
	})
}
