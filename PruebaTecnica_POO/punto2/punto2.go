package main

import "fmt"

type Tarea struct {
	Nombre      string
	Descripcion string
	Responsable string
	Estado      string // "pendiente", "en progreso", "completada"
}

func main() {
	tarea1 := crearTarea("Tarea 1", "Esta es mi tarea1.", "Joaquin")
	fmt.Println(tarea1)
	tarea1.actualizarTarea("en progreso")
	printTarea(tarea1)
	tarea2 := crearTarea("Tarea 2", "Esta es mi tarea2.", "Pepe")
	fmt.Println(tarea2)
	tarea2.actualizarTarea("en progreso")
	printTarea(tarea2)
	tarea3 := crearTarea("Tarea 3", "Esta es mi tarea3.", "Jorge")
	fmt.Println(tarea3)
	tarea3.actualizarTarea("en progreso")
	printTarea(tarea3)

}

func crearTarea(nombre, descripcion, responsable string) *Tarea {
	nuevaTarea := &Tarea{
		Nombre:      nombre,
		Descripcion: descripcion,
		Responsable: responsable,
		Estado:      "pendiente",
	}
	return nuevaTarea
}
func printTarea(tarea *Tarea) {
	fmt.Println("----------------------------------------")
	fmt.Printf("Nombre de la tarea: %s\n", tarea.Nombre)
	fmt.Printf("Descripcion de la tarea: %s\n", tarea.Descripcion)
	fmt.Printf("Responsable: %s\n", tarea.Responsable)
	fmt.Printf("Estado: %s\n", tarea.Estado)
	fmt.Println("----------------------------------------")
}

func (tarea *Tarea) actualizarTarea(nuevoEstado string) {
	tarea.Estado = nuevoEstado
}

func mostrarTareasPendientes(tareas []*Tarea) {
	fmt.Println("Estas son mis tareas pendientes:")
	for _, tarea := range tareas {
		if tarea.Estado == "pendiente" {
			printTarea(tarea)
		}
	}
}
