package utils

import (
	"regexp"
	"strings"
	"unicode"
)

var (
	matchFirstCap = regexp.MustCompile(`(.)[A-Z][a-z]+`)
	matchAllCap   = regexp.MustCompile(`([a-z0-9])([A-Z])`)
)
func CamelToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(str, "${1}_${2}")
	return strings.ToLower(snake)
}
func NormailizeString(text string) string {
	return strings.ToLower(strings.TrimSpace(text))
}

func ConvertRating(rating float32) float32 {
	return float32(int(rating * 10)) / 10
}

func IsNumeric(s string) bool {
    if s == "" {
        return false
    }
    for _, r := range s {
        if !unicode.IsDigit(r) {
            return false
        }
    }

    return true
}