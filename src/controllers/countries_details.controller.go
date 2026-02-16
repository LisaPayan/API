package controllers

import (
	"net/http"
	"projet/helper"
	"projet/services"
)

func DisplayCountrieDetails(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	data, statusCode, err := services.GetCountriesDetails(name)
	if err != nil {
		helper.RedirectToError(w, r, statusCode, err.Error())
		return
	}

	if statusCode != http.StatusOK {
		helper.RedirectToError(w, r, statusCode, "Erreur lors de la récupération des données")
		return
	}

	helper.RenderTemplate(w, r, "details-country", data)
}
