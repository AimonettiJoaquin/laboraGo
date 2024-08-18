/*
1. Multiplicación concurrente de matrices

Descripción:

Implementa un programa en Go que simula la ejecución de múltiples tareas en paralelo utilizando goroutines. Cada tarea realizará un trabajo específico (por ejemplo, procesamiento de datos, cálculos complejos, etc.). Utiliza canales para comunicar el progreso de cada tarea y sincronizar la finalización de todas las goroutines.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generarMatriz(filas, columnas int) [][]int {
	matriz := make([][]int, filas)
	for i := 0; i < filas; i++ {
		matriz[i] = make([]int, columnas)
		for j := 0; j < columnas; j++ {
			matriz[i][j] = rand.Intn(10)
		}
	}
	return matriz
}

// Funcion de imprimir la matriz

func imprimirMatriz(matriz [][]int) {

	for _, fila := range matriz {
		fmt.Println(fila)
	}
}

//Funcion de multiplicacion

func multiplicarMatricesConcurrente(matrizA, matrizB [][]int, resultado chan<- [][]int) {
	filasA := len(matrizA)
	columnasA := len(matrizA[0])
	columnasB := len(matrizB[0])

	resultadoMatriz := make([][]int, filasA)
	for i := 0; i < filasA; i++ {
		resultadoMatriz[i] = make([]int, columnasB)
	}

	for i := 0; i < filasA; i++ {
		for j := 0; j < columnasB; j++ {
			go func(i, j int) {
				suma := 0
				for k := 0; k < columnasA; k++ {
					suma += matrizA[i][k] * matrizB[k][j]
				}
				resultadoMatriz[i][j] = suma
			}(i, j)
		}
	}
	for x := 0; x < filasA; x++ {
		for z := 0; z < columnasB; z++ {
			<-time.After(10 * time.Millisecond)
		}
	}

	resultado <- resultadoMatriz
}

func main() {
	rand.Seed(time.Now().UnixNano())

	//definir tamaño de matriz

	filasA := 3
	columnasA := 2
	filasB := 2
	columnasB := 3

	matrizA := generarMatriz(filasA, columnasA)
	matrizB := generarMatriz(filasB, columnasB)

	//Upwork
	fmt.Println("Matriz A: ")
	imprimirMatriz(matrizA)
	fmt.Println("Matriz B: ")
	imprimirMatriz(matrizB)

	resultado := make(chan [][]int)

	go multiplicarMatricesConcurrente(matrizA, matrizB, resultado)

	fmt.Println("Resultado del producto cruz: ")
	fmt.Println(<-resultado)
}
