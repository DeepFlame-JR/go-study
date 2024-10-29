package main

import "fmt"

func main(){
	// 열거형
	// 상수를 사용하는 일정한 규칙에 따라 숫자 계산
	const(
		Jan = 1
		Feb = 2
		Mar = 3
	)

	fmt.Println(Jan)
	fmt.Println(Feb)
	fmt.Println(Mar)

	fmt.Println("iota")
	const(
		_ = iota * 10 
		A 
		B
		C
	)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}

