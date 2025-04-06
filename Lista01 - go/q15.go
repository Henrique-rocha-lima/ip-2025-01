package main

import "fmt"

func main() {
	var n int
	var x int
	fmt.Println("insira um valor")
	fmt.Scanln(&n)
	if n < 6 || n > 1999 {
		fmt.Println("valor invalido")
		return
	}
	for i := 1; i <= n; i++ {
		if i%2 != 0 {
			continue
		}
		x = i * i
		fmt.Printf("%d^2 = %d\n", i, x)
	}
}
