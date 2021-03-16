package main

import "fmt"

func main() {
	var k int = 1

	if k == 1 {
		fmt.Println("One")
	} else if k == 2 {  //같은 라인
		fmt.Println("Two")
	} else {   //같은 라인
		fmt.Println("Other")
	}

	if val := k * 2; val < 3 {
		fmt.Println(val)
	}

	var name string
	var category = 1

	switch category {
	case 1:
		name = "Paper Book"
	case 2:
		name = "eBook"
	case 3:
		name ="Blog"
	default:
		name = "Other"
	}

	fmt.Println(name)

	check(2)
}

func check(val int) {
	switch val {
	case 1:
		fmt.Println("1 이하")
		fallthrough
	case 2:
		fmt.Println("2 이하")
		fallthrough
	case 3:
		fmt.Println("3 이하")
		fallthrough
	default:
		fmt.Println("default 도달")
	}
}