package valueobject

import (
	"errors"
	"regexp"
	"strings"
)

// Errors block
var (
	ErrInvalidRomanNumeral = errors.New("Not a valid numeral")
	ErrNegativeValue       = errors.New("can not represent negative numbers")
)

var (
	lookup = map[string]uint{
		"I":  1,
		"IV": 4,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
	}

	order = []string{"M", "D", "C", "L", "X", "V", "IV", "I"}
)

// RomanNumeral represents a number
type RomanNumeral struct {
	value string
}

// NewRomanNumeral creates a new numeral
func NewRomanNumeral(v interface{}) (RomanNumeral, error) {
	var n RomanNumeral
	switch t := v.(type) {
	case string:
		match, _ := regexp.MatchString("[MDCLXVI]+", t)
		if !match {
			return n, ErrInvalidRomanNumeral
		}
		n.value = t
	case uint:
		n.value = itoa(t)
	case int:
		if t < 0 {
			return n, ErrNegativeValue
		}
		n.value = itoa(uint(t))
	default:
		return n, ErrInvalidRomanNumeral
	}

	return n, nil
}

// String returns string representation
func (n RomanNumeral) String() string {
	return n.value
}

// Equals checks that two values are the same
func (n RomanNumeral) Equals(value Value) bool {
	otherRomanNumeral, ok := value.(RomanNumeral)
	return ok && n.value == otherRomanNumeral.value
}

func itoa(in uint) string {
	remainder := in
	acc := ""

	for _, k := range order {
		currentVal := lookup[k]
		c := remainder / currentVal
		acc += strings.Repeat(k, int(c))
		remainder = remainder - (c * currentVal)
	}

	return acc
}
