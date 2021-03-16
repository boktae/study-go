package main

import "fmt"

func main() {
	var a []int //슬라이스 변수 선언
	a = []int {1,2,3} //슬라이스에 리터럴값 지정
	fmt.Println(a, cap(a)); // 슬라이스의 길이 및 용량은 내장함수 len(), cap()을 써서 확인
	//-----------	
	s := make([]int, 5, 10) // 길이(Length)와 용량(Capacity)
	fmt.Println(len(s), cap(s))

	//-----------
	var s2 []int
    if s2 == nil {
        println("Nil Slice")
    }
	println(len(s), cap(s)) // 모두 0
	

	//-----------
	s3 := []int{0, 1, 2, 3, 4, 5}
    s3 = s3[2:5]  
	fmt.Println(s3) //2,3,4 출력
	
	//-----------
	s4 := []int{0, 1, 2, 3, 4, 5}
	s4 = s4[2:5]     // 2, 3, 4
	s4 = s4[1:]      // 3, 4
	fmt.Println(s4) // 3, 4 출력
	
	//-----------
	s5 := []int{0, 1}
	//하나 확장
	s5 = append(s5, 2)
	//복수 요소들 확장
	s5 = append(s5, 3,4,5,6)
	fmt.Println(s5)

	//-----------
	 // len=0, cap=3 인 슬라이스
	 sliceA := make([]int, 0, 3)
 
	 // 계속 한 요소씩 추가
	 for i := 1; i <= 15; i++ {
		 sliceA = append(sliceA, i)
		 // 슬라이스 길이와 용량 확인
		 fmt.Println(len(sliceA), cap(sliceA))
	 }
  
	 fmt.Println(sliceA) // 1 부터 15 까지 숫자 출력 

	 //-----------
	 sliceA2 := []int{1, 2, 3}
	 sliceB2 := []int{4, 5, 6}
  
	 sliceA2 = append(sliceA2, sliceB2...)
	 //sliceA = append(sliceA, 4, 5, 6)
  
	 fmt.Println(sliceA2) // [1 2 3 4 5 6] 출력

	 //-----------
	 source := []int{0, 1, 2}
	 target := make([]int, len(source), cap(source)*2)
	 copy(target, source)
	 fmt.Println(target)  // [0 1 2 ] 출력
	 println(len(target), cap(target)) // 3, 6 출력
}