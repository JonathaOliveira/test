package main

import (
	"fmt"
)

func calcularIMC(peso, altura float64) string {
	imc := peso / (altura * altura)

	if imc < 18.5 {
		return "Abaixo do peso"
	} else if imc >= 18.5 && imc < 25 {
		return "Peso normal"
	} else if imc >= 25 && imc < 30 {
		return "Sobrepeso"
	} else if imc >= 30 && imc < 35 {
		return "Obesidade de nível 1"
	} else {
		return "Obesidade de nível 2"
	}
}

func main() {
	peso := 70.0   // em kg
	altura := 1.75 // em metros

	resultadoIMC := calcularIMC(peso, altura)
	fmt.Println("O seu IMC é:", resultadoIMC)
}
