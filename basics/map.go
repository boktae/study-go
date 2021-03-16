package main

import "fmt"

func main () {
	var m map[int]string
	m = make(map[int]string)

	 //추가 혹은 갱신
	 m[901] = "Apple"
	 m[134] = "Grape"
	 m[777] = "Tomato"

	  // 키에 대한 값 읽기
	  str := m[134]
	  fmt.Println(str)
   
	  noData := m[999] // 값이 없으면 nil 혹은 zero 리턴
	  fmt.Println(noData)
   
	  // 삭제
	  delete(m, 777)

	//---------------
	tickers := map[string]string{
        "GOOG": "Google Inc",
        "MSFT": "Microsoft",
        "FB":   "FaceBook",
        "AMZN": "Amazon",
    }
 
    // map 키 체크
    val, exists := tickers["MSFT"]
    if !exists {
        fmt.Println("No MSFT ticker")
	}
	fmt.Println(val);


    // for range 문을 사용하여 모든 맵 요소 출력
    // Map은 unordered 이므로 순서는 무작위
    for key, val := range tickers {
        fmt.Println(key, val)
    }
}