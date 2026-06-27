package main

import (
	"errors"
	"fmt"
	"strings"
)

// ========= FUNCIONES (SINTAXIS BÁSICA Y RETORNOS) =========

// 1. Función básica que no retorna nada
func saludar(nombre string) {
	fmt.Printf("¡Hola %s, bienvenido a la semana 12!\n", nombre)
}

// 2. Función con un único valor de retorno (sintaxis combinada de tipos 'a, b int')
func sumar(a, b int) int {
	return a + b
}

// 3. Función con múltiples valores de retorno (Manejo de Errores Idiomático)
func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("división por cero no permitida")
	}
	return a / b, nil
}

// 4. Función con valores de retorno nombrados (Named return values)
func analizarTexto(texto string) (palabras int, caracteres int, lineas int) {
	caracteres = len(texto)
	lineas = strings.Count(texto, "\n") + 1
	palabras = len(strings.Fields(texto))
	// Return implícito: devuelve las variables declaradas en la firma
	return
}

func main() {
	fmt.Println("=== 06_FUNCIONES_BASICAS ===")

	// 1. Invocar función simple
	saludar("Desarrollador Go")

	// 2. Único retorno
	fmt.Printf("La suma de 5 y 7 es %d\n", sumar(5, 7))

	// 3. Múltiples retornos
	fmt.Println("\n--- Prueba de División ---")
	resultado, err := dividir(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", resultado)
	}

	resError, errDivCero := dividir(10, 0)
	if errDivCero != nil {
		fmt.Println("Error capturado correctamente:", errDivCero)
	} else {
		fmt.Printf("Resultado inesperado: %.2f\n", resError)
	}

	// 4. Retornos Nombrados
	fmt.Println("\n--- Análisis de Texto ---")
	texto := "Hola mundo\nEste es un texto\nde prueba en Go."
	p, c, l := analizarTexto(texto)
	fmt.Printf("Texto analizado:\n%q\nPalabras: %d, Caracteres: %d, Líneas: %d\n", texto, p, c, l)
}
