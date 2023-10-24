package v0

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/kitanoyoru/effective-mobile-task/internal/requests"
	"github.com/kitanoyoru/effective-mobile-task/pkg/utils"
)

func (api *HTTPApi) GetPersonRequestCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			getPersonRequest requests.GetPersonRequest
		)

		if id := chi.URLParam(r, "person_id"); id != "" {
			pid, _ := strconv.Atoi(id)
			getPersonRequest.ID = pid
		}

		ctx := context.WithValue(r.Context(), "GetPersonRequest", &getPersonRequest)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (api *HTTPApi) GetFilterPersonRequestCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			getFilterPersonRequest requests.GetFilterPersonRequest
		)

		query := r.URL.Query()

		if id := query.Get("id"); id != "" {
			pid, _ := strconv.Atoi(id)
			getFilterPersonRequest.ID = pid
		}
		if name := query.Get("name"); name != "" {
			getFilterPersonRequest.Name = name
		}

		ctx := context.WithValue(r.Context(), "GetFilterPersonRequest", &getFilterPersonRequest)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (api *HTTPApi) PostPersonRequestCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			postPersonRequest requests.PostPersonRequest
		)

		if err := json.NewDecoder(r.Body).Decode(&postPersonRequest); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		ctx := context.WithValue(r.Context(), "PostPersonRequest", &postPersonRequest)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (api *HTTPApi) DeletePersonRequestCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			deletePersonRequest requests.DeletePersonRequest
		)

		if err := json.NewDecoder(r.Body).Decode(&deletePersonRequest); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		ctx := context.WithValue(r.Context(), "DeletePersonRequest", &deletePersonRequest)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
