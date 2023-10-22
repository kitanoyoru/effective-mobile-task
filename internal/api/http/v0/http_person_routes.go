package v0

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/kitanoyoru/effective-mobile-task/internal/dtos"
	"github.com/kitanoyoru/effective-mobile-task/internal/service"
)

func getPersonRequestHandler(w http.ResponseWriter, r *http.Request) {
	getPersonDto := r.Context().Value("personDTO").(*dtos.PersonGetDTO)

	resp, err := service.GetPersonResponse(getPersonDto)
	if err != nil {
		render.Render(w, r, ErrRender(err))
	}

	render.Render(w, r, resp)

}

func getListPersonRequestHandler(w http.ResponseWriter, r *http.Request) {}

func postPersonRequestHandler(w http.ResponseWriter, r *http.Request) {}

func patchPersonRequetHandler(w http.ResponseWriter, r *http.Request) {}

func deletePersonRequestHanndler(w http.ResponseWriter, r *http.Request) {}
