package main

import "fmt"

func main() {
	var v [10]int
	var pares []int
	var somaPares int
	var impares []int
	fmt.Println("Digite 10 números inteiros:")
	for i := range v {
		fmt.Scanln(&v[i])
		if v[i]%2 == 0 {
			pares = append(pares, v[i])
			somaPares += v[i]
		} else {
			impares = append(impares, v[i])
		}
	}
	fmt.Println("a) Números pares digitados:", pares)
	fmt.Println("b) Soma dos números pares:", somaPares)
	fmt.Println("c) Números ímpares digitados:", impares)
	fmt.Println("d) Quantidade de números ímpares:", len(impares))
}
