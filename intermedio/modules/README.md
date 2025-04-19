 go mod init github.com/WilliamPerezBeltran/test 
 go mod init <githurepo-propio>/<nombre-del-modulo>

 #  Proyecto Go con Módulos - Ejemplo con `hellomod`

Este proyecto muestra cómo iniciar un módulo en Go, agregar una dependencia externa y ejecutar el programa. Se utiliza el módulo `github.com/donvito/hellomod` como ejemplo.

---

##  Requisitos

- Go instalado (`go version` para verificar)
- Git instalado (opcional para manejar módulos)

---

##  Pasos para crear el proyecto

```bash
# 1. Crear carpeta y entrar en ella
mkdir modules
cd modules/

# 2. Inicializar el módulo
go mod init github.com/WilliamPerezBeltran/test

# 3. Obtener dependencia
go get github.com/donvito/hellomod

# 4. Crear archivo main.go
nvim main.go

# 5. Ejecutar el código
go run main.go

# 6. (Opcional) Compilar el ejecutable
go build main.go
./main


 Estructura del Proyecto

modules/
│
├── go.mod
├── go.sum
└── main.go
Código de Ejemplo (main.go) - Versión 1
package main

import (
    "github.com/donvito/hellomod"
)

func main() {
    hellomod.SayHello()
}

Para este ejemplo, no es necesario usar v2, ya que la versión 1 no tiene cambios mayores.

Si deseas usar la versión 2:
Cambia la importación en main.go:

package main

import (
    "github.com/donvito/hellomod/v2"
)

func main() {
    hellomod.SayHello()
}



¿Dónde se almacenan las dependencias?
Las dependencias descargadas con go get o go mod tidy se guardan en la caché de módulos local de Go:


~/go/pkg/mod
También puedes inspeccionarlas con:

go mod download -json




 Comandos Útiles
Comando	Descripción
go mod init	Inicializa el módulo con un nombre (namespace)
go get <paquete>	Descarga e instala una dependencia externa
go mod tidy	Limpia y sincroniza el go.mod y go.sum
go mod download -json	Descarga y muestra detalles de las dependencias
go run main.go	Compila y ejecuta el archivo principal
go build main.go	Compila el archivo en un binario ejecutable

## 🛠️ Comandos Útiles

| Comando                 | Descripción                                           |
|------------------------|-------------------------------------------------------|
| `go mod init`          | Inicializa el módulo con un nombre (namespace)        |
| `go get <paquete>`     | Descarga e instala una dependencia externa            |
| `go mod tidy`          | Limpia y sincroniza el `go.mod` y `go.sum`            |
| `go mod download -json`| Descarga y muestra detalles de las dependencias       |
| `go run main.go`       | Compila y ejecuta el archivo principal                |
| `go build main.go`     | Compila el archivo en un binario ejecutable           |

Al ejecutar go run main.go, deberías ver:

Hello Module!
