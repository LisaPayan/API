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

	var names []string
	if err := json.Unmarshal(file, &names); err != nil {
		return []string{}
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
	names := GetFavoriteNames()
	exists := false
	var newNames []string

	for _, n := range names {
		if n == name {
			exists = true
		} else {
			newNames = append(newNames, n)
		}
	}

	if !exists {
		newNames = append(newNames, name)
	}

	fileContent, _ := json.MarshalIndent(newNames, "", "  ")
	_ = os.WriteFile(favorisFile, fileContent, 0644)
}

func GetAllFavories() ([]Countrie, error) {
	countries, _, err := GetCountries()
	if err != nil {
		return nil, err
	}

	favNames := GetFavoriteNames()
	var favorites []Countrie

	for _, country := range *countries {
		for _, f := range favNames {
			if country.Name.NameOfficiel == f || country.Name.NameUtilisé == f {
				temp := country
				temp.IsFavorite = true
				favorites = append(favorites, temp)
				break
			}
		}
	}

	return favorites, nil
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
