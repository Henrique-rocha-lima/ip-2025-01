package main

import "fmt"

func main() {
	var total float64
	var popular float64
	var geral float64
	var arquibancada float64
	var cadeiras float64

	fmt.Print("Informe quanto ingressos para popular foram vendidos:")
	fmt.Scanln(&popular)
	if popular < 0 {
		fmt.Print("valor invalido")
		return
	}

	fmt.Print("Informe quanto ingressos para geral foram vendidos:")
	fmt.Scanln(&geral)
	if geral < 0 {
		fmt.Print("valor invalido")
		return
	}

	fmt.Print("Informe quanto ingressos para arquibancada foram vendidos:")
	fmt.Scanln(&arquibancada)
	if arquibancada < 0 {
		fmt.Print("valor invalido")
		return
	}

	fmt.Print("Informe quanto ingressos para cadeiras foram vendidos:")
	fmt.Scanln(&cadeiras)
	if cadeiras < 0 {
		fmt.Print("valor invalido")
		return
	}

	total = popular + geral + arquibancada + cadeiras

	PorcentoPopular := popular / total * 100
	PorcentoGeral := geral / total * 100
	PorcentoArquibancada := arquibancada / total * 100
	PorcentoCadeiras := cadeiras / total * 100

	fmt.Printf("A porcentagem de popular vendida foi: %f \nA porcentagem de geral vendida foi: %f \nA porcentagem de arquibancada vendida foi: %f \nA porcentagem de cadeiras vendida foi: %f \n", PorcentoPopular, PorcentoGeral, PorcentoArquibancada, PorcentoCadeiras)

	totalvendas := popular*1 + geral*5 + arquibancada*10 + cadeiras*20

	fmt.Printf("O valor total arrecadado foi: %f", totalvendas)
}
