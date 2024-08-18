package main

import (
	"fmt"
)

func main() {
	var myVar string = "original message"
	// Puntero: variable que almacena la direccion de memoria de otra variable
	var puntero *string = &myVar
	fmt.Println("Valor de myVar: ", myVar)
	fmt.Println("Direccion de memoria de myVar: ", &myVar)
	fmt.Println("Valor de puntero: ", puntero)
	fmt.Println("Direccion de memoria de puntero: ", &puntero)

	// Dereferenciar el puntero *nombrePuntero
	fmt.Println("Valor al que apunta puntero: ", *puntero)
	*puntero = "Nuevo mensaje"
	fmt.Println("Valor de myVar: ", myVar)
}
