package main

import "fmt"

func main() {
	var A float64
	var B float64
	var C float64
	fmt.Println("Insira A:")
	fmt.Scanln(&A)
	fmt.Println("Insira B:")
	fmt.Scanln(&B)
	fmt.Println("Insira C:")
	fmt.Scanln(&C)
	delta := B*B - 4*A*C
	fmt.Printf("O VALOR DE DELTA E = %.2f\n", delta)
}
