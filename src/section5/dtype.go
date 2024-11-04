package main

import (
	"fmt"
	"math"
	"strings"
	"unicode/utf8"
)

func main(){
	// Bool
	// !, || (or), && (and)
	// 암묵적 변환을 지원하지 않음

	// Integer
	// 정수, 실수, 복소수
	// 32bit, 64bit, unsigned (양수)
	// 정수: 8진수(0) 16진수(0x), 10진수
	var num1 int = 17
	var num2 int = 0x32fa2c

	fmt.Println("num1: ", num1)
	fmt.Println("num2: ", num2)

	var char1 byte = 72
	var char2 byte = 0110
	var char3 byte = 0x48
	fmt.Printf("%c %c %c\n", char1, char2, char3)
	fmt.Printf("%d %d %d\n", char1, char2, char3)
	fmt.Printf("%d %o %x\n", char1, char2, char3) // 10진수, 8진수, 16진수

	// 숫자 연산
	// 타입이 같아야 가능
	// +, -, *, %, /, <<, >>, &, ^
	fmt.Println("maxInt8: ", math.MaxInt8)
	fmt.Println("maxInt16: ", math.MaxInt16)
	fmt.Println("maxInt32: ", math.MaxInt32)
	fmt.Println("maxInt64: ", math.MaxInt64)

	n1 := 100000 // int
	n2 := int16(10000)
	fmt.Println("ex: ", n1+int(n2))


	// 문자열
	// "", ``
	// 문자 char '' 타입이 존재하지 않음 -> rune(int32)로 문자 코드 값을 표현
	// 문자: '' 로 작성
	// 자주 사용하는 escape: \\, \', \"
	str1 := "c:\\go_study\\"
	str2 := `c:\go_sutdy\`
	fmt.Println("str1: ", str1)
	fmt.Println("str2: ", str2)
	fmt.Println("str1 length:", utf8.RuneCountInString(str1))
	fmt.Println("str1 length:", len([]rune(str1)))

	str3 := "Golang"
	for i, char := range str3{
		fmt.Printf("%c(%d)\t", char, i)
	}
	fmt.Println()

	fmt.Println("Golang[0:2]", str3[0:2])

	str4 := "World"
	strSet := []string{}
	strSet = append(strSet, str3)
	strSet = append(strSet, str4)

	fmt.Println("Join: ", strings.Join(strSet, " || "))


}