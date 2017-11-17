package stapiclient

import (
	"net/url"
	"net/http"
	"strings"
	"fmt"
	"encoding/json"
	"errors"
)

const baseURL = "http://stapi.co/api/v1/rest"

type searchResult struct {
	Characters []character `json:"characters"`
}

type character struct {
	Uid              string            `json:"uid"`
	Name             string            `json:"name"`
	CharacterSpecies []characterSpecie `json:"name"`
}

type characterSpecie struct {
	Uid  string `json:"uid"`
	Name string `json:"name"`
}

func getCharacterUid(characterName string) (string, error) {
	searchURL := fmt.Sprintf("%s/character/search", baseURL)

	form := url.Values{}
	form.Add("name", characterName)

	req, err := http.NewRequest("POST", searchURL, strings.NewReader(form.Encode()))
	if err != nil {
		return "", err
	}

	hc := http.Client{}
	res, err := hc.Do(req)

	if err != nil {
		return "", err
	}

	var result searchResult
	json.NewDecoder(res.Body).Decode(&result)

	defer res.Body.Close()

	if len(result.Characters) == 0 {
		return "", errors.New("not found")
	}

	return result.Characters[0].Uid, nil
}

func GetCharacterRace(characterName string) string {
	return "Dummy"
}
