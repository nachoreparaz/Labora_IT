package main

import "fmt"

// // ======= DICCIONARIO DE CONTACTOS ===============

type Contacto struct {
	Nombre    string
	Telefono  string
	Email     string
	Direccion string
}

func agregarContacto(agenda map[string]Contacto, nombre, telefono, correo, direccion string) {
	agenda[nombre] = Contacto{
		Nombre:    nombre,
		Telefono:  telefono,
		Email:     correo,
		Direccion: direccion,
	}
	fmt.Printf("Contacto '%s' agregado a la agenda", nombre)
}

func buscarContacto(agenda map[string]Contacto, nombre string) (Contacto, bool) {
	contacto, encontrado := agenda[nombre]
	if !encontrado {

	}
	return contacto, encontrado
}

func main() {
	contactos := make(map[string]Contacto)
	agregarContacto(contactos, "Contacto 1", "123456789", "nacho@test.com", "Avenida Siempre Viva 123")
	agregarContacto(contactos, "Contacto 2", "123456789", "juan@test.com", "Avenida Numero 1 123")
	agregarContacto(contactos, "Contacto 3", "123456789", "andre@test.com", "Avenida Numero 2 123")

	buscarContacto(contactos, "Contacto 2")
}
