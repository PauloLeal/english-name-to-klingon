package stapiminiclient

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

type characterResult struct {
	Character character `json:"character"`
}

type character struct {
	Uid              string            `json:"uid"`
	Name             string            `json:"name"`
	CharacterSpecies []characterSpecie `json:"characterSpecies"`
}

type characterSpecie struct {
	Uid  string `json:"uid"`
	Name string `json:"name"`
}

func getCharacterUid(characterName string) (string, error) {
	searchURL := fmt.Sprintf("%s/character/search", baseURL)

	form := url.Values{}
	form.Add("name", strings.Title(characterName))

	req, err := http.NewRequest("POST", searchURL, strings.NewReader(form.Encode()))
	if err != nil {
		return "", err
	}

	hc := http.Client{}
	res, err := hc.Do(req)

	if err != nil || res.StatusCode != 200 {
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

func getCharacterSpecieByUid(uid string) (string, error) {
	u := fmt.Sprintf("%s/character?uid=%s", baseURL, uid)
	res, err := http.Get(u)

	defer res.Body.Close()

	if err != nil || res.StatusCode != 200 {
		return "", err
	}

	var result characterResult
	json.NewDecoder(res.Body).Decode(&result)

	s := make([]string, len(result.Character.CharacterSpecies))

	for i, r := range result.Character.CharacterSpecies {
		s[i] = strings.Title(r.Name)
	}

	if len(s) > 0 {
		return strings.Join(s, " / "), nil
	} else {
		return "Unknown", nil
	}
}

func GetCharacterSpecie(characterName string) string {
	uid, err := getCharacterUid(characterName)
	if err != nil {
		return "Unknown"
	}

	specie, err := getCharacterSpecieByUid(uid)
	if err != nil {
		return "Unknown"
	}

	return strings.Title(specie)
}
