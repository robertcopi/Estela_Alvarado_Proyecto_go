package main

import (
	"fmt"
)

// ========= CLOSURES EN GO =========

// 1. Fábrica de cierres (Closure Factory)
// Retorna una función anónima que mantiene una referencia a la variable `count`.
func createCounter() func() int {
	count := 0 // Variable privada encapsulada dentro de la función y de la cual el compilador asegura retención
	increment := func() int {
		count++
		return count
	}
	return increment
}

// 2. Closure aplicado a acumulación de saldos
func nuevoGeneradorSaldo(saldoInicial float64) func(float64) float64 {
	saldo := saldoInicial
	return func(cambio float64) float64 {
		saldo += cambio
		return saldo
	}
}

func main() {
	fmt.Println("=== 09_CLOSURES ===")

	fmt.Println("\n--- 1. Creador de Contadores Múltiples ---")
	// Cada ejecución de createCounter genera un ámbito de "count" distinto en memoria.
	counter1 := createCounter()
	counter2 := createCounter()

	fmt.Println("counter1 invocación 1:", counter1()) // 1
	fmt.Println("counter1 invocación 2:", counter1()) // 2

	// counter2 no está afectado por counter1
	fmt.Println("counter2 invocación 1:", counter2()) // 1
	fmt.Println("counter2 invocación 2:", counter2()) // 2

	fmt.Println("counter1 invocación 3:", counter1()) // 3

	fmt.Println("\n--- 2. Generador de Saldos ---")
	cajaChica := nuevoGeneradorSaldo(100.0)
	cajaFuerte := nuevoGeneradorSaldo(5000.0)

	fmt.Printf("Caja Chica inicial: %.2f\n", cajaChica(0))      // 100
	fmt.Printf("Añadir 50 a Caja Chica: %.2f\n", cajaChica(50)) // 150

	fmt.Printf("Gastar 1000 en Fuerte: %.2f\n", cajaFuerte(-1000)) // 4000
}
