package v0

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/kitanoyoru/effective-mobile-task/internal/requests"
)

func (api *HTTPApi) GetPersonRequestCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			getPersonRequest *requests.GetPersonRequest
		)

		if id := chi.URLParam(r, "id"); id != "" {
			pid, _ := strconv.Atoi(id)
			getPersonRequest.ID = pid
		}

		ctx := context.WithValue(r.Context(), "GetPersonRequest", getPersonRequest)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (api *HTTPApi) GetFilterPersonRequestCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			getFilterPersonRequest *requests.GetFilterPersonRequest
		)

		query := r.URL.Query()

		if id := query.Get("id"); id != "" {
			pid, _ := strconv.Atoi(id)
			getFilterPersonRequest.ID = pid
		}
		if name := query.Get("name"); name != "" {
			getFilterPersonRequest.Name = name
		}

		ctx := context.WithValue(r.Context(), "GetFilterPersonRequest", getFilterPersonRequest)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
