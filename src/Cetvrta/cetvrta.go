package main

import "fmt"

func main() {
	count_x := make(map[int]int)
	count_y := make(map[int]int)

	for i := 0; i < 3; i++ {
		var x, y int
		fmt.Scanf("%d %d", &x, &y)
		count_x[x]++
		count_y[y]++
	}

	for key,value := range count_x {
		if value == 1 {
			fmt.Printf("%d ", key)
			break
		}
	}

	for key,value := range count_y {
		if value == 1 {
			fmt.Println(key)
			break
		}
	}
}
