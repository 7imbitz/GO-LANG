package main

import "fmt"

func main(){
	var H, M, tM, tH, newH, newM int

	fmt.Scan(&H, &M)
	if H == 0 {
		H =24
		tH = H * 60
		tM = tH + M
		tM -=45
		newH = tM / 60
		newM = (tM - (newH * 60)) % 60
	}else {
		tH = H * 60
		tM = tH + M
		tM -=45
		newH = tM / 60
		newM = (tM - (newH * 60))
	}

	fmt.Println(newH, " ", newM)
}
