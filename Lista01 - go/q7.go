package main

import "fmt"

func main() {
	var f float64
	var p float64
	var c float64
	var mm float64
	fmt.Println("Insira a temperatura em fahrenheit:")
	fmt.Scanln(&f)
	fmt.Println("Insira a quantidade de chuva em polegadas:")
	fmt.Scanln(&p)
	c = (5 * (f - 32)) / 9
	fmt.Printf("O VALOR EM CELSIUS = %.2f\n", c)
	mm = 25.4 * p
	fmt.Printf("A QUANTIDADE DE CHUVA E = %.2f", mm)
}
