package main

import (
	"errors"
	"fmt"
	"strings"
)

// Interfaz que define las operaciones principales de la biblioteca
type OperacionesBiblioteca interface {
	VerCatalogo()
	BuscarLibro(criterio, valor string)
	PrestarLibro(titulo, email string) error
	DevolverLibro(titulo string) error
}

// Estructura para representar un libro
type Libro struct {
	titulo    string
	autor     string
	categoria string
	año       int
	prestado  bool
}

// Getters para los campos de la estructura de Libro
func (l *Libro) GetTitulo() string    { return l.titulo }
func (l *Libro) GetAutor() string     { return l.autor }
func (l *Libro) GetCategoria() string { return l.categoria }
func (l *Libro) GetAño() int          { return l.año }
func (l *Libro) GetPrestado() bool    { return l.prestado }

// Setter para cambiar el estado de préstamo
func (l *Libro) SetPrestado(estado bool) { l.prestado = estado }

// Estructura para un usuario
type Usuario struct {
	nombre   string
	email    string
	telefono string
}

// Getters para la estructura Usuario
func (u *Usuario) GetNombre() string   { return u.nombre }
func (u *Usuario) GetEmail() string    { return u.email }
func (u *Usuario) GetTelefono() string { return u.telefono }

// Estructura para la biblioteca
type Biblioteca struct {
	libros   []Libro
	usuarios []Usuario
}

// Implementación de la interfaz OperacionesBiblioteca

// Método para mostrar el catálogo de libros disponibles
func (b *Biblioteca) VerCatalogo() {
	fmt.Println("\nCatálogo de libros disponibles:")
	// Recorrer los libros y mostrar solo los que no están prestados
	for _, libro := range b.libros {
		if !libro.GetPrestado() {
			fmt.Printf("Título: %s, Autor: %s, Categoría: %s, Año: %d\n", libro.GetTitulo(), libro.GetAutor(), libro.GetCategoria(), libro.GetAño())
		}
	}
}

// Método para buscar un libro por varios criterios
func (b *Biblioteca) BuscarLibro(criterio, valor string) {
	fmt.Println("\nResultados de búsqueda:")
	criterio = strings.ToLower(criterio)
	valor = strings.ToLower(valor)

	var encontrados bool
	// Recorrer los libros y buscar según el criterio
	for _, libro := range b.libros {
		switch criterio {
		case "titulo":
			if strings.Contains(strings.ToLower(libro.GetTitulo()), valor) {
				encontrados = true
				fmt.Printf("Título: %s, Autor: %s, Categoría: %s, Año: %d\n", libro.GetTitulo(), libro.GetAutor(), libro.GetCategoria(), libro.GetAño())
			}
		default:
			fmt.Println("Criterio no válido.")
			return
		}
	}

	if !encontrados {
		fmt.Println("No se encontraron libros con ese criterio.")
	}
}

// Método para prestar un libro
func (b *Biblioteca) PrestarLibro(titulo, email string) error {
	for i, libro := range b.libros {
		if strings.ToLower(libro.GetTitulo()) == strings.ToLower(titulo) {
			if libro.GetPrestado() {
				return errors.New("el libro ya está prestado")
			}
			b.libros[i].SetPrestado(true)
			fmt.Printf("¡El libro \"%s\" ha sido prestado exitosamente!\n", titulo)
			return nil
		}
	}
	return errors.New("no se encontró el libro")
}

// Método para devolver un libro
func (b *Biblioteca) DevolverLibro(titulo string) error {
	for i, libro := range b.libros {
		if strings.ToLower(libro.GetTitulo()) == strings.ToLower(titulo) && libro.GetPrestado() {
			b.libros[i].SetPrestado(false)
			fmt.Printf("¡El libro \"%s\" ha sido devuelto exitosamente!\n", titulo)
			return nil
		}
	}
	return errors.New("el libro no está prestado o no existe en la biblioteca")
}

// Método para mostrar el menú
func mostrarMenu() {
	fmt.Println("\nMenú principal:")
	fmt.Println("1. Ver catálogo de libros disponibles")
	fmt.Println("2. Buscar libro por título")
	fmt.Println("3. Prestar libro")
	fmt.Println("4. Devolver libro")
	fmt.Println("5. Salir")
}

func main() {
	// Inicializar la biblioteca con algunos libros
	biblioteca := &Biblioteca{
		libros: []Libro{
			{titulo: "1984", autor: "George Orwell", categoria: "Ficción", año: 1949, prestado: false},
			{titulo: "Don Quijote", autor: "Miguel de Cervantes", categoria: "Clásico", año: 1605, prestado: false},
			{titulo: "El Principito", autor: "Antoine de Saint-Exupéry", categoria: "Ficción", año: 1943, prestado: false},
		},
	}

	// Ciclo principal para mostrar el menú y ejecutar las opciones
	for {
		mostrarMenu() // Mostrar el menú
		var opcion string
		fmt.Print("Seleccione una opción: ")
		fmt.Scanln(&opcion)

		switch opcion {
		case "1":
			biblioteca.VerCatalogo()
		case "2":
			var criterio, valor string
			fmt.Print("Buscar por (titulo): ")
			fmt.Scanln(&criterio)
			fmt.Print("Ingrese el valor de búsqueda: ")
			fmt.Scanln(&valor)
			biblioteca.BuscarLibro(criterio, valor)
		case "3":
			var titulo, email string
			fmt.Print("Ingrese el título del libro: ")
			fmt.Scanln(&titulo)
			fmt.Print("Ingrese su email: ")
			fmt.Scanln(&email)
			err := biblioteca.PrestarLibro(titulo, email)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
			}
		case "4":
			var titulo string
			fmt.Print("Ingrese el título del libro a devolver: ")
			fmt.Scanln(&titulo)
			err := biblioteca.DevolverLibro(titulo)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
			}
		case "5":
			fmt.Println("Gracias por usar el sistema de la biblioteca. ¡Hasta luego!")
			return // Salir del programa
		default:
			fmt.Println("Opción no válida, por favor intente de nuevo.")
		}
	}
}
