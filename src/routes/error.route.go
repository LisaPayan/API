package routes

import (
	"net/http"
	"projet/controllers"
)

func errorRouter(router *http.ServeMux) {
	router.HandleFunc("/error", controllers.ErrorDisplay)
}
