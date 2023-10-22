package v0

import (
	"net/http"

	"github.com/kitanoyoru/effective-mobile-task/internal/requests"
	"github.com/kitanoyoru/effective-mobile-task/pkg/utils"
)

func (api *HTTPApi) getPersonRequestHandler(w http.ResponseWriter, r *http.Request) {
	getPersonRequet := r.Context().Value("GetPersonRequest").(*requests.GetPersonRequest)

	resp, err := api.service.GetPersonResponse(r.Context(), getPersonRequet)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondwithJSON(w, http.StatusOK, resp)
}

func (api *HTTPApi) getFilterPersonRequestHandler(w http.ResponseWriter, r *http.Request) {
	getFilterPersonRequest := r.Context().Value("GetFilterPersonRequest").(*requests.GetFilterPersonRequest)

	resp, err := api.service.GetFilterPersonResponse(r.Context(), getFilterPersonRequest)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondwithJSON(w, http.StatusOK, resp)

}

func (api *HTTPApi) postPersonRequestHandler(w http.ResponseWriter, r *http.Request) {}

func (api *HTTPApi) patchPersonRequetHandler(w http.ResponseWriter, r *http.Request) {}

func (api *HTTPApi) deletePersonRequestHanndler(w http.ResponseWriter, r *http.Request) {}
