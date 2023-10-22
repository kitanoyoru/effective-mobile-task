package v0

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kitanoyoru/effective-mobile-task/internal/dtos"
)

func GetPersonDTOCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			personDto *dtos.PersonGetDTO
		)

		if id := chi.URLParam(r, "id"); id != "" {
			personDto.ID = &id
		}
		if name := chi.URLParam(r, "name"); name != "" {
			personDto.Name = &name
			personDto.WithFilter = true
		}

		ctx := context.WithValue(r.Context(), "personDTO", personDto)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
