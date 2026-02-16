package services

import (
	"encoding/json"
	"fmt"
	"os"
)

type NameFav struct {
	NameOfficiel string `json:"official"`
	NameUtilisé  string `json:"common"`
}

type FavorisData struct {
	NameFav NameFav `json:"name"`
}

const favorisFile = "favoris.json"

func GetFavoriteNames() []string {
	file, err := os.ReadFile(favorisFile)
	if err != nil {
		return []string{}
	}

	var data []FavorisData
	json.Unmarshal(file, &data)
	var names []string
	for _, f := range data {
		names = append(names, f.NameFav.NameOfficiel)
	}
	return names
}

func IsFavorite(name string) bool {
	names := GetFavoriteNames()
	for _, savedName := range names {
		if savedName == name {
			return true
		}
	}
	return false
}

func ToggleFavoriteName(name string) {
	names := GetFavoriteNames() // renvoie []string
	exists := false
	var newNames []string

	for _, savedName := range names {
		if savedName == name {
			exists = true
		} else {
			newNames = append(newNames, savedName)
		}
	}

	if !exists {
		newNames = append(newNames, name)
	}

	var data []FavorisData
	for _, n := range newNames {
		data = append(data, FavorisData{
			NameFav: NameFav{
				NameOfficiel: n,
				NameUtilisé:  n,
			},
		})
	}

	fileContent, _ := json.MarshalIndent(data, "", "  ")
	os.WriteFile(favorisFile, fileContent, 0644)
}

func GetAllFavories() ([]Countrie, error) {
	names := GetFavoriteNames()
	var countriesFav []Countrie

	for _, name := range names {
		countryFav, err := GetCountryByName(name)
		if err == nil {
			temp := countryFav
			temp.IsFavorite = true
			countriesFav = append(countriesFav, temp)
		}
	}
	return countriesFav, nil
}

func GetCountryByName(name string) (Countrie, error) {
	countries, _, err := GetCountries()
	if err != nil {
		return Countrie{}, fmt.Errorf("Impossible de récupérer les pays : %s", err.Error())
	}

	for _, country := range *countries {
		if country.Name.NameUtilisé == name || country.Name.NameOfficiel == name {
			return country, nil
		}
	}

	return Countrie{}, fmt.Errorf("Pays non trouvé : %s", name)
}
