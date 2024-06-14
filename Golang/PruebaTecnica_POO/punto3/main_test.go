package main

import (
	"testing"
)

func TestAgregarLibro(t *testing.T) {
	titulo := "El señor de los anillos"
	libros := make(map[string]*Libro)
	agregarLibro(libros, titulo, "J.R.R. Tolkien", "Fantasia")

	if len(libros) != 1 {
		t.Errorf("Se esperaba 1 libro, se obtuvo %d", len(libros))
	}

	if libros[titulo].Titulo != "El señor de los anillos" {
		t.Errorf("El título del libro no es correcto, se obtuvo %s", libros[titulo].Titulo)
	}
}

/*
func TestBuscarTitulo(t *testing.T) {
	titulo := "El señor de los anillos"
	libros := make(map[string]*Libro)
	agregarLibro(libros, titulo, "J.R.R. Tolkien", "Fantasia")
	buscarTitulo(libros, titulo)
	if len(libros) != 1 {
		t.Errorf("Se esperaba 1 libro, se obtuvo %d", len(libros))
	}
}
*/

func TestCambiarDisponibilidad(t *testing.T) {
	titulo := "El señor de los anillos"
	libros := make(map[string]*Libro)
	agregarLibro(libros, titulo, "J.R.R. Tolkien", "Fantasia")
	cambiarDisponibilidad(libros, titulo)
	if libros[titulo].Estado == true {
		t.Errorf("El libro se encuentra disponible, el estado no ha sido cambiado")
	}
}

func TestEliminarLibro(t *testing.T) {
	titulo := "El señor de los anillos"
	libros := make(map[string]*Libro)
	agregarLibro(libros, titulo, "J.R.R. Tolkien", "Fantasia")
	eliminarLibro(libros, titulo)
	if len(libros) != 0 {
		t.Errorf("Se esperaba 0 libros, se obtuvo %d", len(libros))
	}
}
