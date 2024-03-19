package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))
	
	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// cuidado...
	fmt.Println("Teste " + string(97))

	// int para string
	fmt.Println("Teste " + strconv.Itoa(123))

	// string para int, num é o primeiro valor e o segundo valro colocamos _ para ignorar o valor, neste caso colocaria err 
	// mas como em go nao podemos criar sem usar, podemos ignorar  
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("1") // ou
	// b, _ := strconv.ParseBool("true")  
	if b{
		fmt.Println("Verdadeiro")
	}

	fmt.Println()
	fmt.Println()
}
