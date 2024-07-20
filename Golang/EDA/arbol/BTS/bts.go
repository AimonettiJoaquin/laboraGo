package main

import "fmt"

type Nodo struct {
	valor     int
	izquierdo *Nodo
	derecho   *Nodo
}

type ArbolBinarioBusqueda struct {
	raiz *Nodo
}

// Insertar inserta un nuevo valor en el árbol.
func (t *ArbolBinarioBusqueda) Insertar(valor int) {
	nuevoNodo := &Nodo{valor: valor}

	if t.raiz == nil {
		t.raiz = nuevoNodo
		return
	}

	actual := t.raiz
	for {
		if valor < actual.valor {
			if actual.izquierdo == nil {
				actual.izquierdo = nuevoNodo
				break
			}
			actual = actual.izquierdo
		} else {
			if actual.derecho == nil {
				actual.derecho = nuevoNodo
				break
			}
			actual = actual.derecho
		}
	}
}

// Buscar busca un valor en el árbol.
func (t *ArbolBinarioBusqueda) Buscar(valor int) bool {
	actual := t.raiz
	for actual != nil {
		if valor == actual.valor {
			return true
		} else if valor < actual.valor {
			actual = actual.izquierdo
		} else {
			actual = actual.derecho
		}
	}
	return false
}

// RecorrerInOrden realiza un recorrido en orden del árbol.
func (t *ArbolBinarioBusqueda) RecorrerInOrden() {
	t.recorrerInOrden(t.raiz)
}

func (t *ArbolBinarioBusqueda) recorrerInOrden(nodo *Nodo) {
	if nodo != nil {
		t.recorrerInOrden(nodo.izquierdo)
		fmt.Print(nodo.valor, " ")
		t.recorrerInOrden(nodo.derecho)
	}
}

func main() {
	arbol := &ArbolBinarioBusqueda{}

	arbol.Insertar(50)
	arbol.Insertar(30)
	arbol.Insertar(20)
	arbol.Insertar(40)
	arbol.Insertar(70)
	arbol.Insertar(60)
	arbol.Insertar(80)

	fmt.Println("Recorrido en orden del árbol:")
	arbol.RecorrerInOrden()

	fmt.Println("\nBuscar 50:", arbol.Buscar(50))
	fmt.Println("Buscar 100:", arbol.Buscar(100))
}
