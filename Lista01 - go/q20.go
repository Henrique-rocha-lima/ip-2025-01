package main

import "fmt"

func main() {
	var h int
	var m int
	var s int
	fmt.Println("escreva o tempo em horas, minutos e segundos nessa ordem")
	fmt.Scanln(&h)
	fmt.Scanln(&m)
	fmt.Scanln(&s)
	s2 := h * 3600
	s3 := m * 60
	s4 := s + s2 + s3
	fmt.Printf("O TEMPO EM SEGUNDOS E = %d \n", s4)
}
