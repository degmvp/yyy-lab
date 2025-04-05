package main

import (
	"fmt"
	"math"
)

func main() {
	var i int
	var j int = -1
	var k, l int = 1, 2

	m := 10
	n, o := 20, 30

	fmt.Printf("i: %d, j: %d, k: %d, l: %d, m: %d, n: %d, o: %d\n", i, j, k, l, m, n, o)

	var quadrado func(int) int = criaPotencia(2)

	fmt.Printf("2 elevado ao quadrado: %d\n", quadrado(8))
	fmt.Printf("2 elevado a 16: %d\n", math.Pow(2,16))
}
func criaPotencia(exp int) func(int) int {
	return func(base int) int {
		return int(math.Pow(float64(base), float64(exp)))
	}
}
