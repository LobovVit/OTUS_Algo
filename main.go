package main

import (
	"fmt"
	"log"
	"otus_algo/task03"
)

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
	var a int = 70
	//fmt.Println(task03.Fibo(a+2))
	fmt.Println(task03.Fibo2(a))
	fmt.Println(task03.Fibo3(a + 2))

	fmt.Println("")
	log.Println("main")

}
