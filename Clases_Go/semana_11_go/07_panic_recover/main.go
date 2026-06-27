package main

import (
    "fmt"
)

func main() {
    fmt.Println("===  PANIC  Y  RECOVER  ===")
    
    //  PANIC  BÁSICO  
    demonstrarPanicBasico()
    
    //  RECOVER  PARA  MANEJAR  PANIC  
    demonstrarRecover()
    
    //  CASOS  PRÁCTICOS  
    demonstrarCasosPracticosPanicRecover()
}

func demonstrarPanicBasico() {
    fmt.Println("---  Panic  básico  ---")
    
    //  defer  se  ejecuta  incluso  con  panic  
    defer fmt.Println("3.  Defer  ejecutándose  durante  panic")
    
    fmt.Println("1.  Antes  del  panic")
    fmt.Println("2.  Justo  antes  del  panic")
    
    //  Este  panic  terminará  el  programa  si  no  se  recupera  
    //  panic("¡Algo  salió  terriblemente  mal!")  
    
    fmt.Println("Esta  línea  nunca  se  ejecutaría")
}

func demonstrarRecover() {
    fmt.Println("\n---  Recover  para  manejar  panic  ---")
    
    //  Función  que  puede  hacer  panic  
    funcionPeligrosa := func() {
        defer func() {
            if r := recover(); r != nil {
                fmt.Printf("   🚨  Panic  recuperado:  %v\n", r)
                fmt.Println("   🔄  Continuando  ejecución  normal...")
            }
        }()
        
        fmt.Println("   ⚙  Iniciando  operación  peligrosa...")
        panic("¡Error  crítico  simulado!")
        fmt.Println("   Esta  línea  nunca  se  ejecuta")
    }
    
    fmt.Println("1.  Antes  de  función  peligrosa")
    funcionPeligrosa()
    fmt.Println("2.  Después  de  función  peligrosa  (recuperada)")
    fmt.Println("3.  El  programa  continúa  normalmente")
}

func demonstrarCasosPracticosPanicRecover() {
    fmt.Println("\n---  Casos  prácticos  ---")
    
    //  1.  Servidor  web  que  no  debe  caerse  
    fmt.Println("1.  Simulación  de  servidor  web:")
    simularServidorWeb()
    
    //  2.  Validación  estricta  
    fmt.Println("\n2.  Validación  con  panic/recover:")
    testValidacion()
    
    //  3.  Procesamiento  de  datos  con  recovery  
    fmt.Println("\n3.  Procesamiento  de  lote  con  recovery:")
    procesarLoteDatos()
    
    //  4.  División  segura  
    fmt.Println("\n4.  División  segura:")
    testDivisionSegura()
}

func simularServidorWeb() {
    //  Simular  múltiples  requests  
    requests := []string{"GET  /users", "POST  /users", "GET  /invalid", "DELETE  /users/1"}
    
    for i, request := range requests {
        func() {
            defer func() {
                if r := recover(); r != nil {
                    fmt.Printf("   🚨  Request  %d  falló:  %v\n", i+1, r)
                    fmt.Println("   📝  Logging  error  y  continuando...")
                }
            }()
            
            fmt.Printf("   📥  Procesando  request  %d:  %s\n", i+1, request)
            
            //  Simular  error  en  request  específico  
            if request == "GET  /invalid" {
                panic("endpoint  no  válido")
            }
            
            fmt.Printf("   ✅  Request  %d  completado  exitosamente\n", i+1)
        }()
    }
    
    fmt.Println("   🌐  Servidor  continúa  funcionando")
}

func testValidacion() {
    usuarios := []struct {
        Nombre string
        Edad   int
        Email  string
    }{
        {"Ana", 25, "ana@email.com"},
        {"", 30, "luis@email.com"},        //  Error:  nombre  vacío  
        {"María", -5, "maria@email.com"},  //  Error:  edad  negativa  
        {"Carlos", 35, "carlos@email.com"},
    }
    
    for i, usuario := range usuarios {
        func() {
            defer func() {
                if r := recover(); r != nil {
                    fmt.Printf("   ❌  Usuario  %d  inválido:  %v\n", i+1, r)
                }
            }()
            
            validarUsuario(usuario.Nombre, usuario.Edad, usuario.Email)
            fmt.Printf("   ✅  Usuario  %d  válido:  %s\n", i+1, usuario.Nombre)
        }()
    }
}

func validarUsuario(nombre string, edad int, email string) {
    if nombre == "" {
        panic("nombre  no  puede  estar  vacío")
    }
    if edad < 0 {
        panic("edad  no  puede  ser  negativa")
    }
    if email == "" {
        panic("email  no  puede  estar  vacío")
    }
}

func procesarLoteDatos() {
    datos := []interface{}{1, "texto", 3.14, []int{1, 2, 3}, nil, 42}
    resultados := make([]string, 0)
    
    for i, dato := range datos {
        func() {
            defer func() {
                if r := recover(); r != nil {
                    fmt.Printf("   ⚠  Error  procesando  elemento  %d:  %v\n", i, r)
                    resultados = append(resultados, fmt.Sprintf("ERROR_%d", i))
                }
            }()
            
            resultado := procesarDato(dato)
            resultados = append(resultados, resultado)
            fmt.Printf("   ✅  Elemento  %d  procesado:  %s\n", i, resultado)
        }()
    }
    
    fmt.Printf("   📊  Resultados  finales:  %v\n", resultados)
}

func procesarDato(dato interface{}) string {
    switch v := dato.(type) {
    case int:
        return fmt.Sprintf("INT_%d", v*2)
    case string:
        return fmt.Sprintf("STR_%s", v)
    case float64:
        return fmt.Sprintf("FLOAT_%.2f", v)
    case nil:
        panic("no  se  puede  procesar  nil")
    default:
        panic(fmt.Sprintf("tipo  no  soportado:  %T", v))
    }
}

func testDivisionSegura() {
    operaciones := []struct {
        a, b float64
    }{
        {10, 2},
        {15, 3},
        {20, 0}, //  División  por  cero  
        {25, 5},
    }
    
    for i, op := range operaciones {
        resultado := divisionSegura(op.a, op.b)
        fmt.Printf("   %.1f  ÷  %.1f  =  %s\n", op.a, op.b, resultado)
    }
}

func divisionSegura(a, b float64) string {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Error  en  división:  %v", r)
        }
    }()
    
    if b == 0 {
        panic("división  por  cero")
    }
    
    resultado := a / b
    return fmt.Sprintf("%.2f", resultado)
}

//  Función  utilitaria  para  demostrar  panic  con  stack  trace  
func demonstrarStackTrace() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Panic  recuperado:  %v\n", r)
            //  En  aplicaciones  reales,  aquí  podrías  imprimir  el  stack  trace  completo
        }
    }()
    
    funcionNivel1()
}

func funcionNivel1() {
    funcionNivel2()
}

func funcionNivel2() {
    funcionNivel3()
}

func funcionNivel3() {
    panic("Error  en  función  nivel  3")
}
