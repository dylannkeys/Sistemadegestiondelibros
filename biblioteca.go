package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Estructura para representar un libro
type Libro struct {
	Titulo    string
	Autor     string
	Categoria string
	Año       int
	Prestado  bool
}

// Estructura para un usuario
type Usuario struct {
	Nombre   string
	Email    string
	Teléfono string
}

// Estructura para la biblioteca
type Biblioteca struct {
	Libros        []Libro
	Usuarios      []Usuario
	UsuarioActivo *Usuario
}

// Método para registrar un usuario
func (b *Biblioteca) RegistrarUsuario(nombre, email, telefono string) {
	nuevoUsuario := Usuario{Nombre: nombre, Email: email, Teléfono: telefono}
	b.Usuarios = append(b.Usuarios, nuevoUsuario)
	fmt.Println("¡Usuario registrado exitosamente!")
}

// Método para verificar si el usuario está registrado
func (b *Biblioteca) EstaRegistrado(email string) bool {
	for _, usuario := range b.Usuarios {
		if strings.ToLower(usuario.Email) == strings.ToLower(email) {
			return true
		}
	}
	return false
}

// Método para iniciar sesión
func (b *Biblioteca) IniciarSesion(email string) {
	for _, usuario := range b.Usuarios {
		if strings.ToLower(usuario.Email) == strings.ToLower(email) {
			b.UsuarioActivo = &usuario
			fmt.Printf("¡Bienvenido, %s!\n", usuario.Nombre)
			return
		}
	}
	fmt.Println("No se encontró un usuario con ese correo.")
}

// Método para mostrar todos los libros disponibles
func (b *Biblioteca) VerCatalogo() {
	fmt.Println("\nCatálogo de libros:")
	for _, libro := range b.Libros {
		if !libro.Prestado {
			fmt.Printf("Título: %s, Autor: %s, Categoría: %s, Año: %d\n", libro.Titulo, libro.Autor, libro.Categoria, libro.Año)
		}
	}
}

// Función auxiliar para mostrar el estado de prestado
func prestadoEstado(estado bool) string {
	if estado {
		return "Sí"
	}
	return "No"
}

// Método para prestar un libro
func (b *Biblioteca) PrestarLibro(titulo, email string) {
	if !b.EstaRegistrado(email) {
		fmt.Println("No está registrado. Por favor regístrese primero.")
		var nombre, telefono string
		fmt.Print("Ingrese su nombre: ")
		fmt.Scanln(&nombre)
		fmt.Print("Ingrese su teléfono: ")
		fmt.Scanln(&telefono)
		b.RegistrarUsuario(nombre, email, telefono)
	}

	// Ahora intentamos prestar el libro
	for i, libro := range b.Libros {
		if strings.ToLower(libro.Titulo) == strings.ToLower(titulo) && !libro.Prestado {
			b.Libros[i].Prestado = true
			fmt.Printf("¡El libro \"%s\" ha sido prestado exitosamente!\n", titulo)
			return
		}
	}
	fmt.Println("No se pudo prestar el libro, ya está prestado o no existe.")
}

// Método para devolver un libro
func (b *Biblioteca) DevolverLibro(titulo string) {
	for i, libro := range b.Libros {
		if strings.ToLower(libro.Titulo) == strings.ToLower(titulo) && libro.Prestado {
			b.Libros[i].Prestado = false
			fmt.Printf("¡El libro \"%s\" ha sido devuelto exitosamente!\n", titulo)
			return
		}
	}
	fmt.Println("Este libro no está prestado o no existe en la biblioteca.")
}

// Método para ofrecer la opción de prestar un libro
func (b *Biblioteca) OfrecerPrestarLibro(libro Libro) {
	var respuesta string
	if libro.Prestado {
		fmt.Printf("El libro \"%s\" ya está prestado.\n", libro.Titulo)
	} else {
		fmt.Println("Este libro está disponible. ¿Desea prestarlo? (s/n): ")
		fmt.Scanln(&respuesta)
		if strings.ToLower(respuesta) == "s" {
			var email string
			fmt.Print("Ingrese su email: ")
			fmt.Scanln(&email)
			b.PrestarLibro(libro.Titulo, email)
		} else {
			fmt.Println("No se prestará el libro.")
		}
	}
}

// Método para buscar un libro por varios criterios
func (b *Biblioteca) BuscarLibro(criterio, valor string) {
	fmt.Println("\nResultados de búsqueda:")
	criterio = strings.ToLower(criterio)
	valor = strings.ToLower(valor)

	var encontrados bool
	for _, libro := range b.Libros {
		// Usamos un switch para manejar los diferentes criterios de búsqueda
		switch criterio {
		case "titulo":
			if strings.Contains(strings.ToLower(libro.Titulo), valor) {
				encontrados = true
				fmt.Printf("Título: %s, Autor: %s, Categoría: %s, Año: %d, Prestado: %s\n", libro.Titulo, libro.Autor, libro.Categoria, libro.Año, prestadoEstado(libro.Prestado))
			}
		case "autor":
			if strings.Contains(strings.ToLower(libro.Autor), valor) {
				encontrados = true
				fmt.Printf("Título: %s, Autor: %s, Categoría: %s, Año: %d, Prestado: %s\n", libro.Titulo, libro.Autor, libro.Categoria, libro.Año, prestadoEstado(libro.Prestado))
			}
		case "categoria":
			if strings.Contains(strings.ToLower(libro.Categoria), valor) {
				encontrados = true
				fmt.Printf("Título: %s, Autor: %s, Categoría: %s, Año: %d, Prestado: %s\n", libro.Titulo, libro.Autor, libro.Categoria, libro.Año, prestadoEstado(libro.Prestado))
			}
		case "año":
			// Convertir el valor del año a entero
			valorInt, err := strconv.Atoi(valor)
			if err != nil {
				fmt.Println("Por favor ingrese un año válido.")
				return
			}
			if libro.Año == valorInt {
				encontrados = true
				fmt.Printf("Título: %s, Autor: %s, Categoría: %s, Año: %d, Prestado: %s\n", libro.Titulo, libro.Autor, libro.Categoria, libro.Año, prestadoEstado(libro.Prestado))
			}
		default:
			fmt.Println("Criterio no válido.")
			return
		}
	}

	// Si no se encontraron libros, indicamos al usuario
	if !encontrados {
		fmt.Println("No se encontraron libros con ese criterio.")
	}
}

func mostrarMenu() {
	fmt.Println("\nMenú principal:")
	fmt.Println("1. Ver catálogo de libros")
	fmt.Println("2. Buscar libro")
	fmt.Println("3. Prestar libro")
	fmt.Println("4. Devolver libro")
	fmt.Println("5. Registrarse")
	fmt.Println("6. Iniciar sesión")
	fmt.Println("7. Salir")
}

func main() {
	var biblioteca Biblioteca

	// Inicializamos algunos libros en el catálogo
	biblioteca.Libros = append(biblioteca.Libros,
		Libro{Titulo: "El Gran Gatsby", Autor: "F. Scott Fitzgerald", Categoria: "Ficción", Año: 1925, Prestado: false},
		Libro{Titulo: "1984", Autor: "George Orwell", Categoria: "Ficción", Año: 1949, Prestado: false},
		Libro{Titulo: "Cien Años de Soledad", Autor: "Gabriel García Márquez", Categoria: "Realismo mágico", Año: 1967, Prestado: false},
		Libro{Titulo: "Don Quijote de la Mancha", Autor: "Miguel de Cervantes", Categoria: "Clásico", Año: 1605, Prestado: false},
		Libro{Titulo: "Matar a un ruiseñor", Autor: "Harper Lee", Categoria: "Ficción", Año: 1960, Prestado: false},
		Libro{Titulo: "El Hobbit", Autor: "J.R.R. Tolkien", Categoria: "Fantasía", Año: 1937, Prestado: false},
		Libro{Titulo: "Harry Potter y la piedra filosofal", Autor: "J.K. Rowling", Categoria: "Fantasía", Año: 1997, Prestado: false},
		Libro{Titulo: "El código Da Vinci", Autor: "Dan Brown", Categoria: "Misterio", Año: 2003, Prestado: false},
		Libro{Titulo: "Orgullo y prejuicio", Autor: "Jane Austen", Categoria: "Romántico", Año: 1813, Prestado: false},
		Libro{Titulo: "Los pilares de la Tierra", Autor: "Ken Follett", Categoria: "Histórico", Año: 1989, Prestado: false},
		Libro{Titulo: "La sombra del viento", Autor: "Carlos Ruiz Zafón", Categoria: "Misterio", Año: 2001, Prestado: false},
		Libro{Titulo: "Fahrenheit 451", Autor: "Ray Bradbury", Categoria: "Ciencia ficción", Año: 1953, Prestado: false},
		Libro{Titulo: "En el camino", Autor: "Jack Kerouac", Categoria: "Aventura", Año: 1957, Prestado: false},
		Libro{Titulo: "Frankenstein", Autor: "Mary Shelley", Categoria: "Terror", Año: 1818, Prestado: false},
		Libro{Titulo: "Drácula", Autor: "Bram Stoker", Categoria: "Terror", Año: 1897, Prestado: false},
	)

	// Mostrar el menú y permitir seleccionar opciones
	for {
		mostrarMenu()
		var opcion int
		fmt.Print("Seleccione una opción: ")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			biblioteca.VerCatalogo()
		case 2:
			var criterio, valor string
			fmt.Print("Buscar por (titulo, autor, categoria, año): ")
			fmt.Scanln(&criterio)
			fmt.Print("Ingrese el valor de búsqueda: ")
			fmt.Scanln(&valor)
			biblioteca.BuscarLibro(criterio, valor)
		case 3:
			var titulo, email string
			fmt.Print("Ingrese el título del libro: ")
			fmt.Scanln(&titulo)
			fmt.Print("Ingrese su email: ")
			fmt.Scanln(&email)
			biblioteca.PrestarLibro(titulo, email)
		case 4:
			var titulo string
			fmt.Print("Ingrese el título del libro a devolver: ")
			fmt.Scanln(&titulo)
			biblioteca.DevolverLibro(titulo)
		case 5:
			var nombre, email, telefono string
			fmt.Print("Ingrese su nombre: ")
			fmt.Scanln(&nombre)
			fmt.Print("Ingrese su email: ")
			fmt.Scanln(&email)
			fmt.Print("Ingrese su teléfono: ")
			fmt.Scanln(&telefono)
			biblioteca.RegistrarUsuario(nombre, email, telefono)
		case 6:
			var email string
			fmt.Print("Ingrese su email: ")
			fmt.Scanln(&email)
			biblioteca.IniciarSesion(email)
		case 7:
			fmt.Println("Gracias por usar el sistema de la biblioteca. ¡Hasta luego!")
			return
		default:
			fmt.Println("Opción no válida, por favor intente de nuevo.")
		}
	}
}
