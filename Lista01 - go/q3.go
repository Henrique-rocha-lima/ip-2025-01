package main

import "fmt"

func main() {
	var n1 int
	var n2 int
	var n3 int
	var junt int
	var x int

	fmt.Printf("Insira n1: ")
	fmt.Scanln(&n1)
	if n1 >= 10 || n1 < 0 {
		fmt.Println("Valor invalido")
		return
	}

	fmt.Printf("Insira n2: ")
	fmt.Scanln(&n2)
	if n2 >= 10 || n2 < 0 {
		fmt.Println("Valor invalido")
		return
	}

	fmt.Printf("Insira n3: ")
	fmt.Scanln(&n3)
	if n3 >= 10 || n3 < 0 {
		fmt.Println("Valor invalido")
		return
	}

	junt = n1*100 + n2*10 + n3

	fmt.Println(junt)

	x = junt * junt

	fmt.Println(x)
}
