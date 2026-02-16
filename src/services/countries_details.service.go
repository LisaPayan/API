package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type NameDetails struct {
	NameOfficiel string `json:"official"`
	NameUtilisé  string `json:"common"`
}

type MapsDetails struct {
	GoogleMaps string `json:"googleMaps"`
}

type FlagDetails struct {
	PngFlag string `json:"png"`
	AltFlag string `json:"alt"`
}

type Currency struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type CountryDetails struct {
	Name       NameDetails         `json:"name"`
	Capitale   []string            `json:"capital"`
	Region     string              `json:"region"`
	SousRegion string              `json:"subregion"`
	BordPays   []string            `json:"borders"`
	Maps       MapsDetails         `json:"maps"`
	Pop        int                 `json:"population"`
	Area       float64             `json:"area"`
	Flag       FlagDetails         `json:"flags"`
	Currencies map[string]Currency `json:"currencies"`
	Languages  map[string]string   `json:"languages"`
}

func GetCountriesDetails(name string) (*CountryDetails, int, error) {
	_client := http.Client{
		Timeout: time.Second * 5,
	}

	url := "https://restcountries.com/v3.1/name/" + name

	request, requestErr := http.NewRequest(http.MethodGet, url, nil)
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

	var countrydetails []CountryDetails

	decodeErr := json.NewDecoder(response.Body).Decode(&countrydetails)
	if decodeErr != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("Erreur décode de données - %s", decodeErr.Error())
	}

	return &countrydetails[0], response.StatusCode, nil

}
