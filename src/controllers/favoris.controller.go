package controllers

import (
	"net/http"
	"projet/helper"
	"projet/services"
)

func DisplayFavorites(w http.ResponseWriter, r *http.Request) {
	names, _ := services.GetAllFavories()
	helper.RenderTemplate(w, r, "favoris", names)
}

func ToggleFavorite(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		services.ToggleFavoriteName(name)
	}
	http.Redirect(w, r, "/favoris", http.StatusSeeOther)
}
