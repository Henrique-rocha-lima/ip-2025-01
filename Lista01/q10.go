package main

import "fmt"

func main() {
	var a float64
	var b float64
	var c float64
	var d float64
	fmt.Println("Insira A, B, C e D nessa sequência:")
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)
	fmt.Scanln(&d)
	m := a*d - b*c
	fmt.Printf("O VALOR DO DETERMINANTE E = %.2f\n", m)
}
