package pokedex

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Area struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func Locations(url string) (Area, error) {
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area/"
	}
	res, err := http.Get(url)
	if err != nil {
		return Area{}, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Area{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return Area{}, fmt.Errorf("response failed with status: %v", res.StatusCode)
	}

	area := Area{}
	err = json.Unmarshal(body, &area)
	if err != nil {
		return Area{}, err
	}

	// for location := range area.Results {
	// 	fmt.Println(area.Results[location].Name)
	// }
	return area, nil
}
