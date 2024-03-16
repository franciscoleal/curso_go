package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// números inteiros
	fmt.Println(1, 2, 1000)

	// mostrar o tipo da variável com reflect
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))
	// 							  byte, short 2 bytes, int 4 bytes, long 8 bytes 
	// sem sinal (só positivos... uint8 uint16 uint32 uint64)
	var b byte = 255 // byte é um alias de uint8
	fmt.Println("O byte é", reflect.TypeOf(b))

	// com sinal... int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))
	
	// tabela unicode
	var i2 rune = 'a' // representa um mapeamento da tabela Unicode (int32), rune seria um alias para o int32
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	// números reais (float32, float64)
	var x float32 = 49.99
	var y = float32(49.99)
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo de y é", reflect.TypeOf(y))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99))

	// boolean
	bo := true // nao aceita valores inteiros como 1 ou 0, será interpretado como int mesmo e nao um boolean
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string - somente aspas duplas
	s1 := "Olá meu nome é Francisco"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1)) //len mostra o tamanho da string

	// string com multiplas linhas com backtick
	s2 := `Olá
	meu
	nome
	é
	Francisco`
	fmt.Println("O tamanho da string é", len(s2))

	// char???
	// var x char = 'b'
	char := 'a' // vai ser um tipo int32, então nao tem o tipo char
	fmt.Println("O tipo de char é", reflect.TypeOf(char))
	fmt.Println(char) // vai mostrar o unicode da letra a 
}