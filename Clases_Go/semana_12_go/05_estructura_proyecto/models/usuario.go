package models

type Usuario struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
	Correo string `json:"correo"`
	Activo bool   `json:"activo"`
}
