package main

import "fmt"

func main() {
    // Estrutura anônima com diferentes tipos de campos
    person := struct {
        name    string
        age     int
        height  float64
        isAdmin bool
    }{
        name:    "Ash Ketchum",
        age:     10,
        height: 1.65,
        isAdmin: false,
    }

    // Exibe a estrutura
    fmt.Println(person)
}