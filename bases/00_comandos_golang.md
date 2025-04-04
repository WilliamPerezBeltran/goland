# 📌 Comandos básicos de Go

| Comando           | Descripción |
|------------------|-------------|
| `go version`     | Muestra la versión de Go instalada. |
| `go env`         | Muestra las variables de entorno de Go. |
| `go help`        | Muestra ayuda sobre los comandos de Go. |

---

## 📦 Gestión de módulos y dependencias (Go Modules)

| Comando                   | Descripción |
|---------------------------|-------------|
| `go mod init <nombre>`    | Crea un nuevo módulo (`go.mod`). |
| `go mod tidy`             | Limpia dependencias no utilizadas y descarga las necesarias. |
| `go mod download`         | Descarga todas las dependencias del `go.mod`. |
| `go list -m all`          | Lista todas las dependencias del proyecto. |
| `go get <paquete>`        | Descarga e instala un paquete específico. (Obsoleto para instalar binarios en Go 1.17+) |

---

## ⚙️ Compilación y ejecución

| Comando                  | Descripción |
|--------------------------|-------------|
| `go run <archivo>.go`    | Ejecuta un programa Go sin compilarlo previamente. |
| `go build`               | Compila el código fuente y genera un binario. |
| `go install`             | Compila e instala un paquete en `$GOPATH/bin` o `$HOME/go/bin`. |

---

## 🧪 Testing y benchmarking

| Comando                 | Descripción |
|-------------------------|-------------|
| `go test ./...`        | Ejecuta todas las pruebas en el proyecto. |
| `go test -v`           | Ejecuta pruebas con salida detallada. |
| `go test -bench .`     | Ejecuta pruebas de rendimiento (benchmarks). |

---

## 📊 Inspección y depuración

| Comando                       | Descripción |
|--------------------------------|-------------|
| `go fmt ./...`                 | Formatea automáticamente el código Go. |
| `go vet ./...`                 | Analiza el código en busca de posibles errores. |
| `go doc <paquete>`             | Muestra documentación de un paquete. |
| `go tool cover -html=cover.out` | Genera un informe de cobertura de pruebas en HTML. |

---

## ⚡ Manejo de paquetes y dependencias (GOPATH, sin módulos)

> 🔹 **Nota:** Solo si trabajas sin Go Modules (`GOPATH` clásico).  

| Comando                  | Descripción |
|--------------------------|-------------|
| `go get -u <paquete>`    | Descarga y actualiza un paquete en `GOPATH/src`. |
| `go list ./...`          | Lista los paquetes del proyecto. |

---

## 🔄 Go Repl y compilación cruzada

| Comando                                              | Descripción |
|-----------------------------------------------------|-------------|
| `go install golang.org/x/tools/cmd/godoc@latest`   | Instala la herramienta `godoc` para ver documentación en el navegador. |
| `go build -o <output> <archivo>.go`                | Compila un ejecutable con un nombre específico. |
| `GOOS=linux GOARCH=amd64 go build`                 | Compilación cruzada (por ejemplo, para Linux en arquitectura AMD64). |

---

## 📢 Notas Finales

- Para proyectos modernos, usa **Go Modules** (`go mod init`, `go mod tidy`).  
- `go get` **ya no se usa** para instalar binarios en versiones recientes (usa `go install`).  
- `go fmt`, `go vet` y `go test` son esenciales para mantener código limpio y funcional.  

🚀 **¡Feliz programación en Go!**
