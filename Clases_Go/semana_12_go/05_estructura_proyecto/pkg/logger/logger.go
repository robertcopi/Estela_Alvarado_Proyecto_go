package logger

import "log"

// Funciones expuestas públicamente (Upper case) que pueden ser usadas por cualquier capa
func Info(msg string) {
	log.Println("[INFO]", msg)
}

func Error(msg string) {
	log.Println("[ERROR]", msg)
}
