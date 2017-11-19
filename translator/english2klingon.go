package translator

import (
	"strings"
	"fmt"
	"errors"
)

type translationTable struct {
	EnSeq string
	Kc    int64
}

func getTranslationTable() []translationTable {
	var t []translationTable

	t = append(t, translationTable{EnSeq: "tlh", Kc: '\uF8E4'})

	t = append(t, translationTable{EnSeq: "gh", Kc: '\uF8D5'})
	t = append(t, translationTable{EnSeq: "ng", Kc: '\uF8DC'})
	t = append(t, translationTable{EnSeq: "ch", Kc: '\uF8D2'})

	t = append(t, translationTable{EnSeq: "a", Kc: '\uF8D0'})
	t = append(t, translationTable{EnSeq: "b", Kc: '\uF8D1'})
	t = append(t, translationTable{EnSeq: "d", Kc: '\uF8D3'})
	t = append(t, translationTable{EnSeq: "e", Kc: '\uF8D4'})
	t = append(t, translationTable{EnSeq: "h", Kc: '\uF8D6'})
	t = append(t, translationTable{EnSeq: "i", Kc: '\uF8D7'})
	t = append(t, translationTable{EnSeq: "j", Kc: '\uF8D8'})
	t = append(t, translationTable{EnSeq: "l", Kc: '\uF8D9'})
	t = append(t, translationTable{EnSeq: "m", Kc: '\uF8DA'})
	t = append(t, translationTable{EnSeq: "n", Kc: '\uF8DB'})
	t = append(t, translationTable{EnSeq: "o", Kc: '\uF8DD'})
	t = append(t, translationTable{EnSeq: "p", Kc: '\uF8DE'})
	t = append(t, translationTable{EnSeq: "q", Kc: '\uF8DF'})
	t = append(t, translationTable{EnSeq: "q", Kc: '\uF8E0'})
	t = append(t, translationTable{EnSeq: "r", Kc: '\uF8E1'})
	t = append(t, translationTable{EnSeq: "s", Kc: '\uF8E2'})
	t = append(t, translationTable{EnSeq: "t", Kc: '\uF8E3'})
	t = append(t, translationTable{EnSeq: "u", Kc: '\uF8E5'})
	t = append(t, translationTable{EnSeq: "v", Kc: '\uF8E6'})
	t = append(t, translationTable{EnSeq: "w", Kc: '\uF8E7'})
	t = append(t, translationTable{EnSeq: "y", Kc: '\uF8E8'})

	t = append(t, translationTable{EnSeq: "0", Kc: '\uF8F0'})
	t = append(t, translationTable{EnSeq: "1", Kc: '\uF8F1'})
	t = append(t, translationTable{EnSeq: "2", Kc: '\uF8F2'})
	t = append(t, translationTable{EnSeq: "3", Kc: '\uF8F3'})
	t = append(t, translationTable{EnSeq: "4", Kc: '\uF8F4'})
	t = append(t, translationTable{EnSeq: "5", Kc: '\uF8F5'})
	t = append(t, translationTable{EnSeq: "6", Kc: '\uF8F6'})
	t = append(t, translationTable{EnSeq: "7", Kc: '\uF8F7'})
	t = append(t, translationTable{EnSeq: "8", Kc: '\uF8F8'})
	t = append(t, translationTable{EnSeq: "9", Kc: '\uF8F9'})

	t = append(t, translationTable{EnSeq: ".", Kc: '\uF8FD'})
	t = append(t, translationTable{EnSeq: ",", Kc: '\uF8FE'})
	t = append(t, translationTable{EnSeq: " ", Kc: '\u0020'})
	t = append(t, translationTable{EnSeq: "'", Kc: '\uF8E9'})
	t = append(t, translationTable{EnSeq: "`", Kc: '\uF8E9'}) //contingency

	t = append(t, translationTable{EnSeq: "?", Kc: '\uF8FF'})

	return t
}

func EnglishToKlingon(word string) (string, error) {
	var result []string

	for i := 0; i < len(word); {
		found := false

		for _, t := range getTranslationTable() {
			if strings.Index(strings.ToLower(word)[i:], t.EnSeq) == 0 {
				result = append(result, fmt.Sprintf("0x%04X", t.Kc))
				i += len(t.EnSeq)
				found = true
				break
			}
		}

		if !found {
			return "", errors.New("invalid input")
		}
	}

	return strings.Join(result, " "), nil
}
