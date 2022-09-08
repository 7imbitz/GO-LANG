
package main

import "fmt"

func main() {
    var x int
    
    fmt.Scanln(&x)
    if (x >= 1 && x <= 100) {
        for i := 1; i <= x; i++ {
            fmt.Println(i, "Abracadabra")
        }
    }
}
