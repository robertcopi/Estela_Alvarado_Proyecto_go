package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println(" 🚀 OPTIMIZACIONES DE PERFORMANCE")
	fmt.Println("================================")

	demonstrateSlicePreallocation()
	demonstrateStringBuilding()
	demonstrateTypeConversions()
	demonstrateVariableReuse()
}

func demonstrateSlicePreallocation() {
	fmt.Println("\n--- Preasignación de Slices ---")
	size := 1000

	start := time.Now()
	var badSlice []int
	for i := 0; i < size; i++ {
		badSlice = append(badSlice, i)
	}
	badDuration := time.Since(start)

	start = time.Now()
	goodSlice := make([]int, 0, size)
	for i := 0; i < size; i++ {
		goodSlice = append(goodSlice, i)
	}
	goodDuration := time.Since(start)

	start = time.Now()
	bestSlice := make([]int, size)
	for i := 0; i < size; i++ {
		bestSlice[i] = i
	}
	bestDuration := time.Since(start)

	fmt.Printf("Sin preasignación: %v\n", badDuration)
	fmt.Printf("Con capacidad: %v\n", goodDuration)
	fmt.Printf("Con longitud exacta: %v\n", bestDuration)

	_ = badSlice
	_ = goodSlice
	_ = bestSlice
}

func demonstrateStringBuilding() {
	fmt.Println("\n--- Construcción de Strings ---")
	parts := []string{"Hola", " ", "mundo", " ", "desde", " ", "Go"}

	start := time.Now()
	var badResult string
	for _, part := range parts {
		badResult += part
	}
	badDuration := time.Since(start)

	start = time.Now()
	var builder strings.Builder
	builder.Grow(50)
	for _, part := range parts {
		builder.WriteString(part)
	}
	goodResult := builder.String()
	goodDuration := time.Since(start)

	start = time.Now()
	bestResult := strings.Join(parts, "")
	bestDuration := time.Since(start)

	fmt.Printf("Concatenación: %v -> '%s'\n", badDuration, badResult)
	fmt.Printf("Builder: %v -> '%s'\n", goodDuration, goodResult)
	fmt.Printf("Join: %v -> '%s'\n", bestDuration, bestResult)
}

func demonstrateTypeConversions() {
	fmt.Println("\n--- Conversiones de Tipo ---")
	var numbers []int64 = make([]int64, 1000)
	for i := range numbers {
		numbers[i] = int64(i)
	}

	start := time.Now()
	var badSum int64
	for _, num := range numbers {
		badSum += int64(num)
	}
	badDuration := time.Since(start)

	start = time.Now()
	var goodSum int64
	for _, num := range numbers {
		goodSum += num
	}
	goodDuration := time.Since(start)

	fmt.Printf("Con conversiones innecesarias: %v (suma: %d)\n", badDuration, badSum)
	fmt.Printf("Sin conversiones: %v (suma: %d)\n", goodDuration, goodSum)
}

func demonstrateVariableReuse() {
	fmt.Println("\n--- Reutilización de Variables ---")

	data := make([]map[string]int, 100)
	for i := range data {
		data[i] = map[string]int{"value": i}
	}

	start := time.Now()
	var badTotal int
	for _, item := range data {
		tempValue := item["value"]
		tempSquared := tempValue * tempValue
		badTotal += tempSquared
	}
	badDuration := time.Since(start)

	start = time.Now()
	var goodTotal int
	var value, squared int
	for _, item := range data {
		value = item["value"]
		squared = value * value
		goodTotal += squared
	}
	goodDuration := time.Since(start)

	fmt.Printf("Declaración repetida: %v (total: %d)\n", badDuration, badTotal)
	fmt.Printf("Reutilización: %v (total: %d)\n", goodDuration, goodTotal)
}
