package main

import "fmt"

func main(){
	for i:=1; i <= 9; i=i+2 {
		fmt.Println(i, "X", i, "=", i*i)
	}
}

//for i:=1; i <= 5; i++ {
//	for j:=1; j <= i; j++{
//		fmt.Print("*")
//	}
//	fmt.Println("")
//}

//for i:=2; i <= 9; i++ {
//	for j:=1; j<=9; j++{
//		fmt.Println(i, "x", j, "=", i*j)
//	}
//}