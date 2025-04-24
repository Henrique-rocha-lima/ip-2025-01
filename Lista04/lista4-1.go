package main

import "fmt"

func main() {
	var v [10]int
	s := 0
	for i := range v {
		fmt.Scanln(&v[i])
	}
	for i, val := range v {
		if val > 50 {
			fmt.Printf("%d posição: %d\n", val, i+1)
			s++
		}
	}
	if s == 0 {
		fmt.Println("nenhum numero maior que 50")
	}
}
