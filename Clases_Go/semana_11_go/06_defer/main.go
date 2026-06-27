package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("===  DEFER  ===")
    
    //  DEFER  BÁSICO  
    demonstrarDeferBasico()
    
    //  MÚLTIPLES  DEFERS  
    demonstrarMultiplesDefers()
    
    //  DEFER  CON  VALORES  
    demonstrarDeferConValores()
    
    //  CASOS  PRÁCTICOS  
    demonstrarCasosPracticosDefer()
}

func demonstrarDeferBasico() {
    fmt.Println("---  Defer  básico  ---")
    
    fmt.Println("1.  Inicio  de  función")
    
    defer fmt.Println("4.  Este  mensaje  se  ejecuta  al  final  (defer)")
    
    fmt.Println("2.  En  medio  de  función")
    fmt.Println("3.  Antes  del  return")
    
    //  El  defer  se  ejecuta  aquí  automáticamente  
}

func demonstrarMultiplesDefers() {
    fmt.Println("\n---  Múltiples  defers  (LIFO  -  Last  In,  First  Out)  ---")
    
    defer fmt.Println(" 🥉  Tercer  defer  (se  ejecuta  primero)")
    defer fmt.Println(" 🥈  Segundo  defer  (se  ejecuta  segundo)")
    defer fmt.Println(" 🥇  Primer  defer  (se  ejecuta  último)")
    
    fmt.Println("Código  normal  ejecutándose...")
}

func demonstrarDeferConValores() {
    fmt.Println("\n---  Defer  con  valores  capturados  ---")
    
    x := 10
    defer fmt.Printf("Valor  de  x  en  defer:  %d  (capturado  al  definir  defer)\n", x)
    
    x = 20
    fmt.Printf("Valor  actual  de  x:  %d\n", x)
    
    //  El  defer  usará  x=10,  no  x=20  
    
    //  Para  usar  el  valor  actual,  usar  función  anónima  
    defer func() {
        fmt.Printf("Valor  actual  de  x  en  defer  con  closure:  %d\n", x)
    }()
    
    x = 30
    fmt.Printf("Valor  final  de  x:  %d\n", x)
}

func demonstrarCasosPracticosDefer() {
    fmt.Println("\n---  Casos  prácticos  con  defer  ---")
    
    //  1.  Manejo  de  archivos  
    fmt.Println("1.  Manejo  de  archivos:")
    manejarArchivo()
    
    //  2.  Medición  de  tiempo  
    fmt.Println("\n2.  Medición  de  tiempo:")
    medirTiempoEjecucion()
    
    //  3.  Cleanup  de  recursos  
    fmt.Println("\n3.  Cleanup  de  recursos:")
    simularConexionDB()
    
    //  4.  Logging  de  entrada  y  salida  
    fmt.Println("\n4.  Logging:")
    funcionConLogging("parametro_importante")
    
    //  5.  Mutex  unlocking  
    fmt.Println("\n5.  Manejo  de  mutex:")
    simularMutex()
}

func manejarArchivo() {
    //  Simular  apertura  de  archivo  
    fmt.Println("   📂  Abriendo  archivo...")
    
    //  defer  se  ejecuta  incluso  si  hay  error  
    defer fmt.Println("   🔒  Cerrando  archivo  (defer)")
    
    //  Simular  trabajo  con  archivo  
    fmt.Println("   📝  Escribiendo  datos...")
    fmt.Println("   📖  Leyendo  datos...")
    
    //  Si  hubiera  un  error  aquí,  defer  aún  se  ejecutaría  
    //  return  //  El  defer  se  ejecuta  antes  del  return  
}

func medirTiempoEjecucion() {
    inicio := time.Now()
    
    defer func() {
        duracion := time.Since(inicio)
        fmt.Printf("   ⏱  Función  tardó:  %v\n", duracion)
    }()
    
    fmt.Println("   🔄  Iniciando  operación  costosa...")
    time.Sleep(100 * time.Millisecond) //  Simular  trabajo  
    fmt.Println("   ✅  Operación  completada")
}

func simularConexionDB() {
    fmt.Println("   🔌  Conectando  a  base  de  datos...")
    
    defer fmt.Println("   🔌  Desconectando  de  base  de  datos  (defer)")
    
    //  Simular  múltiples  operaciones  
    fmt.Println("   📊  Ejecutando  query  1...")
    fmt.Println("   📊  Ejecutando  query  2...")
    fmt.Println("   📊  Ejecutando  query  3...")
}

func funcionConLogging(parametro string) {
    fmt.Printf("   📥  ENTRADA:  funcionConLogging(%s)\n", parametro)
    
    defer fmt.Println("   📤  SALIDA:  funcionConLogging")
    
    //  Lógica  de  la  función  
    fmt.Println("   ⚙  Procesando  lógica  de  negocio...")
    
    if parametro == "error" {
        fmt.Println("   ❌  Error  simulado")
        return //  defer  aún  se  ejecuta  
    }
    
    fmt.Println("   ✅  Procesamiento  exitoso")
}

func simularMutex() {
    fmt.Println("   🔐  Adquiriendo  lock...")
    
    defer fmt.Println("   🔓  Liberando  lock  (defer)")
    
    //  Simular  trabajo  en  sección  crítica  
    fmt.Println("   ⚙  Trabajando  en  sección  crítica...")
    time.Sleep(50 * time.Millisecond)
}
