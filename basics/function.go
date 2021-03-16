package main

import "fmt"

func main () {
	msg := "Hello"

	say(msg)
	say2(&msg)
	println(msg) //변경된 메시지 출력
	
	say3("This", "is", "a", "book")

	total := sum(1, 7, 3, 5, 9)
	println(total)
	
	count, total := sum2(1, 7, 3, 5, 9)
	println(count, total)  
	
	count, total = sum3(1, 7, 3, 5, 9)
    println(count, total)  
}

//Pass By Value
func say(msg string) {
	fmt.Println(msg)
}

//Pass By Reference
func say2(msg *string) {
	fmt.Println(*msg);
	*msg = "Changed" //메시지 변경
}

//Variadic function
func say3(msg ...string) {
    for _, s := range msg {
        println(s)
    }
}

//return
func sum(nums ...int) int {
    s := 0
    for _, n := range nums {
        s += n
    }
    return s
}

func sum2(nums ...int) (int, int) {
	s := 0
	count := 0

	for _, n:= range nums {
		s += n
		count++
	}
	return count, s
}

// Named Return Parameter
func sum3(nums ...int) (count int, total int) {
    for _, n := range nums {
        total += n
    }
    count = len(nums)
    return
}
