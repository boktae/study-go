package main

import "fmt"

type Rect struct {
	width, height int
}
// Value Receiver
func (r Rect) area() int {
	return r.width * r.height
}
// Pointer Receiver
func (r *Rect) area2() int {
    r.width++
    return r.width * r.height
}

func main (){
	rect := Rect{10, 20}
	area := rect.area()
	fmt.Println(rect.width, area);

	rect2 := Rect{10, 20}
	area2 := rect2.area2() //메서드 호출
    fmt.Println(rect2.width, area2) // 11 220 출력
}