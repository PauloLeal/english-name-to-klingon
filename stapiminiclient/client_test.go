package stapiminiclient

import (
	"testing"
	"fmt"
)

func TestGetCharacterUid(t *testing.T) {
	t.Run("Uhura", auxTestGetCharacterUid("uhura", "CHMA0000068639", false))
	t.Run("Jim", auxTestGetCharacterUid("jim", "CHMA0000009602", false))
	t.Run("INVALID_", auxTestGetCharacterUid("invalid_", "", true))
}

func auxTestGetCharacterUid(query string, expectedUid string, hasError bool) func(*testing.T) {
	return func(t *testing.T) {
		n, err := getCharacterUid(query)

		if hasError && err == nil {
			t.Error(fmt.Sprintf("Expected getCharacterUid(%s) to return an error. But it returned with result %s", query, n))
		} else if !hasError && err != nil {
			t.Error(fmt.Sprintf("Expected getCharacterUid(%s) to be successful but got an error. %s", query, err))
		} else if !hasError && n != expectedUid {
			t.Error(fmt.Sprintf("Expected uid of %s to be %s. Found %s", query, expectedUid, n))
		}
	}
}

func TestGetCharacterSpecieByUid(t *testing.T) {
	t.Run("Uhura", auxTestGetCharacterSpecieByUid("CHMA0000068639", "Human"))
	t.Run("Jim", auxTestGetCharacterSpecieByUid("CHMA0000009602", "Human"))
	t.Run("B'Etor", auxTestGetCharacterSpecieByUid("CHMA0000003645", "Klingon"))
	t.Run("INVALID_", auxTestGetCharacterSpecieByUid("invalid_", "Unknown"))
}

func auxTestGetCharacterSpecieByUid(uid string, expectedSpecie string) func(*testing.T) {
	return func(t *testing.T) {
		race, err := getCharacterSpecieByUid(uid)

		if err != nil {
			t.Error(fmt.Sprintf("Expected getCharacterSpecieByUid(%s) to be successful but got an error. %s", uid, err))
		} else if race != expectedSpecie {
			t.Error(fmt.Sprintf("Expected race of %s to be %s. Found %s", uid, expectedSpecie, race))
		}
	}
}

func TestGetCharacterRace(t *testing.T) {
	t.Run("testing valid character", auxTestGetCharacterRace("Uhura", "Human"))
	t.Run("testing valid character", auxTestGetCharacterRace("B'Etor", "Klingon"))
	t.Run("testing valid character", auxTestGetCharacterRace("Spock", "Unknown")) //First match on API has no specie information
	t.Run("testing valid character", auxTestGetCharacterRace("INVALID_", "Unknown"))
}

func auxTestGetCharacterRace(characterName string, expectedSpecie string) func(*testing.T) {
	return func(t *testing.T) {
		s := GetCharacterSpecie(characterName)
		if s != expectedSpecie {
			t.Error(fmt.Sprintf("Expected %s race to be %s. Found %s", characterName, expectedSpecie, s))
		}
	}
}
