package main

import "fmt"

func main(){
	// 반드시 Boolean 검사
	// 소괄호 사용 x

	var a int = 20
	b := 20
	if a >= 15{ // 끝을 중괄호로 끝내야 함 (안 붙이면 ;가 붙는 것으로 가정되어 에러 발생)
		fmt.Println("15이상")
	}

	if b >= 25{
		fmt.Println("25이상")
	}else{
		fmt.Println("25미만")
	}

	if c := true; c{ // c 변수는 여기서만 사용 가능
		fmt.Println("True")
	}
}

