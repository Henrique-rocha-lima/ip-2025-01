package main

import "fmt"

func main() {
	var SalMin float64
	var preco float64
	var kW float64

	fmt.Print("insira o valor do salário mínimo: ")
	fmt.Scanln(&SalMin)
	fmt.Print("insira quantos kW a residência consome: ")
	fmt.Scanln(&kW)

	preco = 0.007 * SalMin
	cobrado := preco * kW
	desconto := cobrado * 0.9

	fmt.Printf("Custo por kW: R$ %.2f\n", preco)
	fmt.Printf("Custo do consumo: R$ %.2f\n", cobrado)
	fmt.Printf("Custo com desconto: R$ %.2f\n", desconto)
}
