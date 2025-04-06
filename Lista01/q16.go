package main

import "fmt"

func main() {
	var s float64
	var x float64
	fmt.Println("Qual o salario?")
	fmt.Scanln(&s)
	if s <= 300 {
		x = 1.5 * s
	} else {
		x = 1.3 * s
	}
	fmt.Printf("SALARIO COM REAJUSTE = %.2f", x)
}
