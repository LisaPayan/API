package controllers

import (
	"fmt"
	"log"
	"net/http"
	"projet/helper"
	"projet/services"
	"slices"
	"strconv"
	"strings"
)

func DisplayCountries(w http.ResponseWriter, r *http.Request) {
	data, statusCode, err := services.GetCountries()
	if err != nil {
		helper.RedirectToError(w, r, statusCode, err.Error())
		return
	}

	if statusCode != http.StatusOK {
		helper.RedirectToError(w, r, statusCode, "Erreur lors de la récupération des données")
		return
	}

	helper.RenderTemplate(w, r, "descript-countries", data)
}

func DisplaySearch(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("query")

	query = strings.TrimSpace(strings.ToLower(query))
	if query == "" {
		http.Redirect(w, r, "/countries", http.StatusMovedPermanently)
		return
	}

	data, dataStatus, dataError := services.GetCountries()
	if dataError != nil {
		helper.RedirectToError(w, r, dataStatus, dataError.Error())
		return
	}

	if dataStatus != http.StatusOK {
		helper.RedirectToError(w, r, dataStatus, "Erreur Service")
		return
	}

	searchList := []services.Countrie{}

	for _, item := range *data {
		checkRegion := strings.Contains(strings.ToLower(item.Region), query)
		checkSousRegion := strings.Contains(strings.ToLower(item.SousRegion), query)
		checkNameOff := strings.Contains(strings.ToLower(item.Name.NameOfficiel), query)
		// fmt.Println(item.Name.NameOfficiel, query, checkNameOff)
		checkNameUti := strings.Contains(strings.ToLower(item.Name.NameUtilisé), query)

		if checkRegion || checkSousRegion || checkNameOff || checkNameUti {
			searchList = append(searchList, item)
		}
	}
	helper.RenderTemplate(w, r, "search_countries", searchList)
}

func DisplayFilter(w http.ResponseWriter, r *http.Request) {
	popNbr, _ := strconv.Atoi(strings.TrimSpace(r.FormValue("pop")))

	r.ParseForm()
	region := r.Form["region"]
	language := r.Form["language"]

	for i := range region {
		region[i] = strings.ToLower(region[i])
	}
	for i := range language {
		language[i] = strings.ToLower(language[i])
	}

	data, dataStatusCode, dataError := services.GetCountriesFilter()
	if dataStatusCode != http.StatusOK || dataError != nil {
		log.Printf("Erreur DisplayFilter - %s", dataError.Error())
		http.Error(w, fmt.Sprintf("Erreur service - code : %d \n message: %v", dataStatusCode, dataError.Error()), dataStatusCode)
		return
	}

	valideCountry := []services.Countrie{}

	for _, country := range *data {

		checkRegion := false
		for _, r := range region {
			if strings.EqualFold(strings.TrimSpace(country.Region), strings.TrimSpace(r)) {
				checkRegion = true
				break
			}
		}

		checkLanguage := false
		if country.Languages != nil {
			for _, lang := range country.Languages {
				if slices.Contains(language, strings.ToLower(lang)) {
					checkLanguage = true
					break
				}
			}
		}

		checkPop := country.Pop >= popNbr

		if (checkRegion || len(region) == 0) &&
			(checkLanguage || len(language) == 0) &&
			(checkPop || popNbr == 0) {
			valideCountry = append(valideCountry, country)
		}
	}

	helper.RenderTemplate(w, r, "filter_countries", valideCountry)
}

type PagePagination struct {
	Page int
	Next int
	Prev int
	Data []services.Countrie
}

func DisplayPagination(w http.ResponseWriter, r *http.Request) {
	pageStr := r.FormValue("page")
	pageInt, _ := strconv.Atoi(pageStr)

	if pageInt < 0 {
		pageInt = 0
	}

	startIndex := pageInt * 15
	endIndex := (pageInt * 15) + 15

	data, statusCode, err := services.GetCountries()
	if err != nil {
		helper.RedirectToError(w, r, statusCode, err.Error())
		return
	}

	if statusCode != http.StatusOK {
		helper.RedirectToError(w, r, statusCode, "Erreur Service")
		return
	}

	if endIndex > len((*data)) {
		endIndex = len((*data))
	}

	if startIndex > len((*data)) {
		pageInt = 0
		startIndex = 0
		endIndex = 15
	}

	SelectCountries := (*data)[startIndex:endIndex]

	vieData := PagePagination{pageInt, (pageInt + 1), (pageInt - 1), SelectCountries}

	helper.RenderTemplate(w, r, "pagination", vieData)
}
