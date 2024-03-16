package main

import "fmt"

func main() {
	var a int
	var b float64
	var c bool
	var d string
	var e *int

	// %q imprime string vazia
	// esses são os valores zeros das variáveis da linguagem go
	fmt.Printf("%v %v %v %q %v", a, b, c, d, e)
}