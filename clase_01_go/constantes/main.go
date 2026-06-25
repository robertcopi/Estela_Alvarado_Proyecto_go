package main

import (
	"fmt"
)

const CompanyName = "TechCorp"
const version = "1.0.0"

func main() {
	constantesBasicas()
	enumeraciones()
	constantesComplejas()
}

func constantesBasicas() {
	const pi = 3.14159
	const mensaje = "¡Hola, Go!"
	const activo = true

	fmt.Printf("Pi: %T = %v\n", pi, pi)
	fmt.Printf("Mensaje: %T = %v\n", mensaje, mensaje)
	fmt.Printf("Activo: %T = %v\n", activo, activo)

	const tiempoCompilacion = "Compilado el 2024"
	fmt.Printf("Constantes de empresa: %s v%s\n", CompanyName, version)
	fmt.Printf("Tiempo: %s\n", tiempoCompilacion)

	const a = 42
	const d int = 42
	var x1 int = a
	var x2 int64 = a

	fmt.Printf("x1: %T = %v\n", x1, x1)
	fmt.Printf("x2: %T = %v\n", x2, x2)
	_ = d
}

func enumeraciones() {
	const (
		Lunes = iota
		Martes
		Miercoles
		Jueves
		Viernes
		Sabado
		Domingo
	)

	fmt.Printf("Días de la semana:\n")
	fmt.Printf("Lunes: %d, Martes: %d, Miércoles: %d\n", Lunes, Martes, Miercoles)

	const (
		StatusInactivo = iota + 1
		StatusActivo
		StatusSuspendido
		StatusBloqueado
	)

	fmt.Printf("Estados:\n")
	fmt.Printf("Inactivo: %d, Activo: %d, Suspendido: %d, Bloqueado: %d\n",
		StatusInactivo, StatusActivo, StatusSuspendido, StatusBloqueado)

	const (
		Read = 1 << iota
		Write
		Execute
	)
	fmt.Printf("Permisos (flags):\n")
	fmt.Printf("Read: %d, Write: %d, Execute: %d\n", Read, Write, Execute)

	const ReadWrite = Read | Write
	const FullAccess = Read | Write | Execute

	fmt.Printf("Permisos combinados:\n")
	fmt.Printf("ReadWrite: %d, FullAccess: %d\n", ReadWrite, FullAccess)
}

func constantesComplejas() {
	const (
		Pi    = 3.14159265358979323846
		E     = 2.71828182845904523536
		Phi   = 1.61803398874989484820
		Sqrt2 = 1.41421356237309504880
	)

	const (
		KiB = 1024
		MiB = KiB * 1024
		GiB = MiB * 1024
		TiB = GiB * 1024
	)

	const (
		MaxUsers       = 1000
		SessionTimeout = 30 * 60
		RetryAttempts  = 3
		BufferSize     = 8 * KiB
	)

	fmt.Printf("Constantes matemáticas:\n")
	fmt.Printf("π = %.10f, e = %.10f, φ = %.10f\n", Pi, E, Phi)

	fmt.Printf("\nConstantes de almacenamiento:\n")
	fmt.Printf("1 KiB = %d bytes\n", KiB)
	fmt.Printf("1 MiB = %d bytes\n", MiB)

	fmt.Printf("\nConstantes de configuración:\n")
	fmt.Printf("Max usuarios: %d, Timeout: %ds, Buffer: %d bytes\n", MaxUsers, SessionTimeout, BufferSize)
}
