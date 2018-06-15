package main

import "errors"
import "math"

var (
	// Charset is all the alphabets that can be used
	Charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// ErrInvalidBase is thrown when the base is <= 0
	ErrInvalidBase = errors.New("Invalid base")
)

// Reverse a string
// https://stackoverflow.com/a/10030772/5163807
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// BaseConvertor represents the data structure required for base conversion
type BaseConvertor struct {
	base        uint
	alphabet    []rune
	alphabetMap map[rune]uint
}

// NewBaseConvertor instantiates a new BaseConvertor object
func NewBaseConvertor(base uint) (*BaseConvertor, error) {
	if base == 0 {
		return nil, ErrInvalidBase
	}

	runes := []rune(Charset[0:base])
	runeMap := make(map[rune]uint)

	var i uint
	for i = 0; i < base; i++ {
		runeMap[runes[i]] = i
	}

	return &BaseConvertor{
		base:        base,
		alphabet:    runes,
		alphabetMap: runeMap,
	}, nil
}

// Encode an unsigned integer to another base
func (e *BaseConvertor) Encode(number uint) string {
	remainders := make([]rune, 0)

	n := number
	for n > 0 {
		r := rune(Charset[n%e.base])
		remainders = append(remainders, r)
		n = n / e.base
	}
	return Reverse(string(remainders))
}

// Decode string in given base to unsigned integer
func (e *BaseConvertor) Decode(code string) uint {
	var value float64
	value = 0

	for pos, val := range Reverse(code) {
		value = value + float64(e.alphabetMap[val])*math.Pow(float64(e.base), float64(pos))
	}
	return uint(value)
}
