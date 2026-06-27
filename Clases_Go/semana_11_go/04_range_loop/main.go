package main

import (
    "fmt"
    "strings"
)

func main() {
    opiniones := []string{
        "El  servicio  fue  bueno  y  rápido",
        "El  sistema  es  muy  lento  y  malo",
        "Buen  producto  pero  entrega  lenta",
        "Rápido,  eficiente  y  bueno",
        "Malo  servicio,  lento  y  sin  soporte",
    }
    
    palabrasClave := []string{"bueno", "malo", "rápido", "lento"}
    
    //  Inicializar  el  mapa  de  conteo  
    conteo := make(map[string]int)
    for _, clave := range palabrasClave {
        conteo[clave] = 0
    }
    
    //  Procesar  opiniones  
    for _, opinion := range opiniones {
        //  Convertimos  a  minúsculas  para  uniformidad  
        palabras := strings.Fields(strings.ToLower(opinion))
        
        for _, palabra := range palabras {
            //  Limpiar  comas  o  puntos  si  los  hubiera  (básico)  
            palabra = strings.Trim(palabra, ".,;")
            
            //  Verificamos  si  es  una  palabra  clave  
            if _, existe := conteo[palabra]; existe {
                conteo[palabra]++
            }
        }
    }
    
    //  Mostrar  resultados  
    fmt.Println(" 📊  Conteo  de  palabras  clave:")
    for palabra, cantidad := range conteo {
        fmt.Printf("-  %s:  %d  veces\n", palabra, cantidad)
    }
}
