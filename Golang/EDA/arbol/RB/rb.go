package main

import "fmt"

type Color bool

const (
	Red   Color = true
	Black Color = false
)

type NodoRN struct {
	valor     int
	izquierdo *NodoRN
	derecho   *NodoRN
	color     Color
}

type ArbolRN struct {
	raiz *NodoRN
}

func (t *ArbolRN) Insert(valor int) {
	t.raiz = insertar(t.raiz, valor)
	t.raiz.color = Black // La raíz siempre debe ser negra
}

func insertar(nodo *NodoRN, valor int) *NodoRN {
	if nodo == nil {
		return &NodoRN{valor: valor, color: Red}
	}

	if valor < nodo.valor {
		nodo.izquierdo = insertar(nodo.izquierdo, valor)
	} else if valor > nodo.valor {
		nodo.derecho = insertar(nodo.derecho, valor)
	}

	// Restaurar las propiedades del árbol rojo-negro después de la inserción
	if esRojo(nodo.derecho) && !esRojo(nodo.izquierdo) {
		nodo = rotarIzquierda(nodo)
	}
	if esRojo(nodo.izquierdo) && esRojo(nodo.izquierdo.izquierdo) {
		nodo = rotarDerecha(nodo)
	}
	if esRojo(nodo.izquierdo) && esRojo(nodo.derecho) {
		invertirColores(nodo)
	}

	return nodo
}

func esRojo(nodo *NodoRN) bool {
	if nodo == nil {
		return false
	}
	return nodo.color == Red
}

func rotarIzquierda(nodo *NodoRN) *NodoRN {
	x := nodo.derecho
	nodo.derecho = x.izquierdo
	x.izquierdo = nodo
	x.color = nodo.color
	nodo.color = Red
	return x
}

func rotarDerecha(nodo *NodoRN) *NodoRN {
	x := nodo.izquierdo
	nodo.izquierdo = x.derecho
	x.derecho = nodo
	x.color = nodo.color
	nodo.color = Red
	return x
}

func invertirColores(nodo *NodoRN) {
	nodo.color = !nodo.color
	nodo.izquierdo.color = !nodo.izquierdo.color
	nodo.derecho.color = !nodo.derecho.color
}

func (t *ArbolRN) Search(valor int) bool {
	return search(t.raiz, valor)
}

func search(nodo *NodoRN, valor int) bool {
	if nodo == nil {
		return false
	}
	if valor == nodo.valor {
		return true
	} else if valor < nodo.valor {
		return search(nodo.izquierdo, valor)
	} else {
		return search(nodo.derecho, valor)
	}
}

func (t *ArbolRN) InOrderTraversal() {
	inOrderTraversal(t.raiz)
	fmt.Println()
}

func inOrderTraversal(nodo *NodoRN) {
	if nodo != nil {
		inOrderTraversal(nodo.izquierdo)
		fmt.Print(nodo.valor, " ")
		inOrderTraversal(nodo.derecho)
	}
}

func main() {
	arbol := &ArbolRN{}

	arbol.Insert(50)
	arbol.Insert(30)
	arbol.Insert(20)
	arbol.Insert(40)
	arbol.Insert(70)
	arbol.Insert(60)
	arbol.Insert(80)

	fmt.Println("Recorrido en orden del árbol rojo-negro:")
	arbol.InOrderTraversal()

	fmt.Println("Buscar 50:", arbol.Search(50))
	fmt.Println("Buscar 100:", arbol.Search(100))
}
