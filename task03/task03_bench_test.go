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

func BenchmarkFerato(b *testing.B) {
	var a = 10000
	Ferato(a)
}

func BenchmarkFeratob(b *testing.B) {
	var a = 10000
	Feratob(a)
}

func BenchmarkFibo(b *testing.B) {
	var a = 32
	Fibo(a)
}

func BenchmarkFibo2(b *testing.B) {
	var a = 300
	Fibo2(a)
}

func BenchmarkFibo3(b *testing.B) {
	var a = 302
	Fibo3(a)
}
