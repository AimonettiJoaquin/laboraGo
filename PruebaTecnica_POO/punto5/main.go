/*
5. SGPC - Sistema de gestión de proyectos creativos

Imagina que eres parte de un equipo de desarrollo de software en una empresa de tecnología creativa. Se te ha asignado la tarea de crear un sistema para gestionar
proyectos creativos.
El sistema debe permitir la creación de nuevos proyectos, asignación de miembros del equipo a proyectos, seguimiento del progreso del proyecto y generación de
informes de estado del proyecto. Utiliza tu imaginación para definir cómo serán los proyectos creativos y cómo se gestionarán en este sistema.

Requisitos:

Implementar un sistema para la creación de nuevos proyectos creativos.
Permitir la asignación de miembros del equipo a proyectos.
Seguimiento del progreso del proyecto, incluyendo hitos y fechas límite.
Generación de informes de estado del proyecto, que muestren el progreso, los miembros del equipo asignados y cualquier problema o desafío que surja durante el
desarrollo del proyecto.
Utiliza tu creatividad para definir otros requisitos y funcionalidades que consideres importantes para la gestión eficaz de proyectos creativos.

*/

package main

import "fmt"

func main() {
	m := &Miembros{}
	p := &Proyectos{}
	m.CrearMiembro("Juan", "Desarrollador")
	m.CrearMiembro("Maria", "Diseñadora")
	m.CrearMiembro("Pedro", "Tester")
	p.CrearProyecto("Proyecto 1", "Descripción del proyecto 1", "2023-05-31")
	p.CrearProyecto("Proyecto 2", "Descripción del proyecto 2", "2023-06-30")
	proyecto1 := p.ObtenerProyecto("Proyecto 1")
	p.AsignarMiembro(proyecto1, m.ObtenerMiembro("Juan"))
	p.AsignarMiembro(proyecto1, m.ObtenerMiembro("Maria"))
	p.ActualizarEstado(proyecto1, Terminado)
	proyecto2 := p.ObtenerProyecto("Proyecto 2")
	p.AsignarMiembro(proyecto2, m.ObtenerMiembro("Pedro"))
	p.AsignarMiembro(proyecto2, m.ObtenerMiembro("Juan"))
	p.AsignarMiembro(proyecto2, m.ObtenerMiembro("Maria"))
	p.ActualizarProblema(proyecto2, "Problema con el proyecto 2")
	fmt.Println("miembos proyecto 2:", proyecto2.miembros[0])
}

type ProyectoStatus string

const (
	Activo     ProyectoStatus = "Activo"
	Terminado  ProyectoStatus = "Terminado"
	Suspendido ProyectoStatus = "Suspendido"
)

type Proyecto struct {
	nombre      string
	descripcion string
	miembros    []Miembro
	estado      ProyectoStatus
	fechaLimite string
	problema    string
}

type Miembro struct {
	nombre string
	rol    string
}
type Miembros struct {
	miembros []Miembro
}

type Proyectos struct {
	proyectos []Proyecto
}

func (m *Miembros) CrearMiembro(nombre, rol string) {
	miembro := Miembro{nombre, rol}
	m.miembros = append(m.miembros, miembro)
}

func (p *Proyectos) CrearProyecto(nombre, description, fechaLimite string) {
	proyecto := Proyecto{nombre, description, nil, Activo, fechaLimite, ""}
	p.proyectos = append(p.proyectos, proyecto)
}

func (p *Proyectos) ObtenerProyecto(nombre string) *Proyecto {
	for _, proyecto := range p.proyectos {
		if proyecto.nombre == nombre {
			return &proyecto
		}
	}
	return nil
}

func (m *Miembros) ObtenerMiembro(nombre string) *Miembro {
	for _, miembro := range m.miembros {
		if miembro.nombre == nombre {

			return &miembro
		}
	}
	return nil
}

func (p *Proyectos) AsignarMiembro(proyecto *Proyecto, miembro *Miembro) {
	proyecto.miembros = append(proyecto.miembros, *miembro)
}

func (p *Proyectos) ActualizarEstado(proyecto *Proyecto, estado ProyectoStatus) {
	proyecto.estado = estado
}

func (p *Proyectos) ActualizarProblema(proyecto *Proyecto, problema string) {
	proyecto.problema = problema
}
