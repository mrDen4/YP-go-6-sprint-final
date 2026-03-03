package service

import (
	"strings"
	"unicode"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func isMorse(str string) bool {
	str = strings.TrimSpace(str)

	// if str == "" {
	// 	return false
	// }

	// allowed := ".- /"

	// if !strings.ContainsAny(str, ".-") {
	// 	return false
	// }

	// for _, r := range str {
	// 	if !strings.ContainsRune(allowed, r) {
	// 		return false
	// 	}
	// }
	if strings.ContainsFunc(str, unicode.IsLetter) {
		return false
	}

	return true
}

func Service(str string) string {
	isMorseInput := isMorse(str)

	if isMorseInput {
		return morse.ToText(str)
	}

	return morse.ToMorse(str)
}
