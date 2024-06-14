/*
4. Diccionario de contactos

Un usuario necesita un programa para gestionar sus contactos. El programa debe permitir la adición de nuevos contactos, la búsqueda de contactos por nombre o
número de teléfono, la actualización de la información de un contacto y la eliminación de contactos.

Requisitos:

Implementar una estructura de datos para almacenar la información de cada contacto (nombre, número de teléfono, correo electrónico y dirección).
Permitir la adición de nuevos contactos con su información completa.
Permitir la búsqueda de contactos por nombre o número de teléfono.
Permitir la actualización de la información de un contacto (teléfono, correo electrónico o dirección).
Permitir la eliminación de contactos de la lista.
*/
package main

import "fmt"

type Contacto struct {
	nombre    string
	numero    int
	correo    string
	direccion string
}

func main() {
	var contactos []Contacto
	var contacto1 Contacto
	contacto1 = crearContacto("Juan", 123456789, "pepito@gmail.com", "Calle 1")
	contacto2 := crearContacto("Maria", 987654321, "XXXXXXXXXXXXXXX", "Calle 2")
	contacto3 := crearContacto("Pedro", 111111111, "YYYYYYYYYYYYYYY", "Calle 3")
	contactos = append(contactos, contacto1)
	contactos = append(contactos, contacto2)
	contactos = append(contactos, contacto3)
	/*fmt.Println("-------------------------")
	fmt.Println(contactos)
	fmt.Println("-------------------------")
	fmt.Println(buscarContacto(contactos, "Juan", 0))
	fmt.Println("-------------------------")*/
	eliminarContacto(contactos, "Juan")
	fmt.Println("-------------------------")
	fmt.Println(contactos)
	//actualizarContacto(contactos, "Juan", 987654321, "XXXXXXXXXXXXXXX", "Calle 2")
	//fmt.Println(contactos)

}

func crearContacto(nombre string, numero int, correo string, direccion string) Contacto {
	return Contacto{nombre, numero, correo, direccion}
}

func buscarContacto(contactos []Contacto, nombre string, numero int) Contacto {
	for _, contacto := range contactos {
		if contacto.nombre == nombre || contacto.numero == numero {
			return contacto
		}
	}
	return Contacto{}
}

func eliminarContacto(contactos []Contacto, nombre string) []Contacto {
	for i, contacto := range contactos {
		if contacto.nombre == nombre {
			fmt.Println("Eliminando contacto:", contacto)
			contactos = append(contactos[:i], contactos[i+1:]...)
			return contactos
		}
	}
	return contactos
}

func actualizarContacto(contactos []Contacto, nombre string, numero int, correo string, direccion string) []Contacto {
	for i, contacto := range contactos {
		if contacto.nombre == nombre {
			contactos[i].numero = numero
			contactos[i].correo = correo
			contactos[i].direccion = direccion
			fmt.Println("Actualizando contacto:", contacto)
			return contactos
		}
	}
	return contactos
}
