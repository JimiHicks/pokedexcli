package pokedexapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetLocationAreas() (*GetLocationAreasResponse, error) {
	return GetLocationAreasFromURL("https://pokeapi.co/api/v2/location-area")
}

func GetLocationAreasFromURL(url string) (*GetLocationAreasResponse, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if res.StatusCode > 299 {
		return nil, fmt.Errorf("response failed with status code: %d and\nbody: %s", res.StatusCode, body)
	}

	if err != nil {
		return nil, err
	}

	response := GetLocationAreasResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type GetLocationAreasResponse struct {
	Count    int      `json:"count"`
	Next     *string  `json:"next"`
	Previous *string  `json:"previous"`
	Results  []Result `json:"results"`
}

type Result struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
