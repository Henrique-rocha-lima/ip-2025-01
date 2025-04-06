package main

import "fmt"

func main() {
	var x int
	var y int
	fmt.Println("Insira X e Y")
	fmt.Scanf("%d %d", &x, &y)
	if x%2 != 0 {
		fmt.Println("o primeiro numero não é par")
		return
	}

	for i := 0; i < y; i++ {
		z := i*2 + x
		fmt.Printf("%d ", z)
	}
}
