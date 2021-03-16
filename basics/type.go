package main

import "fmt"

func main() {
	/*
		부울 타입:	bool
		문자열: string
		정수형 타입: int int8 int16  int32 int64
				  uint uin8 uint16 uint32 uint64 uintptr
		Float 및 복수타입: float32 float64 complex64 complex128
		기타 타입: byte = uint8
				rune = int32
	*/

	//복수라인 문자열 (문자열을 별도로 해석하지 않음)
	rawLiteral := `아리랑\n
	아리랑\n
	아라리요`
	fmt.Println(rawLiteral);

	interLiteral := "아리랑아리랑\n아리리요"
	fmt.Println(interLiteral);

	//암묵적(implicit)변환이 일어나지 않으므로 반드시 지정해 주어야 함.
	var i int = 10
	var u uint = uint(i)
	var f float32 = float32(u)
	fmt.Println(i, u, f);

	str := "ABC"
	bytes := []byte(str)
	str2 := string(bytes)
	
	fmt.Println(str, bytes, str2)
}