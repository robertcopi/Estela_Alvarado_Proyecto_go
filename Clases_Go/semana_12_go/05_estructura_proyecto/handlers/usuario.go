package handlers

import (
	"encoding/json"
	"fmt"

	"semana_12_go/05_estructura_proyecto/internal/usuario"
)

// SimularPeticionUsuario simula lo que haría un http.HandlerFunc para devolver datos al cliente final
func SimularPeticionUsuario() {
	fmt.Println(" -> (handlers) Recibiendo petición web...")
	u := usuario.Obtener()

	datosJSON, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		fmt.Println(" -> (handlers) Error procesando JSON:", err)
		return
	}

	fmt.Printf(" -> (handlers) Respuesta HTTP 200 OK con payload:\n%s\n", string(datosJSON))
}
