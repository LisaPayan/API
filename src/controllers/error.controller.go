package controllers

import (
	"net/http"
	"projet/helper"
)

func ErrorDisplay(w http.ResponseWriter, r *http.Request) {

	type Error struct {
		Code    string
		Message string
	}

	data := Error{
		Code:    r.FormValue("code"),
		Message: r.FormValue("message"),
	}

	helper.RenderTemplate(w, r, "error", data)
}
