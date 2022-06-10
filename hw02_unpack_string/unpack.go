package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	// Place your code here.
	runes := []rune(str)
	var sb strings.Builder
	escaping := false

	for i := 0; i < len(runes); i++ {
		if runes[i] == '\\' && !escaping {
			escaping = true
			continue
		}
		if unicode.IsDigit(runes[i]) && !escaping {
			if i == 0 || (i < len(runes)-1 && unicode.IsDigit(runes[i+1])) {
				return "", ErrInvalidString
			}
			if num, err := strconv.Atoi(string(runes[i])); err == nil {
				sb.WriteString(strings.Repeat(string(runes[i-1]), num))
			}
		} else if i == len(runes)-1 || i < len(runes)-1 && !unicode.IsDigit(runes[i+1]) {
			sb.WriteRune(runes[i])
		}
		escaping = false
	}
	return sb.String(), nil
}
