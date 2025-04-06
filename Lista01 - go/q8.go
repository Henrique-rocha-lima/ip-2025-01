package main

import "fmt"

func main() {
	const π = 3.14159
	var h float64
	var r float64
	fmt.Println("insira o raio em metros:")
	fmt.Scanln(&r)
	fmt.Println("insira a altura em metros:")
	fmt.Scanln(&h)
	x := h * 2 * π * r
	y := π * r * r
	preco := (x + (y * 2)) * 100
	fmt.Printf("O VALOR DO CUSTO E = %.2f", preco)
}
