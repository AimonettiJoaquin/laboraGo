package main

import "fmt"

func main() {
	var (
		nombre string = "Joaquin"
		edad   int    = 28
	)

	fmt.Println("Nombre: ", nombre)
	fmt.Println("Edad: ", edad)

	var (
		a              int = 2
		b              int = 5
		suma           int = a + b
		resta          int = a - b
		multiplicacion int = a * b
		division       int = a / b
	)
	fmt.Println("Suma: ", suma)
	fmt.Println("Resta: ", resta)
	fmt.Println("Multiplicacion: ", multiplicacion)
	fmt.Println("Division: ", division)
}
