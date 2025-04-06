package main

import "fmt"

func main() {
	var x int
	fmt.Println("Insira o numero:")
	fmt.Scanln(&x)
	if x%3 == 0 && x%5 == 0 {
		fmt.Println("O NUMERO E DIVISIVEL\n")
	} else {
		fmt.Println("O NUMERO NAO E DIVISIVEL\n")
	}
}
