package main

import "testing"

var testNumbers = []uint{0, 1, 2, 4, 8, 9, 27, 938641, 987948554334, 1000000000, 999999999}
var testBases = []uint{0, 2, 8, 16, 62, 26, 36, 52}

func TestBaseConversion(t *testing.T) {
	for _, base := range testBases {
		baseconv, err := NewBaseConvertor(base)
		if err != nil && base != 0 {
			t.Error("Base 0 should not be accepted")
		}
		if err != nil {
			continue
		}

		for _, n := range testNumbers {
			if baseconv.Decode(baseconv.Encode(n)) != n {
				t.Error("Base", base, "N=", n, "failed")
			}
		}
	}
}
