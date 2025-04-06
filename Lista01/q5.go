package main

import "fmt"

func main() {
	var V float64
	var X string
	var Y int
	var Z float64
	fmt.Printf("Qual sua conta? ")
	fmt.Scanln(&Y)
	fmt.Printf("Qual seu consumo de agua em m³? ")
	fmt.Scanln(&Z)
	fmt.Printf("Você é consumidor residencial (r), comercial (c) ou industrial (i)? ")
	fmt.Scanln(&X)
	if X == "r" || X == "i" || X == "c" {
	} else {
		fmt.Println("invalido")
		return
	}
	fmt.Println("Conta = ", Y)
	if X == "r" {
		V = 5 + (0.05 * Z)
	}
	if X == "c" {
		if Z >= 80 {
			V = ((Z - 80) * 0.25) + 500
		} else {
			V = 500
		}
	}
	if X == "i" {
		if Z >= 80 {
			V = ((Z - 100) * 0.04) + 800
		} else {
			V = 800
		}
	}

	fmt.Print("Valor da conta = ", V)
}
