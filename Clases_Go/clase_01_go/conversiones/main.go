package main

import (
	"fmt"
	"strconv"
)

func main() {
	conversionesBasicas()
	conversionesStringNumeros()
	conversionesStringBytesRunes()
}

func conversionesBasicas() {
	var i int = 42
	var f float64 = 3.14

	var result1 = float64(i) + f
	var result2 = i + int(f)

	fmt.Printf("Conversiones numéricas:\n")
	fmt.Printf("int %d + float64 %.2f = %.2f\n", i, f, result1)
	fmt.Printf("int %d + int(%.2f) = %d\n", i, f, result2)

	var a int8 = 100
	var b int16 = 200

	var suma16 = int16(a) + b

	fmt.Printf("Conversiones entre enteros:\n")
	fmt.Printf("int8(%d) + int16(%d) = %d\n", a, b, suma16)

	var grande int64 = 1000000
	var pequeño int8 = int8(grande)

	fmt.Printf("Overflow: int64(%d) -> int8(%d)\n", grande, pequeño)
}

func conversionesStringNumeros() {
	strNumero := "123"
	strFloat := "3.14159"
	strBool := "true"

	numero, err1 := strconv.ParseInt(strNumero, 10, 64)
	if err1 != nil {
		fmt.Printf("Error convertir '%s' a int: %v\n", strNumero, err1)
	} else {
		fmt.Printf("String '%s' -> int64: %d\n", strNumero, numero)
	}

	numeroInt, err2 := strconv.Atoi(strNumero)
	if err2 == nil {
		fmt.Printf("String '%s' -> int: %d\n", strNumero, numeroInt)
	}

	flotante, err3 := strconv.ParseFloat(strFloat, 64)
	if err3 == nil {
		fmt.Printf("String '%s' -> float64: %.5f\n", strFloat, flotante)
	}

	booleano, err4 := strconv.ParseBool(strBool)
	if err4 == nil {
		fmt.Printf("String '%s' -> bool: %t\n", strBool, booleano)
	}

	var entero int = 456
	strDesdeInt := strconv.Itoa(entero)
	fmt.Printf("int %d -> string: '%s'\n", entero, strDesdeInt)

	strBinario := "1010"
	binario, _ := strconv.ParseInt(strBinario, 2, 64)
	fmt.Printf("Binario '%s' (base 2) -> %d\n", strBinario, binario)
}

func conversionesStringBytesRunes() {
	texto := "Hola 世界 🌍"

	bytes := []byte(texto)
	fmt.Printf("String: '%s'\n", texto)
	fmt.Printf("Bytes: %v\n", bytes)

	textoRecuperado := string(bytes)
	fmt.Printf("Bytes de vuelta a string: '%s'\n", textoRecuperado)

	runes := []rune(texto)
	fmt.Printf("Runes: %v\n", runes)

	textoDesdeRunes := string(runes)
	fmt.Printf("Runes de vuelta a string: '%s'\n", textoDesdeRunes)
}
