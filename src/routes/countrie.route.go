package routes

import (
	"net/http"
	"projet/controllers"
)

func countriesRoutes(router *http.ServeMux) {
	router.HandleFunc("/countries", controllers.DisplayCountries)

	router.HandleFunc("/countries/search", controllers.DisplaySearch)

	router.HandleFunc("/filter", controllers.DisplayFilter)

	router.HandleFunc("/countries/pagination", controllers.DisplayPagination)

	router.HandleFunc("/countries/details", controllers.DisplayCountrieDetails)

	router.HandleFunc("/favoris", controllers.DisplayFavorites)

	router.HandleFunc("/favorites/toggle", controllers.ToggleFavorite)

	router.HandleFunc("/favoris/add", controllers.DisplayFavorites)

}
