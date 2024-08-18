package main

import "testing"

func TestCrearMiembro(t *testing.T) {
	m := &Miembros{}
	m.CrearMiembro("Luis", "Desarrollador")

	if len(m.miembros) != 1 {
		t.Errorf("No se encontro el miembro")
	}

	if m.miembros[0].nombre != "Luis" {
		t.Errorf("El nombre del miembro no es correcto, se esperaba Luis")
	}
}

func TestAgregarProyecto(t *testing.T) {
	p := &Proyectos{}
	p.CrearProyecto("Proyecto 1", "Descripción del proyecto 1", "2023-05-31")

	if len(p.proyectos) != 1 {
		t.Errorf("No se encontro el proyecto")
	}

	if p.proyectos[0].nombre != "Proyecto 1" {
		t.Errorf("El nombre del proyecto no es correcto, se esperaba Proyecto 1")
	}
}

func TestObtenerProyecto(t *testing.T) {
	p := &Proyectos{}
	p.CrearProyecto("Proyecto 1", "Descripción del proyecto 1", "2023-05-31")
	proyecto1 := p.ObtenerProyecto("Proyecto 1")
	if proyecto1 == nil {
		t.Errorf("No se encontro el proyecto")
	}
}

func TestObtenerMiembro(t *testing.T) {
	m := &Miembros{}
	m.CrearMiembro("Luis", "Desarrollador")
	miembro1 := m.ObtenerMiembro("Luis")
	if miembro1 == nil {
		t.Errorf("No se encontro el miembro")
	}
}

func TestAsignarMiembro(t *testing.T) {
	p := &Proyectos{}
	m := &Miembros{}
	p.CrearProyecto("Proyecto 1", "Descripción del proyecto 1", "2023-05-31")
	m.CrearMiembro("Luis", "Desarrollador")
	p.AsignarMiembro(p.ObtenerProyecto("Proyecto 1"), m.ObtenerMiembro("Juan"))
}
