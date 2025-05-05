package pokedex

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Area struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

var cache = NewCache(5 * time.Second)

func Locations(url string) (Area, error) {
	area := Area{}
	data, hit := cache.Get(url)
	if hit {
		json.Unmarshal(data, &area)
		return area, nil
	} else {
		res, err := http.Get(url)

		if err != nil {
			return Area{}, err
		}
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return Area{}, err
		}
		if res.StatusCode > 299 {
			return Area{}, fmt.Errorf("response failed with status: %v", res.StatusCode)
		}
		err = json.Unmarshal(body, &area)
		if err != nil {
			return Area{}, err
		}
		cache.Add(url, body)

	}

	return area, nil
}
