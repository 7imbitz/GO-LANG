package main

import (
"fmt"
)
func main() {

	var n int
	fmt.Scanln(&n)

	var neg_tmp int = 0
	var temp int
	for i:=0;i<n;i++ {
		fmt.Scan(&temp)
		if (temp < 0) { neg_tmp++ }
	}

	fmt.Println(neg_tmp)
}
