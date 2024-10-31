package lib

import(
	"fmt"
)


func init(){
	// 패키지 로드시에 가장 먼저 호출 
	// main보다 먼저 실행 됨
	fmt.Println("lib function init")
  }
  

// 대문자: 패키지 외부 접근 가능
// 소문자: 패키지 외부 접근 불가능
func CheckNum(c int32) bool {
	return c > 10
}
