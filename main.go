package main

import (
	"fmt"
	"log"
	"math/rand"
	"otus_algo/task04"
)

func generateMass(len int) []int {
	ret := make([]int, len)
	for i := 0; i < len; i++ {
		ret[i] = rand.Intn(len * 10)
	}
	return ret
}

func main() {
	//
	//var a = 1234567890
	//var b = 123
	//res := task03.Fcalc(a, b)
	//log.Println(res)
	//
	//a = 1234567890
	//b = 123
	//res = task03.Fcalc2(a, b)
	//log.Println(res)

	//просты числа - решето
	//var a int = 1000
	//res := task03.Feratob(a)
	//res2 := task03.Ferato(a)
	//for i, _ := range res{
	//		fmt.Print(res[i])
	//	    fmt.Print("=")
	//	    fmt.Print(res2[i])
	//	    fmt.Println(";")
	//
	//}

	//Фибоначи
	//var a int = 70
	////fmt.Println(task03.Fibo(a+2))
	//fmt.Println(task03.Fibo2(a))
	//fmt.Println(task03.Fibo3(a + 2))

	//сортировка
	//var a[] int = make([]int, 10)
	//var a  = []int{1,200,3,4,5,700,4,100}

	var a = generateMass(200)
	task04.SortShell(a)
	for idx, val := range a {
		fmt.Printf("idx = %d val=%d \n", idx, val)
	}

	fmt.Println("")
	log.Println("main")

}
