package main

import "fmt"

type Tarea struct {
	Nombre      string
	Descripcion string
	Responsable string
	Estado      string
}

func crearTarea(nombre string, descripcion string, responsable string) *Tarea {
	nuevaTarea := &Tarea{
		Nombre:      nombre,
		Descripcion: descripcion,
		Responsable: responsable,
		Estado:      "pendiente",
	}
	return nuevaTarea
}

func cambiarEstado(tarea *Tarea, nuevoEstado string) {
	tarea.Estado = nuevoEstado
}

func mostrarTareasPendientes(tareas []*Tarea) {
	for _, tarea := range tareas {
		if tarea.Estado == "pendiente" {
			fmt.Println(tarea.Nombre)
		}
	}
}

func main() {
	tarea1 := crearTarea("Tarea 1", "Primera Tarea", "Nacho")
	tarea2 := crearTarea("Tarea 2", "Segunda Tarea", "Andre")

	tareas := []*Tarea{tarea1, tarea2}
	mostrarTareasPendientes(tareas)

	cambiarEstado(tarea2, "Completada")
	mostrarTareasPendientes(tareas)
}
