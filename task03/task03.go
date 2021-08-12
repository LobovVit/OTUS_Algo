package task03

import (
	"math"
	"math/big"
)

func Fcalc(a int, b int) int {
	for i := 1; i >= 1; i++ {
		if a > b {
			a = a - b
		} else if a < b {
			b = b - a
		} else if a == b {
			return a
		}
	}
	return a
}

func Fcalc2(a int, b int) int {
	for i := 1; i >= 1; i++ {
		if (a == 0) || (b == 0) {
			return a + b
		} else if a > b {
			a = a % b
		} else if a < b {
			b = b % a
		}
	}
	return a + b
}

func Ferato(a int) []int {
	mas := make([]int, a)
	//наполнение
	for i := 1; i < a; i++ {
		mas[i] = i

	}
	//выгрызаем
	for i := 1; i < a; i++ {
		if mas[i] != 0 && mas[i] != 1 {
			for index, _ := range mas {
				if index%i == 0 && index != i {
					mas[index] = 0
				}
			}
		}
	}
	return mas
	//Фибоначи
	//фибоначи 2
	//фиб матрица
}

func Feratob(a int) []bool {
	mas := make([]bool, a)

	//выгрызаем
	for i := 1; i < a; i++ {
		if !mas[i] && i != 1 {
			for index, _ := range mas {
				if index%i == 0 && index != i {
					mas[index] = true
				}
			}
		}
	}
	return mas
}

//Фибоначи
func Fibo(n int) int {
	if n <= 2 {
		return 1
	} else {
		return Fibo(n-1) + Fibo(n-2)
	}
}

//фибоначи 2
func Fibo2(n int) *big.Int {
	var a = big.NewInt(1)
	var b = big.NewInt(1)
	var f = new(big.Int)
	for i := 1; i <= n; i++ {
		f = a.Add(a, b)
		a = b
		b = f
	}
	return f
}

//фибо 3
func Fibo3(n int) *big.Int {
	fi := (1 + math.Sqrt(5)) / 2
	f := (math.Pow(fi, float64(n)) / math.Sqrt(5)) + 0.5
	return big.NewInt(int64(f))
}

//фиб матрица !!!
// (1 1) x N
// (1 0)
