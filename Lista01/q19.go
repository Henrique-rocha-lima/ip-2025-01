package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Println("Digite um número:")
	fmt.Scanln(&n)

	if n <= 1 {
		fmt.Println("Número inválido!")
		return
	}

	var soma float64 = 0

	for i := 1; i <= n; i++ {
		soma += 1.0 / float64(i)
	}

	fmt.Printf("%f\n", soma)
}
