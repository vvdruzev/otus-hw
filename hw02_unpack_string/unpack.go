package hw02unpackstring

import (
	"errors"
	"fmt"
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
		return "", fmt.Errorf("некорректная строка")
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
			if _, err = strconv.Atoi(str2[i] + str2[i+1]); err == nil {
				return "", fmt.Errorf("некорректная строка")
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
