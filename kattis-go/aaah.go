package main

import "fmt"

func main(){
	var str1 string
	var str2 string

	fmt.Scanln(&str1)
	fmt.Scanln(&str2)
	//str compare
	if len(str1) < len(str2){
		fmt.Println("no")
	}else {
		fmt.Println("go")
	}
}
