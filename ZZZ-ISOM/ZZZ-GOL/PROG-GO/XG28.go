
package main

import "fmt"

// Function to increment a value using a pointer
func increment(ptr *int) {
    *ptr = *ptr + 1 // Dereference the pointer and increment the value
}

func main() {
    // Declare a variable
    y := 10

    // Pass the address of y to the function
    increment(&y)

    // Print the updated value of y
    fmt.Println("Updated value of y:", y)

    x := 44

    increment(&x)

    // Print the updated value of y
    fmt.Println("Updated value of x:", x)


}