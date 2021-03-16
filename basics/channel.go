package main

import (
	"fmt"
	"time"
)

func main() {
	//정수형 채널을 생성한다
	ch := make(chan int)

	go func(){
		ch <- 123
	} ()

	var i int
	i = <- ch //채널로부터 123을 받는다.
	fmt.Println(i)


	//-------------
	done := make(chan bool)
	go func(){
		for i:= 0; i < 10; i++ {
			fmt.Println(i)
		}
		done <- true
	} ()

	//위의 go루틴이 끝날 때까지 대기
	<-done


	//-------------
	//c := make(chan int)
	//c <- 1   //수신루틴이 없으므로 데드락 
	//fmt.Println(<-c) //코멘트해도 데드락 (별도의 Go루틴없기 때문)

	c := make(chan int, 1)
    c <- 101 //수신자가 없더라도 보낼 수 있다.
	fmt.Println(<-c)
	

	//-------------
	ch2 := make(chan string, 1)
    sendChan(ch2)
	receiveChan(ch2)
	

	//-------------
	ch3 := make(chan int, 2)
	ch3 <- 1
	ch3 <- 2

	close(ch3)

	// fmt.Println(<-ch3)
	// fmt.Println(<-ch3)
	// if _, success := <-ch3; !success {
	// 	fmt.Println("더이상 데이타 없음.")
	// }

	// 방법2
    // 위 표현과 동일한 채널 range 문
    for i := range ch3 {
        fmt.Println(i)
	}
	
	//--------------
	done1 := make(chan bool)
    done2 := make(chan bool)
 
    go run1(done1)
    go run2(done2)
 
EXIT:
    for {
        select {
        case <-done1:
            fmt.Println("run1 완료")
        case <-done2:
            fmt.Println("run2 완료")
			break EXIT
		// default:
		// 	fmt.Println("----")
        }
    }
}


func sendChan(ch chan<- string) {
	ch <- "Data"
}

func receiveChan(ch <-chan string) {
	data := ch
	fmt.Println(data)
}

func run1(done chan bool) {
    time.Sleep(1 * time.Second)
    done <- true
}
 
func run2(done chan bool) {
    time.Sleep(2 * time.Second)
    done <- true
}