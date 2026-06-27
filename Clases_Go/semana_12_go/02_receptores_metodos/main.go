package main

import (
	"errors"
	"fmt"
	"time"
)

// === MODELOS (STRUCTS) ===

type Libro struct {
	ID               int
	Titulo           string
	Autor            string
	Disponible       bool
	FechaPublicacion time.Time
}

type Usuario struct {
	ID              int
	Nombre          string
	LibrosPrestados []int // IDs de los libros
	Activo          bool
}

type Biblioteca struct {
	Nombre   string
	Libros   map[int]Libro
	Usuarios map[int]Usuario
}

// === RECEPTORES DE VALOR (Solo lectura, no modifican el estado original) ===

// Método con receptor de valor de Libro
func (l Libro) ImprimirDetalle() {
	estado := "No disponible"
	if l.Disponible {
		estado = "Disponible"
	}
	fmt.Printf("📖 [%d] %s por %s (%s)\n", l.ID, l.Titulo, l.Autor, estado)
}

// Método con receptor de valor de Usuario
func (u Usuario) PuedePrestarLibro() bool {
	return u.Activo && len(u.LibrosPrestados) < 3 // Máximo 3 libros
}

// === RECEPTORES DE PUNTERO (Modifican el estado original) ===

// Método con receptor de puntero de Biblioteca
func (b *Biblioteca) AgregarLibro(libro Libro) {
	libro.Disponible = true // Por defecto lo hacemos disponible
	b.Libros[libro.ID] = libro
	fmt.Printf("📚 Libro agregado: %s\n", libro.Titulo)
}

// Método con receptor de puntero de Biblioteca
func (b *Biblioteca) RegistrarUsuario(usuario Usuario) {
	usuario.Activo = true
	usuario.LibrosPrestados = []int{}
	b.Usuarios[usuario.ID] = usuario
	fmt.Printf("👤 Usuario registrado: %s\n", usuario.Nombre)
}

// Método con receptor de puntero y manejo de errores (PrestarLibro)
func (b *Biblioteca) PrestarLibro(usuarioID int, libroID int) error {
	// Verificar que usuario existe usando receptor de valor
	usuario, okUsuario := b.Usuarios[usuarioID]
	if !okUsuario {
		return errors.New("usuario no encontrado")
	}

	// Verificar si el usuario puede pedir préstamos
	if !usuario.PuedePrestarLibro() {
		return errors.New("el usuario no está activo o alcanzó límite de préstamos")
	}

	// Verificar existencia de libro
	libro, okLibro := b.Libros[libroID]
	if !okLibro {
		return errors.New("libro no encontrado")
	}

	// Verificar disponibilidad
	if !libro.Disponible {
		return errors.New("el libro no está disponible")
	}

	// ===== MUTACIÓN DE ESTADO =====
	// Modificar copia de struct libro y actualizar mapa
	libro.Disponible = false
	b.Libros[libroID] = libro

	// Modificar copia de struct usuario y actualizar mapa
	usuario.LibrosPrestados = append(usuario.LibrosPrestados, libroID)
	b.Usuarios[usuarioID] = usuario

	fmt.Printf("✅ Préstamo exitoso: '%s' para %s\n", libro.Titulo, usuario.Nombre)
	return nil
}

// Método con receptor de puntero y manejo de errores (DevolverLibro)
func (b *Biblioteca) DevolverLibro(usuarioID int, libroID int) error {
	usuario, okUsuario := b.Usuarios[usuarioID]
	if !okUsuario {
		return errors.New("usuario no encontrado")
	}

	// Buscar si el usuario tiene ese libro prestado
	indicePrestamo := -1
	for i, id := range usuario.LibrosPrestados {
		if id == libroID {
			indicePrestamo = i
			break
		}
	}

	if indicePrestamo == -1 {
		return errors.New("el usuario no tiene prestado este libro")
	}

	// Mover libro a disponible
	libro := b.Libros[libroID]
	libro.Disponible = true
	b.Libros[libroID] = libro

	// Quitar libro del arreglo de prestados (slice removal pattern)
	usuario.LibrosPrestados = append(usuario.LibrosPrestados[:indicePrestamo], usuario.LibrosPrestados[indicePrestamo+1:]...)
	b.Usuarios[usuarioID] = usuario

	fmt.Printf("🔙 Devolución exitosa: '%s' por %s\n", libro.Titulo, usuario.Nombre)
	return nil
}

// Método constructor/inicializador (Función normal, no método)
func ConstructorBiblioteca(nombre string) *Biblioteca {
	return &Biblioteca{
		Nombre:   nombre,
		Libros:   make(map[int]Libro),
		Usuarios: make(map[int]Usuario),
	}
}

func main() {
	fmt.Println("=== SISTEMA DE BIBLIOTECA (RECEPTORES Y MÉTODOS) ===")

	miBiblioteca := ConstructorBiblioteca("Biblioteca Nacional")

	miBiblioteca.AgregarLibro(Libro{ID: 101, Titulo: "Don Quijote", Autor: "Miguel de Cervantes", FechaPublicacion: time.Date(1605, 1, 1, 0, 0, 0, 0, time.UTC)})
	miBiblioteca.AgregarLibro(Libro{ID: 102, Titulo: "Cien años de soledad", Autor: "Gabriel García Márquez", FechaPublicacion: time.Date(1967, 5, 30, 0, 0, 0, 0, time.UTC)})

	miBiblioteca.RegistrarUsuario(Usuario{ID: 201, Nombre: "Roberto"})

	fmt.Println("\n--- Catálogo Inicial ---")
	for _, l := range miBiblioteca.Libros {
		l.ImprimirDetalle() // Usando receptor de valor
	}

	fmt.Println("\n--- Operaciones ---")
	err := miBiblioteca.PrestarLibro(201, 101) // Operación por puntero (muta estado)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// Intento fallido
	err = miBiblioteca.PrestarLibro(201, 101)
	if err != nil {
		fmt.Printf("Error esperado: %v\n", err)
	}

	fmt.Println("\n--- Catálogo tras préstamo ---")
	for _, l := range miBiblioteca.Libros {
		l.ImprimirDetalle()
	}

	fmt.Println("\n--- Devolución ---")
	err = miBiblioteca.DevolverLibro(201, 101)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println("\n--- Estado Final del usuario ---")
	fmt.Printf("Libros prestados para Roberto: %v\n", miBiblioteca.Usuarios[201].LibrosPrestados)
}
