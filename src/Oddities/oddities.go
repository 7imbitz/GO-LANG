package main

import "fmt"

func main(){
	var n int
	var x int

	fmt.Scanln(&n)
	for i := 1; i <= n; i++{
		fmt.Scanln(&x)
		if x >= -10 && x <=10 {
			if x % 2 == 0 { 
				fmt.Println(x, " is even")
			} else {
				fmt.Println(x, " is odd")
			}	
		} else {
			fmt.Println(x, "wrong integer")
		}
	}
}
