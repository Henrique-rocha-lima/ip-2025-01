package main

import "fmt"

func main() {
	var media float64
	var nota1 float64
	var nota2 float64
	var nota3 float64
	fmt.Print("informe a primeira nota: ")
	fmt.Scanln(&nota1)
	if nota1 < 0 {
		fmt.Println("valor invalido")
		return
	}
	if nota1 > 10 {
		fmt.Println("valor invalido")
		return
	}
	fmt.Print("informe a segunda nota: ")
	fmt.Scanln(&nota2)
	if nota2 < 0 {
		fmt.Println("valor invalido")
		return
	}
	if nota2 > 10 {
		fmt.Println("valor invalido")
		return
	}
	fmt.Print("informe a terceira nota: ")
	fmt.Scanln(&nota3)
	if nota3 < 0 {
		fmt.Println("valor invalido")
		return
	}
	if nota3 > 10 {
		fmt.Println("valor invalido")
		return
	}

	media = (nota1 + nota2 + nota3) / 3

	fmt.Printf("\nSua media é de: %v\n", media)
	if media >= 6 {
		fmt.Println("Aprovado!")
	} else {
		fmt.Println("Reprovado")
	}
}
