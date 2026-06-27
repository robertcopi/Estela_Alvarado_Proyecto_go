package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"unicode/utf8"
	"unsafe"
)

func main() {
	tiposBooleanos()
	tiposEnteros()
	tiposPuntoFlotante()
	tiposComplejos()
	stringsAndRunes()
}

func tiposBooleanos() {
	var activo bool = true
	var disponible bool = false
	var validado = true
	var configurado bool

	fmt.Printf("activo: %t, disponible: %t, validado: %t, configurado: %t\n",
		activo, disponible, validado, configurado)

	resultado := activo && disponible
	negacion := !activo

	fmt.Printf("resultado: %t, negación: %t\n", resultado, negacion)
}

func tiposEnteros() {
	var a int8 = 127
	var b int16 = 32767
	var c int32 = 2147483647
	var d int64 = 9223372036854775807

	var ua uint8 = 255
	var ub uint16 = 65535
	var uc uint32 = 4294967295
	//var ud uint64 = 18446744073709551615

	var e int = 42
	var f uint = 42

	var g byte = 255
	var h rune = 'A'
	var i uintptr = 0x12345678

	fmt.Printf("Tamaños en bytes:\n")
	fmt.Printf("int8: %d, int16: %d, int32: %d, int64: %d\n",
		unsafe.Sizeof(a), unsafe.Sizeof(b), unsafe.Sizeof(c), unsafe.Sizeof(d))
	fmt.Printf("uint8: %d, uint16: %d, uint32: %d\n",
		unsafe.Sizeof(ua), unsafe.Sizeof(ub), unsafe.Sizeof(uc))
	fmt.Printf("int: %d, uint: %d, uintptr: %d\n",
		unsafe.Sizeof(e), unsafe.Sizeof(f), unsafe.Sizeof(i))

	// Uso de types (suppressing not used warning)
	_ = g
	_ = h
}

func tiposPuntoFlotante() {
	var precio32 float32 = 19.99
	var temperatura32 float32 = 36.5

	var precio64 float64 = 19.99
	var temperatura64 float64 = 36.5

	var pi = 3.14159265359

	fmt.Printf("Precisión float32: %.10f\n", precio32)
	fmt.Printf("Precisión float64: %.10f\n", precio64)

	area := pi * math.Pow(5.0, 2)
	fmt.Printf("Área del círculo: %.2f\n", area)

	// Uso de types
	_ = temperatura32
	_ = temperatura64
}

func tiposComplejos() {
	var c1 complex64 = 3 + 4i
	var c2 complex128 = 5 + 12i
	var c3 = 1 + 2i

	var c4 = complex(3.0, 4.0)
	var c5 = complex(float32(3.0), float32(4.0))

	fmt.Printf("c1: %v, c2: %v, c3: %v\n", c1, c2, c3)
	fmt.Printf("c4: %v, c5: %v\n", c4, c5)

	realPart := real(c2)
	imagPart := imag(c2)
	fmt.Printf("Parte real: %.2f, Parte imaginaria: %.2f\n", realPart, imagPart)

	modulo := cmplx.Abs(c2)
	fase := cmplx.Phase(c2)
	conjugado := cmplx.Conj(c2)

	fmt.Printf("Módulo: %.2f, Fase: %.2f, Conjugado: %v\n", modulo, fase, conjugado)
}

func stringsAndRunes() {
	var saludo string = "Hola"
	var mensaje = "¡Bienvenido a Go!"
	var vacio string

	var ruta = `C:\Users\Juan\Documents\archivo.txt`
	var multilinea = `Este es un
    string de múltiples
    líneas`

	fmt.Printf("Saludo: '%s', Mensaje: '%s', Vacío: '%s'\n", saludo, mensaje, vacio)
	fmt.Printf("Ruta: %s\n", ruta)
	fmt.Printf("Multilínea:\n%s\n", multilinea)

	longitud := len(mensaje)
	bytes := []byte(mensaje)
	fmt.Printf("Longitud en bytes: %d, Bytes: %v\n", longitud, bytes)

	texto := "Hola 世界 🌍"
	fmt.Printf("Texto: %s\n", texto)
	fmt.Printf("Longitud en runes: %d\n", utf8.RuneCountInString(texto))

	for i, r := range texto {
		fmt.Printf("   Posición %d: %c (U+%04X)\n", i, r, r)
	}
}
