package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var result string
	runes := []rune(str)
	if len(runes) == 0 {
		return result, nil
	}
	remember := runes[0]

	_, err := strconv.Atoi(string(remember))
	if err == nil {
		return result, ErrInvalidString
	}

	for i := 1; i < len(runes); i++ {
		current := runes[i]
		if _, err = strconv.Atoi(string(current)); err == nil {
			if _, err = strconv.Atoi(string(remember)); err == nil {
				return "", ErrInvalidString
			}
			number, _ := strconv.Atoi(string(current))
			result += strings.Repeat(string(remember), number)
			remember = current
			continue
		}
		if _, err = strconv.Atoi(string(remember)); err != nil {
			result += string(remember)
		}

		remember = current
	}
	if _, err = strconv.Atoi(string(remember)); err != nil {
		result += string(remember)
	}
	return result, nil
}
