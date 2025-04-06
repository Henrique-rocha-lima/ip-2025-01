package main

import "fmt"

func main() {
	var x int
	var y int
	var z int
	fmt.Println("Insira o valor inicial, a razão e a quantidade de argumentos a serem somados (separados por espaço):")
	fmt.Scanf("%d %d %d", &x, &y, &z)
	k := make([]int, z)
	var soma int = 0
	for i := 0; i < z; i++ {
		k[i] = int(x + i*y)
		soma += k[i]
	}
	fmt.Printf("%d", soma)
}
