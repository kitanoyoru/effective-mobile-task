package v0

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

const (
	apiPrefix = "/api/v0"

	apiBaseRoutesPrefix   = "/"
	apiPersonRoutesPrefix = "/person"
)

func InitHTTPRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route(apiPrefix, func(r chi.Router) {
		r.Route(apiBaseRoutesPrefix, func(r chi.Router) {
			r.Get("/", renderIndexPage)
			r.Get("/version", renderAppVersionPage)
		})
		r.Route(apiPersonRoutesPrefix, func(r chi.Router) {
			r.With(GetPersonDTOCtx).Get("/", getPersonRequestHandler)
			r.With(paginateMiddleware).Get("/list", getListPersonRequestHandler)
			r.Post("/", postPersonRequestHandler)
			r.Patch("/", patchPersonRequetHandler)
			r.Delete("/", deletePersonRequestHanndler)
		})
	})

	return r
}

func paginateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// just a stub.. some ideas are to look at URL query params for something like
		// the page number, or the limit, and send a query cursor down the chain
		next.ServeHTTP(w, r)
	})
}

type ErrResponse struct {
	Err            error `json:"-"`
	HTTPStatusCode int   `json:"-"`

	StatusText string `json:"status"`
	AppCode    int64  `json:"code,omitempty"`
	ErrorText  string `json:"error,omitempty"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}

func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 422,
		StatusText:     "Error rendering response.",
		ErrorText:      err.Error(),
	}
}
