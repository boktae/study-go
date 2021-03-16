package main

import (
	"fmt"
	"os"
)

func main() {
	 // 잘못된 파일명을 넣음
	 openFile("Invalid.txt")
 
	 // recover에 의해
	 // 이 문장 실행됨
	 println("Done") 
	
	 /*
	f, err := os.Open("1.txt")
	if err != nil {
		panic(err)
	}

	//main 마지막에 파일 close 실행
	defer f.Close()
	// 파일 읽기
    bytes := make([]byte, 1024)
    f.Read(bytes)
	fmt.Println(len(bytes))
	*/
}


func openFile(fn string) {
    // defer 함수. panic 호출시 실행됨
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("OPEN ERROR", r)
        }
    }()
 
    f, err := os.Open(fn)
    if err != nil {
        panic(err)
    }
 
    defer f.Close()
}