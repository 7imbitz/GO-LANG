package main 

import "fmt"

func main() {
    var str1,str2 string
    
    fmt.Scan(&str1)
    fmt.Scan(&str2)
    
    if (len(str1) < len(str2)) {
        fmt.Println("no")
    } else {
        fmt.Println("go")
    }
}
