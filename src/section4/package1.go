package main

import(
	"fmt"
	checkUp "go-study/src/section4/lib"  // GOPATH 설정필요
)

func init(){
  // 패키지 로드시에 가장 먼저 호출 
  // main보다 먼저 실행 됨
  fmt.Println("Init Method start!")
}

func init(){
	fmt.Println("Init Method start! 222")
  }
  

func main(){
	fmt.Println("10보다 큰 수?: ", checkUp.CheckNum(4))
}