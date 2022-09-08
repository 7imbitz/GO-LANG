package main

import "fmt"

func main() {
    var r1,s,r2 int

    fmt.Scan(&r1,&s)
    if (r1 >= -1000 && r1 <= 1000 && s >= -1000 && s <= 1000) {
        r2 = s * 2 - r1

    }
    fmt.Print(r2)
}
