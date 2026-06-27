package main

import (
    "fmt"
    "math/rand"
    "strings"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    secreto := rand.Intn(10) + 1
    var input string
    
    jugadores := []string{"Jugador  1", "Jugador  2"}
    intentosMax := 3
    intentos := map[string]int{"Jugador  1": 0, "Jugador  2": 0}
    
    fmt.Println(" 🎮  ¡Bienvenidos  al  juego  de  adivinanza!")
    fmt.Println(" 🔢  Adivina  el  número  secreto  entre  1  y  10.  Escribe  'salir'  para  terminar.")
    
JUEGO:
    for {
        for _, jugador := range jugadores {
            if intentos[jugador] >= intentosMax {
                continue
            }
            
            fmt.Printf(" 👉  %s,  intento  %d:  ", jugador, intentos[jugador]+1)
            fmt.Scanln(&input)
            
            if strings.ToLower(input) == "salir" {
                fmt.Println(" 🚪  El  juego  ha  sido  cancelado  por  el  usuario.")
                goto FIN
            }
            
            var guess int
            _, err := fmt.Sscanf(input, "%d", &guess)
            if err != nil {
                fmt.Println(" ❌  Entrada  no  válida.  Escribe  un  número.")
                continue
            }
            
            if guess%2 == 0 {
                fmt.Println(" ⚠  Los  pares  no  traen  suerte.  Intenta  con  otro  número  impar.")
                continue
            }
            
            intentos[jugador]++
            
            if guess == secreto {
                fmt.Printf(" 🎉  ¡%s  adivinó  el  número  secreto!  Era  %d  🎯\n", jugador, secreto)
                break JUEGO
            } else {
                fmt.Println(" ❌  Incorrecto.  Sigue  intentando.")
            }
        }
        //  Verifica  si  ambos  jugadores  agotaron  sus  intentos  
        if intentos["Jugador  1"] >= intentosMax && intentos["Jugador  2"] >= intentosMax {
            fmt.Println(" 😢  Ambos  jugadores  agotaron  sus  intentos.")
            break
        }
    }
FIN:
    fmt.Println(" 🎯  Fin  del  juego.  El  número  secreto  era:", secreto)
}
