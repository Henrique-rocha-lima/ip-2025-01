package main

import "fmt"

func main() {
	var n int
	fmt.Println("Quantas medidas quer converter? ")
	fmt.Scanln(&n)
	f := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&f[i])
	}
	for i := 0; i < n; i++ {
		c := (5 * (f[i] - 32)) / 9
		fmt.Printf("%.2f fahrenheit equivale a %.2f celsius\n", f[i], c)
	}
}
