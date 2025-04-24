package main

import "fmt"

func main() {
	var v1 [10]int
	var v2 [5]int
	var paresResultantes []int
	var imparesResultantes []int
	fmt.Println("Digite 10 números para o primeiro vetor:")
	for i := range v1 {
		fmt.Scanln(&v1[i])
	}
	fmt.Println("Digite 5 números para o segundo vetor:")
	for i := range v2 {
		fmt.Scanln(&v2[i])
	}
	somaV2 := 0
	for _, val := range v2 {
		somaV2 += val
	}
	for _, val := range v1 {
		if val%2 == 0 {
			paresResultantes = append(paresResultantes, val+somaV2)
		} else {
			imparesResultantes = append(imparesResultantes, val+somaV2)
		}
	}
	fmt.Println("Vetor resultante 1:", paresResultantes)
	fmt.Println("Vetor resultante 2:", imparesResultantes)
}
