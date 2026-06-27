package usuario

import (
	"fmt"
	"semana_12_go/05_estructura_proyecto/models"
)

// La capa internal maneja lógica de negocio que no debería exponerse a otros módulos/proyectos

func Ejecutar() {
	fmt.Println(" -> (internal/usuario) Servicio de usuarios ejecutándose...")
}

// Obtener retorna un modelo que puede viajar hacia los handlers o repositorios
func Obtener() models.Usuario {
	return models.Usuario{
		ID:     101,
		Nombre: "Ana Torres",
		Correo: "ana@empresa.com",
		Activo: true,
	}
}
