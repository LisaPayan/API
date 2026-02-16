package routes

import (
	"net/http"
	"projet/controllers"
)

func PrincipaleRoute(router *http.ServeMux) {
	router.HandleFunc("/", controllers.DisplayMenu)
}
