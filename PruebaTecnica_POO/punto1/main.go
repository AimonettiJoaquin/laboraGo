/*
1. Gestión de inventario

Una pequeña tienda necesita un programa para gestionar su inventario de productos. El programa debe permitir la adición de nuevos productos, la actualización de la
cantidad disponible de un producto, la eliminación de productos y la visualización del inventario completo.

Requisitos:

Implementar una estructura de datos para almacenar la información de cada producto (nombre, precio y cantidad disponible).
Permitir la adición de nuevos productos con sus respectivas cantidades.
Permitir la actualización de la cantidad disponible de un producto existente.
Permitir la eliminación de productos del inventario.
Mostrar el inventario completo, incluyendo el nombre, precio y cantidad disponible de cada producto.
*/
package main

import (
	"fmt"
)

type Producto struct {
	Nombre             string
	Precio             float32
	CantidadDisponible int
}

func main() {
	producto1 := crearProducto("Leche", 1.5, 10)
	fmt.Println(producto1)
	fmt.Println("---------------")
	fmt.Println("Actualizando producto")
	fmt.Println("---------------")
	producto1.actualizarCantidadDisponible(5)
	fmt.Println(producto1)
	fmt.Println("---------------")
	fmt.Println("Eliminando producto")
	fmt.Println("---------------")
	eliminarProducto(producto1)
	fmt.Println(producto1)
	fmt.Println("---------------")
}

func crearProducto(nombre string, precio float32, cantidadDisponible int) *Producto {
	return &Producto{
		Nombre:             nombre,
		Precio:             precio,
		CantidadDisponible: cantidadDisponible,
	}
}

func (p *Producto) actualizarCantidadDisponible(cantidad int) {
	p.CantidadDisponible = cantidad
}

func eliminarProducto(p *Producto) {
	p.Nombre = ""
	p.Precio = 0
	p.CantidadDisponible = 0
}
