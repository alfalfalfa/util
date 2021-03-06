package util

import (
	"strings"
	"unicode"
)

// CamelToSnake converts a given string to snake case
func CamelToSnake(s string) string {
	var result string
	var words []string
	var lastPos int
	rs := []rune(s)

	for i := 0; i < len(rs); i++ {
		if i > 0 && unicode.IsUpper(rs[i]) {
			if initialism := startsWithInitialism(s[lastPos:]); initialism != "" {
				words = append(words, initialism)

				i += len(initialism) - 1
				lastPos = i
				continue
			}

			words = append(words, s[lastPos:i])
			lastPos = i
		}
	}

	// append the last word
	if s[lastPos:] != "" {
		words = append(words, s[lastPos:])
	}

	for k, word := range words {
		if k > 0 {
			result += "_"
		}

		result += strings.ToLower(word)
	}

	return result
}

// SnakeToCamel returns a string converted from snake case to uppercase
func SnakeToCamel(s string) string {
	var result string

	words := strings.Split(s, "_")

	for _, word := range words {
		if upper := strings.ToUpper(word); commonInitialisms[upper] {
			result += upper
			continue
		}

		w := []rune(word)
		w[0] = unicode.ToUpper(w[0])
		result += string(w)
	}

	return result
}

func SnakeToLowerCamel(s string) string {
	var result string

	words := strings.Split(s, "_")

	fw := true
	for _, word := range words {
		if fw {
			fw = false
			if upper := strings.ToUpper(word); commonInitialisms[upper] {
				result += strings.ToLower(upper)
				continue
			}

			w := []rune(word)
			w[0] = unicode.ToLower(w[0])
			result += string(w)
			continue
		}
		if upper := strings.ToUpper(word); commonInitialisms[upper] {
			result += upper
			continue
		}

		w := []rune(word)
		w[0] = unicode.ToUpper(w[0])
		result += string(w)
	}

	return result
}

// startsWithInitialism returns the initialism if the given string begins with it
func startsWithInitialism(s string) string {
	var initialism string
	// the longest initialism is 5 char, the shortest 2
	for i := 1; i <= 5; i++ {
		if len(s) > i-1 && commonInitialisms[s[:i]] {
			initialism = s[:i]
		}
	}
	return initialism
}

// commonInitialisms, taken from
// https://github.com/golang/lint/blob/3d26dc39376c307203d3a221bada26816b3073cf/lint.go#L482
var commonInitialisms = map[string]bool{
	"API":   true,
	"ASCII": true,
	"CPU":   true,
	"CSS":   true,
	"DNS":   true,
	"EOF":   true,
	"GUID":  true,
	"HTML":  true,
	"HTTP":  true,
	"HTTPS": true,
	//"ID":    true,
	"IP":   true,
	"JSON": true,
	"LHS":  true,
	"QPS":  true,
	"RAM":  true,
	"RHS":  true,
	"RPC":  true,
	"SLA":  true,
	"SMTP": true,
	"SSH":  true,
	"TLS":  true,
	"TTL":  true,
	"UI":   true,
	"UID":  true,
	"UUID": true,
	"URI":  true,
	"URL":  true,
	"UTF8": true,
	"VM":   true,
	"XML":  true,
}

type CaseString string

func NewCaseString(s string) CaseString {
	//????????????????????????????????????????????????
	return CaseString(strings.ToLower(CamelToSnake(s)))
}
func (this CaseString) Lower() string {
	return strings.ToLower(string(this))
}
func (this CaseString) Upper() string {
	return strings.ToUpper(this.Lower())
}
func (this CaseString) LowerSnake() string {
	return strings.ToLower(CamelToSnake(this.Lower()))
}
func (this CaseString) UpperSnake() string {
	return strings.ToUpper(CamelToSnake(this.Lower()))
}
func (this CaseString) LowerCamel() string {
	return SnakeToLowerCamel(this.Lower())
}
func (this CaseString) UpperCamel() string {
	return SnakeToCamel(this.Lower())
}
