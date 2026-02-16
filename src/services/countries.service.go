package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Name struct {
	NameOfficiel string `json:"official"`
	NameUtilisé  string `json:"common"`
}

type Maps struct {
	GoogleMaps string `json:"googleMaps"`
}

type Flag struct {
	PngFlag string `json:"png"`
	AltFlag string `json:"alt"`
}

type Countrie struct {
	Name       Name              `json:"name"`
	Languages  map[string]string `json:"languages"`
	Capitale   []string          `json:"capital"`
	Region     string            `json:"region"`
	SousRegion string            `json:"subregion"`
	BordPays   []string          `json:"borders"`
	Maps       Maps              `json:"maps"`
	Pop        int               `json:"population"`
	Area       float64           `json:"area"`
	Flag       Flag              `json:"flags"`
	IsFavorite bool              `json:"isFavorite"`
}

func GetCountries() (*[]Countrie, int, error) {
	_client := http.Client{
		Timeout: time.Second * 10,
	}

	request, requestErr := http.NewRequest(http.MethodGet, "https://restcountries.com/v3.1/all?fields=name,flags,region,language", nil)
	if requestErr != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("Erreur préparation requête - %s", requestErr.Error())
	}

	response, responseErr := _client.Do(request)

	if responseErr != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("Erreur envoi de la requête - %s", responseErr.Error())
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, response.StatusCode, fmt.Errorf("Erreur réponse - %s", response.Status)
	}

	var countries []Countrie

	decodeErr := json.NewDecoder(response.Body).Decode(&countries)
	if decodeErr != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("Erreur décode de données - %s", decodeErr.Error())
	}

	return &countries, response.StatusCode, nil
}
