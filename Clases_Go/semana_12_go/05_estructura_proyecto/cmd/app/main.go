package main

import (
	"fmt"
	"semana_12_go/05_estructura_proyecto/handlers"
	"semana_12_go/05_estructura_proyecto/internal/usuario"
	"semana_12_go/05_estructura_proyecto/pkg/logger"
)

func main() {
	logger.Info("Iniciando aplicación estructurada...")
	fmt.Println("=== ESTRUCTURA DE PROYECTOS GO ===")

	// Utilizando capa de servicio interno
	usuario.Ejecutar()

	// Simulación de handler/controlador HTTP
	handlers.SimularPeticionUsuario()

	logger.Info("Aplicación finalizada correctamente.")
}
