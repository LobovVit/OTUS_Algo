package task03

import (
	"testing"
)

func BenchmarkCalc(b *testing.B) {
	var a = 1234567890
	var bb = 123
	Fcalc(a, bb)
}

func BenchmarkCalc2(b *testing.B) {
	var a = 1234567890
	var bb = 123
	Fcalc2(a, bb)
}
