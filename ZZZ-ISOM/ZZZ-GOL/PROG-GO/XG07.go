package main

import (
    "fmt"
    "sync"
)

func main() {
    // Cria um canal para receber os resultados
    results := make(chan int)

    // Cria um WaitGroup para sincronizar as goroutines
    var wg sync.WaitGroup

    // Números para calcular o quadrado
    numbers := []int{1, 2, 3, 4, 5}

    // Inicia uma goroutine para cada número
    for _, num := range numbers {
        wg.Add(1) // Incrementa o contador do WaitGroup
        go func(n int) {
            defer wg.Done() // Decrementa o contador ao finalizar
            results <- n * n // Envia o quadrado para o canal
        }(num)
    }

    // Goroutine para fechar o canal após todas as goroutines terminarem
    go func() {
        wg.Wait() // Espera todas as goroutines finalizarem
        close(results) // Fecha o canal
    }()

    // Imprime os resultados recebidos do canal
    for result := range results {
        fmt.Println("Quadrado:", result)
    }
}