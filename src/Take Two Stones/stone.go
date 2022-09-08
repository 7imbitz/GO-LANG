package main

import "fmt"

func main() {
    var n int

    fmt.Scan(&n)
    if (n >= 1 && n <= 10000000 ) {
        if (n % 2 == 0 ) {
            fmt.Print("Bob")
        } else {
            fmt.Print ("Alice")
        }
    }
}
