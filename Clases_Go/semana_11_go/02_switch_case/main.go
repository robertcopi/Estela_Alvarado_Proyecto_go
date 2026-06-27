package main

import (
    "fmt"
    "runtime"
    "time"
)

func main() {
    fmt.Println("===  ESTRUCTURAS  SWITCH  ===")
    
    //  SWITCH  BÁSICO  
    dia := time.Now().Weekday()
    
    switch dia {
    case time.Monday:
        fmt.Println(" 😴  Lunes  -  Inicio  de  semana")
    case time.Tuesday:
        fmt.Println(" 💪  Martes  -  A  trabajar")
    case time.Wednesday:
        fmt.Println(" 🐪  Miércoles  -  Mitad  de  semana")
    case time.Thursday:
        fmt.Println(" 🚀  Jueves  -  Casi  llegamos")
    case time.Friday:
        fmt.Println(" 🎉  Viernes  -  ¡Fin  de  semana  próximo!")
    case time.Saturday, time.Sunday:
        fmt.Println(" 🏖  Fin  de  semana")
    default:
        fmt.Println(" 🤔  Día  desconocido")
    }
    
    //  SWITCH  CON  INICIALIZACIÓN  
    switch mes := time.Now().Month(); mes {
    case time.December, time.January, time.February:
        fmt.Println(" ❄  Época  de  verano  (Hemisferio  Sur)")
    case time.March, time.April, time.May:
        fmt.Println(" 🍂  Otoño")
    case time.June, time.July, time.August:
        fmt.Println("🧥  Invierno")
    case time.September, time.October, time.November:
        fmt.Println(" 🌸  Primavera")
    }
    
    //  SWITCH  SIN  EXPRESIÓN  (actúa  como  if-else-if)  
    hora := time.Now().Hour()
    temperatura := 22.0
    
    switch {
    case hora < 6:
        fmt.Println(" 🌃  Madrugada")
    case hora < 12 && temperatura > 20:
        fmt.Println(" 🌞  Mañana  agradable")
    case hora < 12:
        fmt.Println(" 🌅  Mañana  fresca")
    case hora < 18 && temperatura > 25:
        fmt.Println(" ☀  Tarde  calurosa")
    case hora < 18:
        fmt.Println(" 🌤  Tarde  normal")
    default:
        fmt.Println(" 🌙  Noche")
    }
    
    //  SWITCH  CON  FALLTHROUGH  (poco  común)  
    numero := 3
    
    switch numero {
    case 1:
        fmt.Print("uno")
        fallthrough
    case 2:
        fmt.Print("dos")
        fallthrough
    case 3:
        fmt.Print("tres")
        fallthrough
    case 4:
        fmt.Print("cuatro")
    }
    fmt.Println() //  Nueva  línea  
    
    //  SWITCH  CON  TYPE  ASSERTION  
    var interfaz interface{} = "texto"
    
    switch valor := interfaz.(type) {
    case string:
        fmt.Printf("Es  string:  '%s'  (longitud:  %d)\n", valor, len(valor))
    case int:
        fmt.Printf("Es  entero:  %d\n", valor)
    case float64:
        fmt.Printf("Es  float:  %.2f\n", valor)
    case bool:
        fmt.Printf("Es  booleano:  %t\n", valor)
    case nil:
        fmt.Println("Es  nil")
    default:
        fmt.Printf("Tipo  desconocido:  %T\n", valor)
    }
    
    //  CASOS  PRÁCTICOS  CON  SWITCH  
    demonstrarCasosPracticosSwitch()
}

func demonstrarCasosPracticosSwitch() {
    fmt.Println("\n---  Casos  prácticos  con  switch  ---")
    
    //  1.  Procesamiento  de  códigos  de  estado  HTTP  
    statusCode := 404
    
    switch statusCode {
    case 200:
        fmt.Println(" ✅  OK")
    case 201:
        fmt.Println(" ✅  Creado")
    case 400:
        fmt.Println(" ❌  Petición  incorrecta")
    case 401:
        fmt.Println(" 🔐  No  autorizado")
    case 403:
        fmt.Println(" 🚫  Prohibido")
    case 404:
        fmt.Println(" 🔍  No  encontrado")
    case 500:
        fmt.Println(" 💥  Error  interno  del  servidor")
    default:
        if statusCode >= 200 && statusCode < 300 {
            fmt.Println(" ✅  Éxito")
        } else if statusCode >= 400 && statusCode < 500 {
            fmt.Println(" ❌  Error  del  cliente")
        } else if statusCode >= 500 {
            fmt.Println(" 💥  Error  del  servidor")
        } else {
            fmt.Printf(" 🤔  Código  desconocido:  %d\n", statusCode)
        }
    }
    
    //  2.  Categorización  de  archivos  por  extensión  
    filename := "documento.pdf"
    extension := filename[len(filename)-3:]
    
    switch extension {
    case "pdf":
        fmt.Println(" 📄  Documento  PDF")
    case "doc", "docx":
        fmt.Println(" 📝  Documento  de  Word")
    case "xls", "xlsx":
        fmt.Println(" 📊  Hoja  de  cálculo")
    case "jpg", "png", "gif":
        fmt.Println(" 🖼  Imagen")
    case "mp4", "avi", "mov":
        fmt.Println(" 🎬  Video")
    case "mp3", "wav", "flac":
        fmt.Println(" 🎵  Audio")
    default:
        fmt.Printf(" 📁  Archivo  de  tipo:  %s\n", extension)
    }
    
    //  3.  Lógica  de  permisos  por  rol  
    rol := "admin"
    accion := "delete_user"
    
    switch rol {
    case "super_admin":
        fmt.Println(" 🔑  Acceso  total  -  Todas  las  acciones  permitidas")
    case "admin":
        switch accion {
        case "create_user", "edit_user", "view_user":
            fmt.Println(" ✅  Acción  permitida  para  admin")
        case "delete_user":
            fmt.Println(" ⚠  Acción  sensible  -  Requiere  confirmación")
        default:
            fmt.Println(" ❌  Acción  no  permitida  para  admin")
        }
    case "moderator":
        switch accion {
        case "view_user", "edit_user":
            fmt.Println(" ✅  Acción  permitida  para  moderador")
        default:
            fmt.Println(" ❌  Acción  no  permitida  para  moderador")
        }
    case "user":
        switch accion {
        case "view_user":
            fmt.Println(" ✅  Solo  visualización  permitida")
        default:
            fmt.Println(" ❌  Acción  no  permitida  para  usuario  regular")
        }
    default:
        fmt.Println(" ❌  Rol  no  reconocido")
    }
    
    //  4.  Procesamiento  por  sistema  operativo  
    os := runtime.GOOS
    
    switch os {
    case "linux":
        fmt.Println(" 🐧  Configuración  para  Linux")
        configurarLinux()
    case "darwin":
        fmt.Println(" 🍎  Configuración  para  macOS")
        configurarMacOS()
    case "windows":
        fmt.Println("🪟  Configuración  para  Windows")
        configurarWindows()
    default:
        fmt.Printf(" 🤔  Sistema  operativo  no  soportado:  %s\n", os)
    }
    
    //  5.  State  machine  simple  
    estado := "inicio"
    evento := "login_exitoso"
    
    nuevoEstado := procesarEstado(estado, evento)
    fmt.Printf("Estado:  %s  ->  Evento:  %s  ->  Nuevo  Estado:  %s\n", estado, evento, nuevoEstado)
}

func configurarLinux() {
    fmt.Println("   -  Configurando  paths  de  Linux")
    fmt.Println("   -  Estableciendo  permisos  UNIX")
}

func configurarMacOS() {
    fmt.Println("   -  Configurando  paths  de  macOS")
    fmt.Println("   -  Configurando  Keychain")
}

func configurarWindows() {
    fmt.Println("   -  Configurando  paths  de  Windows")
    fmt.Println("   -  Configurando  Registry")
}

func procesarEstado(estadoActual, evento string) string {
    switch estadoActual {
    case "inicio":
        switch evento {
        case "login_exitoso":
            return "autenticado"
        case "registro":
            return "registrando"
        default:
            return "inicio"
        }
    case "autenticado":
        switch evento {
        case "logout":
            return "inicio"
        case "timeout":
            return "sesion_expirada"
        default:
            return "autenticado"
        }
    case "sesion_expirada":
        switch evento {
        case "relogin":
            return "autenticado"
        case "timeout_final":
            return "inicio"
        default:
            return "sesion_expirada"
        }
    default:
        return "inicio"
    }
}
