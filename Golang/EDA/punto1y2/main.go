/*
Ejercicio 1: Árbol de búsqueda de prefijos

Diseña e implementa una estructura de datos en Go para un árbol de búsqueda de prefijos (Trie) que almacene un conjunto de palabras. Proporciona funciones para insertar palabras en el árbol y buscar palabras por prefijo.
*/
package main

import "fmt"

type Node struct {
	children map[rune]*Node
	isEnd    bool
}

type Trie struct {
	root *Node
}

func NewTrie() *Trie {
	return &Trie{root: &Node{children: make(map[rune]*Node), isEnd: false}}
}

func (t *Trie) Insert(word string) {
	//iniciamos desde la raiz
	node := t.root

	//iteramos por cada caracter de la palabra
	for _, char := range word {
		//si no existe el nodo, lo creamos
		if _, ok := node.children[char]; !ok {
			node.children[char] = &Node{children: make(map[rune]*Node), isEnd: false}
		}
		//avanzamos al siguiente nodo
		node = node.children[char]
	}
	//marcamos el final de la palabra
	node.isEnd = true
}

func (t *Trie) Exist(word string) bool {
	//iniciamos desde la raiz
	node := t.root

	//iteramos por cada caracter de la palabra
	for _, char := range word {
		//si no existe el nodo, no hay coincidencias
		if _, ok := node.children[char]; !ok {
			return false
		}
		//avanzamos al siguiente nodo
		node = node.children[char]
	}

	//si llegamos al final del prefijo, retornamos true
	return node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	//iniciamos desde la raiz
	node := t.root

	//iteramos por cada caracter del prefijo
	for _, char := range prefix {
		//si no existe el nodo, no hay coincidencias
		if _, ok := node.children[char]; !ok {
			return false
		}
		//avanzamos al siguiente nodo
		node = node.children[char]
	}

	//si llegamos al final del prefijo, retornamos true
	return true
}

func (t *Trie) SearchWithPrefix(prefix string) []string {
	//iniciamos desde la raiz
	var result []string
	node := t.root

	//iteramos por cada caracter del prefijo
	for _, char := range prefix {
		//si no existe el nodo, no hay coincidencias
		if _, ok := node.children[char]; !ok {
			return result
		}
		//avanzamos al siguiente nodo
		node = node.children[char]
	}

	//si llegamos al final del prefijo, retornamos todas las palabras que comienzan con el prefijo
	t.search(node, prefix, &result)
	return result
}

func (t *Trie) search(node *Node, prefix string, result *[]string) {
	//si es el final de una palabra, la agregamos al resultado
	if node.isEnd {
		*result = append(*result, prefix)
	}
	//recorremos los hijos del nodo
	for char, child := range node.children {
		//llamamos recursivamente a search con el hijo y el prefijo actual
		t.search(child, prefix+string(char), result)
	}
}

// main
func main() {
	//creamos el trie
	trie := NewTrie()
	//datos a insertar
	words := []string{"hola", "holas", "mundo", "c++", "python", "java"}

	//interamos
	for _, word := range words {
		trie.Insert(word)
	}

	fmt.Println(trie.Exist("hola"))

	fmt.Println(trie.Exist("s"))

	fmt.Println("-----------------")
	fmt.Println(trie.StartsWith("hol"))

	fmt.Println("-----------------")
	fmt.Println(trie.SearchWithPrefix("hol"))
}
