package main

import "fmt"

// ========================= EJERCICIO 3 ===============
/*
Una biblioteca necesita un sistema para gestionar su colección de libros.
El programa debe permitir la adición de nuevos libros, la búsqueda de libros por título o autor,
la actualización del estado de un libro (disponible o prestado) y la eliminación de libros.

Requisitos:

Implementar una estructura de datos para almacenar la información de cada libro (título, autor, género y estado).
Permitir la adición de nuevos libros a la colección.
Permitir la búsqueda de libros por título o autor.
Permitir la actualización del estado de un libro a "disponible" o "prestado".
Permitir la eliminación de libros de la colección.

*/

type Biblioteca struct {
	Titulo string
	Autor  string
	Genero string
	Estado string
}

func agregarLibro(biblioteca map[string]Biblioteca, titulo, autor, genero, estado string) {
	biblioteca[titulo] = Biblioteca{
		Titulo: titulo,
		Autor:  autor,
		Genero: genero,
		Estado: estado,
	}
	fmt.Printf("\n Libro '%s' agregado a la biblioteca", titulo)
}

func buscarPorTitulo(biblioteca map[string]Biblioteca, titulo string) (Biblioteca, bool) {
	libro, ok := biblioteca[titulo]
	if !ok {
		fmt.Println("\nLibro no encontrado")
	}
	return libro, ok
}

func buscarPorAutor(biblioteca map[string]Biblioteca, autor string) []Biblioteca {
	var libros []Biblioteca
	for _, libro := range biblioteca {
		if libro.Autor == autor {
			libros = append(libros, libro)
		}
	}
	return libros
}

func actualizarEstado(biblioteca map[string]Biblioteca, titulo string, estado string) {
	if libro, ok := biblioteca[titulo]; ok {
		libro.Estado = estado
		biblioteca[titulo] = libro
		fmt.Printf("\nLibro '%s' actualizado", titulo)
	}
}

func eliminarLibro(biblioteca map[string]Biblioteca, titulo string) {
	if _, ok := biblioteca[titulo]; ok {
		delete(biblioteca, titulo)
	} else {
		fmt.Println("\nLibro no encontrado")
	}
}

func mostrarBiblioteca(biblioteca map[string]Biblioteca) {
	for _, libro := range biblioteca {
		fmt.Println("\nLibro: ", libro.Titulo, libro.Autor, libro.Genero, libro.Estado)
	}
}

func main() {
	biblioteca := make(map[string]Biblioteca)

	fmt.Print("\nBiblioteca Al Crearse\n")
	agregarLibro(biblioteca, "Libro 1", "Autor 1", "Genero 1", "Disponible")
	agregarLibro(biblioteca, "Libro 2", "Autor 2", "Genero 2", "Disponible")
	agregarLibro(biblioteca, "Libro 3", "Autor 3", "Genero 3", "Disponible")
	mostrarBiblioteca(biblioteca)

	fmt.Print("\nBiblioteca Al Actualizarse:\n")
	actualizarEstado(biblioteca, "Libro 3", "Prestado")
	mostrarBiblioteca(biblioteca)

	fmt.Print("\nBiblioteca Al Eliminarse\n")
	eliminarLibro(biblioteca, "Libro 2")
	mostrarBiblioteca(biblioteca)
}
