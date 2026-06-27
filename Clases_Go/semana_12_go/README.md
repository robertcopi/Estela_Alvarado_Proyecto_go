# Go Semana 12: Funciones, Structs, Pointers e Interfaces

Este repositorio contiene los ejercicios resueltos de la **Semana 12**, implementando todos los módulos prácticos explicados en la clase "Funciones y Structs.pdf". Cada concepto ha sido debidamente aislado en su propio módulo ejecutable (`main.go`).

## Estructura del Proyecto

```text
semana_12_go/
├── 01_structs/                 # Structs básicos, structs anónimos y valores de campos.
├── 02_receptores_metodos/      # Receivers de Valor (Value Receivers) y Receivers de Puntero (Pointer Receivers). Sistema de biblioteca.
├── 03_punteros/                # Operaciones de memoria, ref y unref con *, &, y la función new().
├── 04_interfaces/              # Polimorfismo mediante el sistema de notificaciones. Empleo de Type Assertions y Type Switches.
├── 05_estructura_proyecto/     # Simulación de un esqueleto Go estándar con múltiples capas (api, internal, pkg, models, cmd).
├── 06_funciones_basicas/       # Aspectos básicos de funciones, Named Returns y control de errores por retorno múltiple.
├── 07_funciones_variadicas/    # Funciones con argumentos variádicos utilizando el operador '...'.
├── 08_funciones_tipos/         # Higher-Order unctions, composición de funciones, filtros y sort mediante closures pasados como tipos.
├── 09_closures/                # Funciones anidadas y captura de estado variable mediante Closures (generadores e incrementadores).
├── go.mod                      # Módulo 'semana_12_go' para control de packages.
└── README.md                   # Repositorio central de documentación.
```

## Requerimientos

- Instalar Go (versión sugerida: `>= 1.20`)
- Validar las variables de entorno de `GOPATH` o activar módulos (`go env -w GO111MODULE=on`).

## Ejecución

Los módulos han sido programados para no interdependerse entre subcarpetas, pero todos pertenecen al módulo general `semana_12_go`. Para ejecutar cada uno localmente, usa los siguientes comandos desde la raíz de `semana_12_go`:

```bash
# Validar, formatear y compilar la base de código completa
go fmt ./...
go vet ./...
go build ./...

# 01
go run ./01_structs

# 02
go run ./02_receptores_metodos

# 03
go run ./03_punteros

# 04
go run ./04_interfaces

# 05
go run ./05_estructura_proyecto/cmd/app

# 06
go run ./06_funciones_basicas

# 07
go run ./07_funciones_variadicas

# 08
go run ./08_funciones_tipos

# 09
go run ./09_closures
```

## Conceptos Aplicados
- **Orientación a estructuras:** En Go no existen las Clases. Se emplea composición por medio de **Structs**. 
- **Interfaces implícitas:** El Duck Typing de Go no requiere usar un *implements*. Un tipo satisface una Interfaz simplemente declarando sus métodos coincidentes.
- **Closures y Tipos Primera Clase:** Go trata a las funciones como valores transportables, devolvibles y encapsulables para conservar referencias de memoria limitadas.
