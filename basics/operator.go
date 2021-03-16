package main

import "fmt"

func main() {
	/*
		관계연산자
		a == b
		a != c
		a >= b
	*/

	/*
		논리연산자
		A && B
		A || !(C && B)
	*/

	/*
		Bitwise 연산자
		c = (a & b) << 5
	*/

	/*
		할당연산자
		a = 100
		a *= 10
		a >>= 2
		a |= 1
	*/

	//	포인터연산자
	var k int = 10
	var p = &k
	fmt.Println(k, &k, *p);
	k = 20
	fmt.Println(k, &k, *p);
}
