package main

import "fmt"

func main() {
    var n,x int 

    fmt.Scan(&n)
    if (n >=1 && n <=20) {
        for k := 1; k <= n; k++ {
            fmt.Scan(&x)
            if (x >= -10 && x <= 10) {
                if (x % 2 == 0){
                    fmt.Println(x, " is even")
                } else {
                    fmt.Println(x, " is odd")
                }
            } else {
                fmt.Println(x, " is out of range")
            }
        }
    }
}
