package main

import "fmt"

func main(){
	var arr [6]int
	cor := [6]int{ 1,1,2,2,2,8, }

	for i := 0; i < 6; i++{
		fmt.Scan(&arr[i])
		fmt.Print(cor[i] - arr[i], " ")
	}
}
