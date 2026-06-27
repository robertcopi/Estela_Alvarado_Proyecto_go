package main

import (
	"fmt"
	"strings"
	"time"
)

// ========= FUNCIONES VARIÁDICAS =========

// 1. Función variádica básica con rango variable de float64
func calcularPromedio(numeros ...float64) float64 {
	if len(numeros) == 0 {
		return 0
	}

	var suma float64
	for _, num := range numeros {
		suma += num
	}
	return suma / float64(len(numeros))
}

// 2. Función variádica con parámetros fijos al inicio
func LogActividad(nivel string, usuario string, acciones ...string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] [%s] - Usuario: %s\n", timestamp, nivel, usuario)

	if len(acciones) == 0 {
		fmt.Println("   (Sin acciones registradas)")
		return
	}

	for i, accion := range acciones {
		fmt.Printf("   %d. %s\n", i+1, accion)
	}
	fmt.Println("--------------------------------------------------")
}

// 3. Función variádica avanzada pasando funciones de retorno de texto () -> string
func GenerarReporte(titulo string, secciones ...func() string) string {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("=== %s ===\n", strings.ToUpper(titulo)))
	builder.WriteString(fmt.Sprintf("Generado: %s\n\n", time.Now().Format("02/01/2006 15:04:05")))

	for i, seccion := range secciones {
		builder.WriteString(fmt.Sprintf("--- Sección %d ---\n", i+1))
		builder.WriteString(seccion())
		builder.WriteString("\n\n")
	}

	return builder.String()
}

func main() {
	fmt.Println("=== 07_FUNCIONES_VARIADICAS ===")

	fmt.Println("\n--- 1. calcularPromedio() ---")
	prom1 := calcularPromedio(8.5, 9.0, 7.5, 8.0)
	prom2 := calcularPromedio() // No hay error aunque esté vacío
	fmt.Printf("Promedio de 4 notas: %.2f\n", prom1)
	fmt.Printf("Promedio sin notas: %.2f\n", prom2)

	fmt.Println("\n--- 2. LogActividad() ---")
	// Diferente cantidad de parámetros dinámicos
	LogActividad("INFO", "roberto")
	LogActividad("WARN", "ana_t", "Cambio de contraseña")
	LogActividad("AUDIT", "sysadmin", "Backup DB", "Actualización sistema", "Limpieza caché")

	// 3. Desempaquetado de Slices usando el operador variádico ...
	fmt.Println("\n--- 3. Desempaquetado de Slices ---")
	accionesProgramadas := []string{"Mantenimiento servidor", "Envío correos marketing", "Generación reportes"}
	// expande el slice al desempaquetarlo para que pase como variádico arguments
	LogActividad("SYSTEM", "cronjob", accionesProgramadas...)

	// 4. Reporte Generator usando funciones variádicas anonimas
	fmt.Println("\n--- 4. Generar Reporte con variádicos func() ---")
	sec1 := func() string { return "Las metas se alcanzaron al 95%." }
	sec2 := func() string { return "Se redujo el Opex un 15%." }
	sec3 := func() string { return "El equipo de desarrollo finalizó 4 módulos." }

	reporteCompleto := GenerarReporte("Resumen Anual", sec1, sec2, sec3)
	fmt.Println(reporteCompleto)
}
