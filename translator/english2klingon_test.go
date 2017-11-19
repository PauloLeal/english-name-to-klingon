package translator

import (
	"testing"
	"fmt"
)

func TestEnglishToKlingon(t *testing.T) {
	t.Run("Uhura", auxTestEnglishToKlingon("Uhura", "0xF8E5 0xF8D6 0xF8E5 0xF8E1 0xF8D0", false))
	t.Run("Jim", auxTestEnglishToKlingon("Jim", "0xF8D8 0xF8D7 0xF8DA", false))
	t.Run("Ch'Targh", auxTestEnglishToKlingon("Ch'Targh", "0xF8D2 0xF8E9 0xF8E3 0xF8D0 0xF8E1 0xF8D5", false))
	t.Run("with spaces", auxTestEnglishToKlingon("Jim Jim", "0xF8D8 0xF8D7 0xF8DA 0x0020 0xF8D8 0xF8D7 0xF8DA", false))
	t.Run("invalid", auxTestEnglishToKlingon("kirk", "invalid", true))

}

func auxTestEnglishToKlingon(englishWord string, klingonWord string, hasError bool) func(*testing.T) {
	return func(t *testing.T) {
		translated, err := EnglishToKlingon(englishWord)

		if err != nil && !hasError {
			t.Error(fmt.Sprintf("Expected translation of %s to be successful", englishWord))
		} else if err == nil && hasError {
			t.Error(fmt.Sprintf("Expected translation of %s to be an error", englishWord))
		} else if !hasError && translated != klingonWord {
			t.Error(fmt.Sprintf("Expected translation of %s \nto be\t%s. \nGot\t\t%s", englishWord, klingonWord, translated))
		}

	}
}
