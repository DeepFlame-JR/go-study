package main

import "fmt"

func main(){
	for i:=0;i<5;i++{
		fmt.Println(i)
	}

	i := 0
	for{
		if i > 3{
			break
		}
		fmt.Println("Infinte")
		i++
	}

	loc := []string{"Seoul", "Busan", "Incheon"}
	for index, name := range loc{
		fmt.Println(index, name)
	}

	for i,j := 0,0 ; i<=10 ; i,j=i+1,j+10{
		fmt.Println(i, j)
	}


Loop1:
	for i:=0 ; i <=3 ; i++{
		for j:=0 ; j < 3 ; j++{
			if i == 1 && j == 1{
				break Loop1
			}
			fmt.Println(i, j)
		}
	}
}	

