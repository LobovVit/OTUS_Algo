package main

import (
	"log"
	"otus_algo/task03"
)

func main() {
	var a = 1234567890
	var b = 123
	res := task03.Fcalc(a, b)
	log.Println(res)

	a = 1234567890
	b = 123
	res = task03.Fcalc2(a, b)
	log.Println(res)

	log.Println("main")
}
