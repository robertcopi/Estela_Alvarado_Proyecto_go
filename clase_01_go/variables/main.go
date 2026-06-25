package main

import (
	"fmt"
	"time"
)

var globalCounter int = 0
var variableGlobal = "global"

func main() {
	declaracionBásica()
	alcanceVariables()
	patronesInicializacion()
}

func declaracionBásica() {
	var edad int = 25
	var nombre string = "Juan"
	var activo bool = true

	var contador int
	var mensaje string

	fmt.Printf("Con valores: edad=%d, nombre=%s, activo=%t\n", edad, nombre, activo)
	fmt.Printf("Zero values: contador=%d, mensaje='%s'\n", contador, mensaje)

	var x = 42
	var y = 3.14
	var z = "hello"

	fmt.Printf("x: %T = %v, y: %T = %v, z: %T = %v\n", x, x, y, y, z, z)

	nombrecorto := "Ana"
	edadcorto := 30
	fmt.Printf("Declaración corta: %s, %d\n", nombrecorto, edadcorto)
}

func alcanceVariables() {
	var localVar = "local"
	fmt.Printf("Global: %d, Local: %s\n", globalCounter, localVar)

	if true {
		var blockVar = "bloque"
		fmt.Printf("Desde bloque - Global: %d, Local: %s, Bloque: %s\n", globalCounter, localVar, blockVar)

		var localVar = "sombreada"
		fmt.Printf("Variable sombreada: %s\n", localVar)
	}

	for i := 0; i < 3; i++ {
		var loopVar = fmt.Sprintf("iteración_%d", i)
		fmt.Printf("Loop: i=%d, loopVar=%s\n", i, loopVar)
	}
}

func patronesInicializacion() {
	var tiempo = time.Now()
	var timestamp = tiempo.Unix()
	var año = tiempo.Year()

	var mensaje string
	if año%2 == 0 {
		mensaje = "Año par"
	} else {
		mensaje = "Año impar"
	}

	var configuracion map[string]string
	if configuracion == nil {
		configuracion = make(map[string]string)
		configuracion["env"] = "development"
	}

	fmt.Printf("Tiempo: %v\n", tiempo)
	fmt.Printf("Timestamp: %d, Año: %d\n", timestamp, año)
	fmt.Printf("Mensaje: %s\n", mensaje)
	fmt.Printf("Configuración: %v\n", configuracion)
}
