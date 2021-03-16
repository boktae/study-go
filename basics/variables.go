package main

import "fmt"

func main() {
	var a int
	var b float32 = 11.0
	var i, j, k int = 1,2,3

	fmt.Println(a, b, i, j, k)

	const c int = 10
	const s string  ="Hi"
	s2 := "Hi2"

	fmt.Println(c, s, s2);

	const (
		Visa = "Visa"
		Master = "MasterCard"
		Amex = "American Express"
	)
	
	fmt.Println(Visa, Master, Amex)

	const (
		Apple = iota
		Grape
		Orange
	)

	fmt.Println(Apple, Grape, Orange)
}