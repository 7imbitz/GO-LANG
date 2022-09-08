package main

import (
    "fmt"
)

func main() {
    var x,y int 

    fmt.Scanln(&x)
    fmt.Scanln(&y)
    if (x >= -1000 && x <=1000 && y >= -1000 && y <=1000 && x != 0 && y != 0) {
        if (x >= 0) {
            if ( y >= 0 ) {fmt.Print("1")} else {fmt.Print("4")}
        } else {
            if ( y >= 0 ) {fmt.Print("2")} else {fmt.Print("3")}
        }
    }
}
