package controllers

import (
	"net/http"
	"projet/helper"
)

func DisplayMenu(w http.ResponseWriter, r *http.Request) {
	helper.RenderTemplate(w, r, "menu", nil)
}
