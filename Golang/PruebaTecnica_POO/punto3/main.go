/* Una biblioteca necesita un sistema para gestionar su colección de libros. El programa debe permitir la adición de nuevos libros, la búsqueda de libros por título o autor, la actualización del estado de un libro (disponible o prestado) y la eliminación de libros.

Requisitos:

Implementar una estructura de datos para almacenar la información de cada libro (título, autor, género y estado).
Permitir la adición de nuevos libros a la colección.
Permitir la búsqueda de libros por título o autor.
Permitir la actualización del estado de un libro a "disponible" o "prestado".
Permitir la eliminación de libros de la colección.
*/

package main

import "fmt"

func main() {
	libros := make(map[string]*Libro)
	agregarLibro(libros, "El señor de los anillos", "J.R.R. Tolkien", "Fantasia")
	agregarLibro(libros, "El hobbit", "J.R.R. Tolkien", "Fantasia")
	agregarLibro(libros, "El juego de tronos", "George R.R. Martin", "Fantasia")

	buscarTitulo(libros, "El señor de los anillos")
	buscarAutor(libros, "J.R.R. Tolkiens")

	cambiarDisponibilidad(libros, "El hobbit")
	fmt.Println(libros["El hobbit"])

	eliminarLibro(libros, "El juego de tronos")
}

type Libro struct {
	Titulo string
	Autor  string
	Genero string
	Estado bool
}

func agregarLibro(libros map[string]*Libro, titulo, autor, genero string) {
	nuevoLibro := &Libro{
		Titulo: titulo,
		Autor:  autor,
		Genero: genero,
		Estado: true,
	}
	libros[titulo] = nuevoLibro
	fmt.Printf("Libros '%s' agregado a la colección \n", titulo)
}

func buscarTitulo(libros map[string]*Libro, titulo string) {
	if libro, ok := libros[titulo]; ok {
		fmt.Printf("Libro encontrado: %s \n", libro.Titulo)
	} else {
		fmt.Printf("Libro no encontrado: %s \n", titulo)
	}
}

func buscarAutor(libros map[string]*Libro, autor string) {
	for titulo, libro := range libros {
		if libro.Autor == autor {
			fmt.Printf("Libro encontrado: %s \n", titulo)
		}
	}
}
func cambiarDisponibilidad(libros map[string]*Libro, titulo string) {
	if libro, ok := libros[titulo]; ok {
		fmt.Println("*************************")
		fmt.Println("Cambiando disponibilidad")
		fmt.Println("*************************")
		fmt.Printf("Libro encontrado: %s \n", libro.Titulo)
		libro.Estado = !libro.Estado
		fmt.Println("*************************")
		fmt.Printf("Libro '%s' actualizado: %t \n", titulo, libro.Estado)
	} else {
		fmt.Println("*************************")
		fmt.Printf("Libro no encontrado: %s \n", titulo)
	}
}

func eliminarLibro(libros map[string]*Libro, titulo string) {
	if _, ok := libros[titulo]; ok {
		delete(libros, titulo)
		fmt.Printf("Libro '%s' eliminado \n", titulo)
	} else {
		fmt.Printf("Libro no encontrado: %s \n", titulo)
	}
}
