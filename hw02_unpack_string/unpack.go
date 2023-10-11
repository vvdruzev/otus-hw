package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	end := len(str)
	if end == 0 {
		return "", nil
	}
	_, err := strconv.Atoi(str[0:1])
	if err == nil {
		return "", ErrInvalidString
	}
	if end == 1 {
		return str, nil
	}

	str2 := strings.Split(str, "")

	var result string
	var c string
	end = len(str2)

	for i := 0; i < end-1; i++ {
		c += str2[i]
		if i <= end-2 {
			if _, err = strconv.Atoi(str2[i+1] + str2[i+2]); err == nil {
				return "", ErrInvalidString
			}
		}
		number, err := strconv.Atoi(str2[i+1])
		if err != nil {
			result += c
			c = ""
			if i == end-2 {
				result += str2[i+1]
			}
			continue
		}
		result += strings.Repeat(c, number)
		c = ""
		i++
		if i == end-2 {
			result += str2[i+1]
		}
	}
	return result, nil
}
