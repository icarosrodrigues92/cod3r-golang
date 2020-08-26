package main

import (
	"fmt"
	m "math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo (float64) foi inferido pelo compilardor. Duck!!!

	// forma reduzida de criar uma var
	area := PI * m.Pow(raio, 2)
	//se não usar variavel e não usar, da erro de compilação
	fmt.Println("A área da circuferência é", area)

	//Blocos de construção
	const (
		a = 1
		b = "dois"
	)

	var (
		c = 3.3
		d = true
	)
	fmt.Println(a, b, c, d)

	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "epa!"
	fmt.Println(g, h, i)
}
