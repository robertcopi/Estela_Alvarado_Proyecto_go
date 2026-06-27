package main

import (
	"fmt"
)

// Definición básica de un struct
type Usuario struct {
	ID     int
	Nombre string
	Email  string
	Activo bool
}

// Struct con tags (útil para JSON, validaciones, etc.)
type Producto struct {
	ID     int     `json:"id"  db:"product_id"`
	Nombre string  `json:"nombre"  validate:"required"`
	Precio float64 `json:"precio"  validate:"min=0"`
}

// Struct anidado
type Direccion struct {
	Calle  string
	Ciudad string
	CP     string
}

type Empresa struct {
	Nombre    string
	Direccion Direccion
	Empleados []Usuario
}

func main() {
	fmt.Println("=== ESTRUCTURAS EN GO (STRUCTS) ===")

	// 1. Inicialización con valores cero
	var u1 Usuario
	fmt.Printf("1. Valores cero: %+v\n", u1)

	// 2. Inicialización con valores específicos
	u2 := Usuario{
		ID:     1,
		Nombre: "Juan",
		Email:  "juan@email.com",
		Activo: true,
	}
	fmt.Printf("2. Valores específicos: %+v\n", u2)

	// 3. Inicialización parcial
	u3 := Usuario{
		Nombre: "María",
		Email:  "maria@email.com",
	}
	fmt.Printf("3. Inicialización parcial: %+v\n", u3)

	// 4. Usando punteros
	u4 := &Usuario{
		ID:     2,
		Nombre: "Pedro",
	}
	fmt.Printf("4. Puntero a struct: %+v\n", u4)

	// 5. Structs Anónimos
	var person struct {
		name string
		age  int
		pet  string
	}
	person.name = "bob"
	person.age = 50
	person.pet = "dog"
	fmt.Printf("5. Struct anónimo variable: %+v\n", person)

	pet := struct {
		name string
		kind string
	}{
		name: "Fido",
		kind: "dog",
	}
	fmt.Printf("   Struct anónimo literal: %+v\n", pet)

	// 6. Structs anidados
	empresa := Empresa{
		Nombre: "Tech Corp",
		Direccion: Direccion{
			Calle:  "Av. Siempreviva 123",
			Ciudad: "Lima",
			CP:     "10001",
		},
		Empleados: []Usuario{u2, u3},
	}
	fmt.Printf("6. Struct anidado: Empresa=%s, Ciudad=%s, NumEmpleados=%d\n", empresa.Nombre, empresa.Direccion.Ciudad, len(empresa.Empleados))
}
