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
		} else {
			utils.RespondWithError(w, http.StatusBadRequest, "id must be specified and be a part of the path")
			return
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
		if surname := query.Get("surname"); surname != "" {
			getFilterPersonRequest.Surname = surname
		}
		if gender := query.Get("gender"); gender != "" {
			getFilterPersonRequest.Gender = gender
		}
		if countryID := query.Get("countryID"); countryID != "" {
			getFilterPersonRequest.CountryID = countryID
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

		if id := chi.URLParam(r, "person_id"); id != "" {
			pid, _ := strconv.Atoi(id)
			deletePersonRequest.ID = pid
		} else {
			utils.RespondWithError(w, http.StatusBadRequest, "id must be specified and be a part of the path")
			return
		}

		ctx := context.WithValue(r.Context(), "DeletePersonRequest", &deletePersonRequest)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (api *HTTPApi) PatchPersonRequestCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			patchPersonRequest requests.PatchPersonRequest
		)

		if id := chi.URLParam(r, "person_id"); id != "" {
			pid, _ := strconv.Atoi(id)
			patchPersonRequest.ID = pid
		} else {
			utils.RespondWithError(w, http.StatusBadRequest, "id must be specified and be a part of the path")
			return
		}

		if err := json.NewDecoder(r.Body).Decode(&patchPersonRequest); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		ctx := context.WithValue(r.Context(), "PatchPersonRequest", &patchPersonRequest)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
