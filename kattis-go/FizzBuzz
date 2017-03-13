package main

import "fmt"

func main(){
	var x int
	var y int
	var n int 

	fmt.Scanln(&x, &y, &n)
	if x >=1 && x<y && y<=n && n <= 100{
		for i := 1; i <= n; i++{
			if i % x == 0 && i % y == 0 {
				fmt.Println("FizzBuzz")
			} else if i % x == 0{
				fmt.Println("Fizz")
			} else if i % y == 0  {
				fmt.Println("Buzz")
			} else {
				fmt.Println(i)
			}
		}	
	}
}
