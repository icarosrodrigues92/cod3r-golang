package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" Linha.")

	//comando anterior não quebrou a linha. Só quebra depois de printar.
	fmt.Println(" Nova")
	fmt.Println("linha.")

	x := 3.131516
	fmt.Println("O valor de x é", x)
	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é" + xs)
	fmt.Printf(" O valor de x é %.2f", x)

	a := 1
	b := 1.999
	c := false
	d := "opa"
	// /n = pula linha, %s = string, %f float, %f.X float com X decimais, %t boolean
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n %v %v %v %v", a, b, c, d)
}
