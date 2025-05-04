package pokedexcli

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Locations(conf *config) Area {
	endPoint := "https://pokeapi.co/api/v2/location-area/"
	res, err := http.Get(endPoint)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return fmt.Errorf("response failed with status: %v", res.StatusCode)
	}
	err = json.Unmarshal(body, &conf)
	if err != nil {
		return err
	}
	area := Area{}
	err = json.Unmarshal(body, &area)
	if err != nil {
		return err
	}

	// for location := range area.Results {
	// 	fmt.Println(area.Results[location].Name)
	// }
	return area
}
