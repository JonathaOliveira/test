package main

import "fmt"

func main() {
	var inteiro = 10

	var decimal = 20.22

	var descricao = "teste string"

	var verdadeoufalso = false

	fmt.Printf("%v >>> (%T)\n", inteiro, inteiro)
	fmt.Printf("%v >>> (%T)\n", decimal, decimal)
	fmt.Printf("%v >>> (%T)\n", descricao, descricao)
	fmt.Printf("%v >>> (%T)\n", verdadeoufalso, verdadeoufalso)
}
