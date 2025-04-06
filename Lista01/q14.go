package main

import "fmt"

func main() {
	var a float64
	var h float64
	fmt.Println("Qual a altura da piramide em metros?")
	fmt.Scanln(&h)
	fmt.Println("Qual o comprimento de 1 das arestas da base em metros?")
	fmt.Scanln(&a)
	V := 0.86602540378443864676372317075294 * a * a * h
	fmt.Printf("O VOLUME DA PIRAMIDE E = %.2f METROS CUBICOS\n", V)
}
