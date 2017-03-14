package main

import "fmt"

func main(){
	var str string
	var n int
	var p int 

	fmt.Scanln(&n, &p)
	for i := 0; i < n; i++{
		fmt.Scanln(&str)
	}

	fmt.Println(p)
}
