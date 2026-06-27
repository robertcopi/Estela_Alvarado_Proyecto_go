# Semana 11 - Estructuras de Control en Go

Este repositorio contiene las implementaciones en código descritas durante la sesión correspondiente a la Semana 11 del Taller de Lenguajes de Programación, que abarca el tema "Estructuras de Control en Go".

## Directorios y Módulos

La organización contiene siete sub-módulos para correr, testear y explorar individualmente:

1. **`01_if_else`**: Uso de las condicionales if, if-else, if-else-if (en cascada) y condicionales con inicialización que permiten controlar la validez de variables de manera temprana.
2. **`02_switch_case`**: Expresiones regulares de switch, fallthrough (en cascada) y categorizaciones múltiples, como por ejemplo la resolución de códigos de estado HTTP o roles.
3. **`03_for_loop`**: Diferentes presentaciones del bucle general que es usado en Go de manera única y universal: infinito, estructurado o similar a *while*. Se incluyó el "Ordenamiento de Burbuja" y validadores de límite.
4. **`04_range_loop`**: Cómo extraer las distantes keys/values, o iteraciones en arrays, slices y cadenas Unicode interactivas usando el iterador Range.
5. **`05_control_flujo`**: Muestra cómo alterar la normal ejecución interactiva del `for` infinito con tags, labels explícitos, `break`, `continue` y el clásico `goto`.
6. **`06_defer`**: Explicación en código del funcionamiento LIFO (Last In First Out) para la recolección temprana de métodos finalizadores de ejecución de recursos o archivos.
7. **`07_panic_recover`**: Manejo superior en vez de un simple error para lidiar con el fracaso inevitable pero interrumpiendo un proceso y controlando cómo falla pacíficamente o informativamente usando `recover`.

## ¿Cómo ejecutar cada módulo?

Debido a que están inicializados como repositorios disjuntos locales independientes, basta con ir en tu terminal al directorio especificado e importarlo:
```bash
# Estando en la raíz del proyecto global
cd semana_11_go/01_if_else
go run .
```

O si solo los compilas:
```bash
go build
./01_if_else     # (O en caso de tener OS Windows usas: .\01_if_else.exe)
```
