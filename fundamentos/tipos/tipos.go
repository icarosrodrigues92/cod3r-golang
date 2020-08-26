package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//númericos inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	//sem sinal (só positivos)...
	// uint8-byte uint16-short uint32-int uint64-long
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	i1 := math.MaxInt64
	fmt.Println("O máximo de int é ", i1)
	fmt.Println("O tipo i1 é ", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento da tabla unicode (int32)
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	//números reais (float32, float64). Por padrão é 64 a não ser que você defina
	var x float32 = 49.99
	fmt.Println("O rune é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99))

	//boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string
	s1 := "Olá meu nome é Ícaro"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1))

	// string com multiplas linhas
	s2 := `Olá
	meu
	nome
	é
		ícaro`
	fmt.Println(s2 + "!")
	fmt.Println("O tamanho da string é", len(s2))

	//char??? não existe
	char := 'b'
	fmt.Println("O tipo de char é", reflect.TypeOf(char))
	fmt.Println(char)
}
