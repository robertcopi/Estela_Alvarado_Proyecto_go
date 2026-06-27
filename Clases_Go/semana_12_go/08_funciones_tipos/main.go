package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

// ========= FUNCIONES COMO TIPOS (FIRST-CLASS CITIZENS) =========

type Libro struct {
	ID               int
	Titulo           string
	Autor            string
	Disponible       bool
	Precio           float64
	FechaPublicacion time.Time
}

// 1. Definir tipos de función
type FiltroLibro func(Libro) bool
type ComparadorLibro func(Libro, Libro) bool
type TransformadorLibro func(Libro) Libro

// 2. Función de Orden Superior que acepta otra función (filtro) como parámetro
func FiltrarLibros(libros []Libro, filtro FiltroLibro) []Libro {
	var resultado []Libro
	for _, libro := range libros {
		if filtro(libro) {
			resultado = append(resultado, libro)
		}
	}
	return resultado
}

// 3. Función que retorna otra función (Fábrica de funciones)
func CrearFiltroAutor(autorBuscado string) FiltroLibro {
	return func(libro Libro) bool {
		return strings.Contains(strings.ToLower(libro.Autor), strings.ToLower(autorBuscado))
	}
}

// 4. Función de orden superior para transformar libros (Map)
func MapearLibros(libros []Libro, transformador TransformadorLibro) []Libro {
	resultado := make([]Libro, len(libros))
	for i, libro := range libros {
		resultado[i] = transformador(libro)
	}
	return resultado
}

// 5. Función para ordenar con comparador personalizado
func OrdenarLibros(libros []Libro, comparador ComparadorLibro) {
	sort.Slice(libros, func(i, j int) bool {
		return comparador(libros[i], libros[j])
	})
}

// 6. Función que combina múltiples filtros (composición)
func CombinarFiltros(filtros ...FiltroLibro) FiltroLibro {
	return func(libro Libro) bool {
		for _, filtro := range filtros {
			if !filtro(libro) {
				return false
			}
		}
		return true
	}
}

// Simulación de Base de Datos
func obtenerLibros() []Libro {
	return []Libro{
		{1, "El Hobbit", "J.R.R. Tolkien", true, 20.0, time.Date(1937, 9, 21, 0, 0, 0, 0, time.UTC)},
		{2, "El Señor de los Anillos", "J.R.R. Tolkien", false, 35.0, time.Date(1954, 7, 29, 0, 0, 0, 0, time.UTC)},
		{3, "Dune", "Frank Herbert", true, 22.0, time.Date(1965, 8, 1, 0, 0, 0, 0, time.UTC)},
		{4, "Cien años de soledad", "Gabriel García Márquez", true, 18.0, time.Date(1967, 5, 30, 0, 0, 0, 0, time.UTC)},
	}
}

func main() {
	fmt.Println("=== 08_FUNCIONES_COMO_TIPOS ===")
	libros := obtenerLibros()

	// Filtros predefinidos como funciones literales
	filtroDisponibles := func(l Libro) bool { return l.Disponible }
	filtroBaratos := func(l Libro) bool { return l.Precio < 25.0 }

	// Fábrica dinámica de filtros
	filtroTolkien := CrearFiltroAutor("Tolkien")

	// Uso individual de Filtrar
	disponibles := FiltrarLibros(libros, filtroDisponibles)
	fmt.Printf("\nLibrería: \n- %d disponibles\n", len(disponibles))

	librosTolkien := FiltrarLibros(libros, filtroTolkien)
	fmt.Printf("- %d libros de Tolkien\n", len(librosTolkien))

	// Composición de filtros
	filtroComplejo := CombinarFiltros(filtroDisponibles, filtroBaratos)
	librosEspeciales := FiltrarLibros(libros, filtroComplejo)
	fmt.Printf("- %d libros disponibles y baratos (< 25$)\n", len(librosEspeciales))

	// Función Transformadora (Aplicamos descuento)
	aplicarDescuento := func(l Libro) Libro {
		l.Precio = l.Precio * 0.9 // 10% de descuento
		return l
	}
	librosConDescuento := MapearLibros(libros, aplicarDescuento)
	fmt.Printf("- Descuento aplicado a %d libros\n", len(librosConDescuento))

	// Función Comparadora
	fmt.Println("\nOrdenando catálogo...")
	OrdenarLibros(libros, func(a, b Libro) bool {
		return a.FechaPublicacion.Before(b.FechaPublicacion) // Más antiguos primero
	})

	for _, l := range libros {
		fmt.Printf("   %s (%d)\n", l.Titulo, l.FechaPublicacion.Year())
	}
}
