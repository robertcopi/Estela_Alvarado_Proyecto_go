package main

import (
	"fmt"
)

type Persona struct {
	Nombre string
	Edad   int
	Email  string
}

func (p *Persona) Saludar() string {
	return fmt.Sprintf("Hola, soy %s", p.Nombre)
}
func main() {
	if true {
		fmt.Println("Mal formateado")
	}
	for i := 0; i < 5; i++ {
		fmt.Printf("Número: %d\n", i)
	}
	p := &Persona{
		Nombre: "Juan",
		Edad:   30,
		Email:  "juan@ejemplo.com"}
	fmt.Println(p.Saludar())
	numeros := []int{1, 2, 3, 4, 5}
	mapa := map[string]int{"uno": 1, "dos": 2, "tres": 3}
	matriz := [][]int{{1, 2}, {3, 4}, {5, 6}}
	for _, num := range numeros {
		if num%2 == 0 {
			fmt.Printf("Par: %d\n", num)
		} else {
			fmt.Printf("Impar: %d\n", num)
		}
	}
	_ = mapa
	_ = matriz
}
