package main

import (
    "fmt"
    "math/big"
    "unsafe"
)

// Função para obter o tamanho em bytes de um valor
func getSizeInBytes(value interface{}) string {
    switch v := value.(type) {
    case uint64:
        return fmt.Sprintf("%d bytes", unsafe.Sizeof(v))
    case *big.Int:
        return fmt.Sprintf("%d bytes (big.Int)", len(v.Bytes()))
    default:
        return "unknown size"
    }
}

func main() {
    m := make(map[string]interface{})

    m["k1"] = uint64(7)
    m["k2"] = uint64(2147483648)
    m["k3"] = uint64(4294967296)
    m["k4"], _ = new(big.Int).SetString("18446744073709551616", 10)

    for key, value := range m {
        size := getSizeInBytes(value)
        fmt.Printf("%s: %v (%s)\n", key, value, size)
    }
}