package v0

import (
	"html/template"
	"net/http"

	"github.com/kitanoyoru/effective-mobile-task/pkg/utils"
)

func (api *HTTPApi) renderIndexPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("index.html").ParseFiles(api.service.GetTemplatePath("index.html"))
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithTemplate(w, tmpl, nil)
}
