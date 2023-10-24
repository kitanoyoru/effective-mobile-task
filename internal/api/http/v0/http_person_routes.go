package v0

import (
	"net/http"

	"github.com/kitanoyoru/effective-mobile-task/internal/requests"
	"github.com/kitanoyoru/effective-mobile-task/pkg/utils"
	"github.com/sirupsen/logrus"
)

func (api *HTTPApi) getPersonRequestHandler(w http.ResponseWriter, r *http.Request) {
	getPersonRequest := r.Context().Value("GetPersonRequest").(*requests.GetPersonRequest)

	logrus.Debug(getPersonRequest)

	resp, err := api.service.GetPerson(r.Context(), getPersonRequest)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, resp)
}

func (api *HTTPApi) getFilterPersonRequestHandler(w http.ResponseWriter, r *http.Request) {
	getFilterPersonRequest := r.Context().Value("GetFilterPersonRequest").(*requests.GetFilterPersonRequest)

	resp, err := api.service.FilterPerson(r.Context(), getFilterPersonRequest)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, resp)

}

func (api *HTTPApi) postPersonRequestHandler(w http.ResponseWriter, r *http.Request) {
	postPersonRequest := r.Context().Value("PostPersonRequest").(*requests.PostPersonRequest)

	resp, err := api.service.AddPerson(r.Context(), postPersonRequest)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, resp)
}

func (api *HTTPApi) patchPersonRequetHandler(w http.ResponseWriter, r *http.Request) {
}

func (api *HTTPApi) deletePersonRequestHanndler(w http.ResponseWriter, r *http.Request) {
	deletePersonRequest := r.Context().Value("DeletePersonRequest").(*requests.DeletePersonRequest)

	resp, err := api.service.DeletePerson(r.Context(), deletePersonRequest)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, resp)
}
