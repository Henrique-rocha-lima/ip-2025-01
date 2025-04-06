package main

import "fmt"

func main() {
	var h int
	fmt.Println("Quantas horas a carroça foi alugada")
	fmt.Scanln(&h)
	if h == 0 {
		fmt.Printf("R$25.00\n")
	}
	if h == 1 {
		fmt.Printf("R$20.00\n")
	}
	if h == 2 {
		fmt.Printf("R$15.00\n")
	}
	if h >= 3 {
		x := h / 3 * 10
		fmt.Printf("R$%d.00", x)
	}
	if h < 0 {
		fmt.Printf("Valor invalido")
	}
}
