package main

import "fmt"

func main() {
    var VAR1 int64 = 2147483648
    var VAR2 uint32 = 2147483648

    fmt.Println(VAR1) // 2147483648
    fmt.Println(VAR2) // 2147483648
}