package main

import (
	"fmt"
	"runtime"
	"sync"
)

func say(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(s, "***", i)
	}
}

func main() {
	/*
		say("sync")

		go say("Async1")
		go say("Async2")
		go say("Async3")

		time.Sleep(time.Second * 3)
	*/

	// 4개의 CPU 사용
	runtime.GOMAXPROCS(4)

	// WaitGroup 생성. 2개의 go루틴을 기다림
	var wait sync.WaitGroup
	wait.Add(2)

	go func() {
		defer wait.Done() // 끝나면 Done 호출
		fmt.Println("Hello")
	}()

	go func(msg string) {
		defer wait.Done()
		fmt.Println(msg)
	}("Hi")

	wait.Wait()
}
