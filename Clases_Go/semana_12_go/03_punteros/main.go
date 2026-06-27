package main

import (
	"fmt"
)

// Person struct para demostración de pasos por referencia
type Person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("=== PUNTEROS Y GESTIÓN DE MEMORIA ===")

	// 1. Declaración básica y Operador de Dirección &
	var x int = 42
	var p *int = &x // p guarda la dirección de memoria de x

	fmt.Printf("1. Valor de x: %d, Dirección de x: %v, Valor de p: %v\n", x, &x, p)

	// 2. Desreferenciación con *
	fmt.Printf("2. Desreferencia de p (*p): %d\n", *p)

	// 3. Modificando el valor a través del puntero
	*p = 100
	fmt.Printf("3. Tras '*p = 100', el valor original de x es: %d\n", x)

	// 4. Uso de la función predefinida new()
	ptrNew := new(int) // Asigna memoria para un int y devuelve un puntero (con zero value 0)
	fmt.Printf("4. Puntero creado con new() valor: %v, desreferenciado: %d\n", ptrNew, *ptrNew)

	*ptrNew = 15
	fmt.Printf("   Valor modificado en el heap: %d\n", *ptrNew)
	ptrNew = nil // Liberar la referencia de memoria

	// 5. Punteros con Funciones (Pasar por referencia)
	fmt.Println("\n--- Funciones y paso por referencia ---")

	val := 10
	fmt.Println("Antes de modifyValue:", val)
	modifyValue(&val)
	fmt.Println("Después de modifyValue:", val)

	// 6. Punteros con Structs Complejos
	fmt.Println("\n--- Punteros en Estructuras de Datos ---")

	per := Person{Name: "John", Age: 30}
	fmt.Println("Before modification:", per)
	modifyPerson(&per) // Enviamos la dirección de memoria del struct, no una copia
	fmt.Println("After modification:", per)
}

// Función que modifica el valor original porque recibe un puntero
func modifyValue(ptr *int) {
	*ptr = 230
}

// Función que modifica los campos de un Struct directamente en su bloque de memoria
func modifyPerson(ptr *Person) {
	ptr.Age = 31 // Go permite acceder directamente a los campos a través del puntero sin necesidad de (*ptr).Age explícito
	ptr.Name = "John Doe"
}
