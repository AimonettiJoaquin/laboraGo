package main

import (
	"fmt"
	"math"
)

type NodoAVL struct {
	valor     int
	izquierdo *NodoAVL
	derecho   *NodoAVL
	altura    int
}

type ArbolAVL struct {
	raiz *NodoAVL
}

// Función para obtener la altura de un nodo.
func altura(nodo *NodoAVL) int {
	if nodo == nil {
		return -1
	}
	return nodo.altura
}

// Función para calcular el balance de un nodo.
func balance(nodo *NodoAVL) int {
	if nodo == nil {
		return 0
	}
	return altura(nodo.izquierdo) - altura(nodo.derecho)
}

// Función para rotar a la izquierda en el nodo dado.
func rotarIzquierda(nodo *NodoAVL) *NodoAVL {
	derecha := nodo.derecho
	nodo.derecho = derecha.izquierdo
	derecha.izquierdo = nodo
	nodo.altura = int(math.Max(float64(altura(nodo.izquierdo)), float64(altura(nodo.derecho))) + 1)
	derecha.altura = int(math.Max(float64(altura(derecha.izquierdo)), float64(altura(derecha.derecho))) + 1)
	return derecha
}

// Función para rotar a la derecha en el nodo dado.
func rotarDerecha(nodo *NodoAVL) *NodoAVL {
	izquierda := nodo.izquierdo
	nodo.izquierdo = izquierda.derecho
	izquierda.derecho = nodo
	nodo.altura = int(math.Max(float64(altura(nodo.izquierdo)), float64(altura(nodo.derecho))) + 1)
	izquierda.altura = int(math.Max(float64(altura(izquierda.izquierdo)), float64(altura(izquierda.derecho))) + 1)
	return izquierda
}

// Función para insertar un valor en el árbol AVL.
func (t *ArbolAVL) Insertar(valor int) {
	t.raiz = insertar(t.raiz, valor)
}

func insertar(nodo *NodoAVL, valor int) *NodoAVL {
	if nodo == nil {
		return &NodoAVL{valor: valor, altura: 0}
	}

	if valor < nodo.valor {
		nodo.izquierdo = insertar(nodo.izquierdo, valor)
	} else if valor > nodo.valor {
		nodo.derecho = insertar(nodo.derecho, valor)
	} else {
		return nodo // El valor ya existe
	}

	// Actualizar altura del nodo actual
	nodo.altura = int(math.Max(float64(altura(nodo.izquierdo)), float64(altura(nodo.derecho))) + 1)

	// Calcular el balance del nodo
	balance := balance(nodo)

	// Casos de desequilibrio y rotaciones
	if balance > 1 { // Desbalance hacia la izquierda
		if valor < nodo.izquierdo.valor {
			// Rotación simple a la derecha
			return rotarDerecha(nodo)
		} else {
			// Rotación doble a la izquierda-derecha
			nodo.izquierdo = rotarIzquierda(nodo.izquierdo)
			return rotarDerecha(nodo)
		}
	}
	if balance < -1 { // Desbalance hacia la derecha
		if valor > nodo.derecho.valor {
			// Rotación simple a la izquierda
			return rotarIzquierda(nodo)
		} else {
			// Rotación doble a la derecha-izquierda
			nodo.derecho = rotarDerecha(nodo.derecho)
			return rotarIzquierda(nodo)
		}
	}

	return nodo
}

// Función para buscar un valor en el árbol AVL.
func (t *ArbolAVL) Buscar(valor int) bool {
	return buscar(t.raiz, valor)
}

func buscar(nodo *NodoAVL, valor int) bool {
	if nodo == nil {
		return false
	}
	if valor == nodo.valor {
		return true
	} else if valor < nodo.valor {
		return buscar(nodo.izquierdo, valor)
	} else {
		return buscar(nodo.derecho, valor)
	}
}

// Función para realizar un recorrido en orden del árbol AVL.
func (t *ArbolAVL) RecorrerInOrden() {
	recorrerInOrden(t.raiz)
	fmt.Println()
}

func recorrerInOrden(nodo *NodoAVL) {
	if nodo != nil {
		recorrerInOrden(nodo.izquierdo)
		fmt.Print(nodo.valor, " ")
		recorrerInOrden(nodo.derecho)
	}
}

func main() {
	arbol := &ArbolAVL{}

	arbol.Insertar(50)
	arbol.Insertar(30)
	arbol.Insertar(20)
	arbol.Insertar(40)
	arbol.Insertar(70)
	arbol.Insertar(60)
	arbol.Insertar(80)

	fmt.Println("Recorrido en orden del árbol AVL:")
	arbol.RecorrerInOrden()

	fmt.Println("Buscar 50:", arbol.Buscar(50))
	fmt.Println("Buscar 100:", arbol.Buscar(100))
}
