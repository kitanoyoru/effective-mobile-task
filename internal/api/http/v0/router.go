package v0

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/kitanoyoru/effective-mobile-task/internal/service"
)

const (
	apiPrefix = "/api/v0"

	apiBaseRoutesPrefix   = "/"
	apiPersonRoutesPrefix = "/person"
)

type HTTPApi struct {
	service *service.Service
}

func NewHTTPApi(service *service.Service) *HTTPApi {
	return &HTTPApi{
		service,
	}
}

func (api *HTTPApi) GetHTTPRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}))

	r.Route(apiPrefix, func(r chi.Router) {
		r.Route(apiBaseRoutesPrefix, func(r chi.Router) {
			r.Get("/", api.renderIndexPage)
			r.Get("/version", api.renderAppVersionPage)
		})
		r.Route(apiPersonRoutesPrefix, func(r chi.Router) {
			r.With(api.GetPersonRequestCtx).Get("/{person_id}", api.getPersonRequestHandler)
			r.With(api.GetFilterPersonRequestCtx).Get("/", api.getFilterPersonRequestHandler)
			r.With(api.PostPersonRequestCtx).Post("/", api.postPersonRequestHandler)
			r.With(api.PatchPersonRequestCtx).Patch("/{person_id}", api.patchPersonRequestHandler)
			r.With(api.DeletePersonRequestCtx).Delete("/{person_id}", api.deletePersonRequestHanndler)
		})
	})

	return r
}
