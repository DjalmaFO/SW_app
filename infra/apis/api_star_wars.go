package apis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type APIResponse struct {
	Results []struct {
		Films []string `json:"films"`
	} `json:"results"`
}

// CountFilms => Utiliza API pública da StarWar para retorna a quantidade de aparições de determinado planeta em filmes
func CountFilms(name string) (amount int, err error) {
	resp, err := http.Get(fmt.Sprintf("https://swapi.dev/api/planets/?search=%s", name))
	if err != nil {
		log.Println(err.Error())
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return
	}

	data := APIResponse{}
	if err = json.Unmarshal(body, &data); err != nil {
		log.Println(err.Error())
		return
	}

	if len(data.Results) == 0 {
		return
	}

	return len(data.Results[0].Films), nil
}
