package task03

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
