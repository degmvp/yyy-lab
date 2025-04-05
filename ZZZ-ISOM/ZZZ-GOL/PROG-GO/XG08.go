package main

import (
    "fmt"

)

func main() {
    a := 20.45
    b := 34.89

    // Subtraction of two floating-point numbers
    c := b - a

    var V1 int64 = 2147483648
    var V2 uint32 = 2147483648

    // Perform subtraction on V1 and V2
    V1 = V1 - 1073741824
    V2 = V2 - 1073741824

    // Display the result
    fmt.Printf("Result c  is: %f\n", c)
	fmt.Printf("Result v1 is: %d\n", V1)
	fmt.Printf("Result v2 is: %d\n", V2)

    // Display the type of c variable
    fmt.Printf("The type of c is: %T\n", c)
    fmt.Printf("The type of V1 is: %T\n", V1)
    fmt.Printf("The type of V2 is: %T\n", V2)
}
