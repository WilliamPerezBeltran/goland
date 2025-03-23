# Resumen: Valores Primitivos en Go

En Go, los valores primitivos son los tipos básicos de datos que optimizan el uso de memoria y mejoran la eficiencia del código.

## Números Enteros  
- `int`: Se adapta a la arquitectura del sistema (`int32` en 32 bits, `int64` en 64 bits).  
- `int8` (-128 a 127), `int16`, `int32`, `int64` (rangos aumentan exponencialmente).  
- `uint` (sin signo) permite solo valores positivos, maximizando almacenamiento (`uint8`, `uint16`, `uint32`, `uint64`).  

## Números Decimales  
- `float32`: Rango ±1.8 × 10⁻³⁸ a ±3.4 × 10³⁸.  
- `float64`: Rango ±2.23 × 10⁻³⁰⁸ a ±1.8 × 10³⁰⁸.  
Útiles para cálculos precisos en matemáticas y ciencia.  

## Otros Tipos  
- `string`: Maneja texto, delimitado por comillas dobles.  
- `bool`: Solo admite `true` o `false`, clave en estructuras condicionales.  
- `complex64` y `complex128`: Para números complejos con partes real e imaginaria (`10 + 8i`).  

Definir correctamente los tipos en Go mejora el rendimiento y fomenta buenas prácticas de programación. 🚀

