package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var result string
	runes := []rune(str)
	if len(runes) == 0 {
		return result, nil
	}
	remember := runes[0]

	if unicode.IsDigit(remember) {
		return result, ErrInvalidString
	}

	for i := 1; i < len(runes); i++ {
		current := runes[i]
		if unicode.IsDigit(current) {
			if unicode.IsDigit(remember) {
				return "", ErrInvalidString
			}
			number, _ := strconv.Atoi(string(current))
			result += strings.Repeat(string(remember), number)
			remember = current
			continue
		}
		if !unicode.IsDigit(remember) {
			result += string(remember)
		}

		remember = current
	}
	if !unicode.IsDigit(remember) {
		result += string(remember)
	}
	return result, nil
}
