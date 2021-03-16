package main

import "fmt"

type person struct {
	name string
	age int
}

type dict struct {
	data map[int]string
}

//생성자함수 정의
func newDict() *dict {
	d := dict{}
	d.data = map[int]string{}
	return &d //포인터 전달
}

func main() {
	p := person{}

	p.name = "Lee"
	p.age = 10

	fmt.Println(p)

	var p1 person 
	p1 = person{"Bob", 20}
	p2 := person{name: "Sean", age: 50}
	fmt.Println(p1)
	fmt.Println(p2)

	//-------------
	dic := newDict()//생성자 호출
	dic.data[1] = "A"
	fmt.Println(dic);
}