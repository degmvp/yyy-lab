package main

import "fmt"

func main() {
    // Declare a variable
    x := 42

    // Get the memory address of x using the address-of operator (&)
    ptr := &x

    // Print the memory address
    fmt.Println("Memory address of x:", ptr)

    // Access the value stored at the memory address using the pointer indirection operator (*)
    fmt.Println("Value stored at the memory address:", *ptr)

    // Modify the value using the pointer
    *ptr = 100
    fmt.Println("New value of x:", x)
}