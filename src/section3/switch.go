package main

import "fmt"

func main(){
	a := 30/15

	switch a{
	case 2,4,6:
		fmt.Println("짝수")
		fallthrough // 아래 내용이 조건에 맞지 않아도 넘어감
	case 1,3,5:
		fmt.Println("홀수")
	case 100:
		fmt.Println("100이다")
	}

}	

