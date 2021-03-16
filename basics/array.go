package main

import "fmt"

func main() {
	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	fmt.Println(a[1])

	var a1 = [3]int{1,2,3}
	var a2 = [...]int{1,2,3}

	fmt.Println(a1[2]);
	fmt.Println(a2[2]);

	var a3 = [2][3]int{
        {1, 2, 3},
        {4, 5, 6},  //끝에 콤마 추가
    }
    println(a3[1][2])
}