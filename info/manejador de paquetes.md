# Resumen

## ¿Qué es un manejador de paquetes en programación?
Los lenguajes de programación, como Python o Go, a menudo incluyen un manejador de paquetes, una herramienta esencial que permite a los desarrolladores compartir y gestionar librerías de código. Por ejemplo, Python utiliza `pip` y Go utiliza `goget`. Estos manejadores facilitan la reutilización de código y el acceso a librerías desarrolladas por la comunidad.

## ¿Cómo se utiliza goget en Go?
Para utilizar el manejador de paquetes en Go, simplemente necesitas ejecutar `goget`. No es necesario instalar nada adicional, ya que Go ya lo incluye al momento de su instalación. Con `goget`, puedes descargar librerías como **GoTour**, que fue desarrollada por Google para facilitar la introducción a Go de manera online.

## ¿Cómo instalar librerías con goget?
Ejecuta el comando:
```sh
goblanc.org/x/tour
```
Para obtener feedback durante la descarga, añade la bandera `-v` para ver los detalles.
Usa la bandera `-u` para asegurarte de que la librería se descargue nuevamente, aunque ya esté instalada.

## ¿Dónde se instala el código descargado?
Una vez descargado, el binario instalable se encuentra en la carpeta `bin` de tu `GOPATH`.
El código fuente queda en la carpeta `src`, organizada de acuerdo a la estructura de directorios.

## Descubriendo la librería GoTour
**GoTour** ofrece una plataforma offline para aprender Go. Es como un mini-curso, permitiéndote:

- Activar o desactivar opciones de sintaxis.
- Ejecutar código y restablecerlo a su estado original.
- Cambiar el idioma, aunque se recomienda usar inglés por la completitud de los recursos.

Se pueden explorar distintos módulos sobre conceptos básicos de Go, como:

- Declaración de variables.
- Estructuras.
- Interfaces y concurrencia.

## ¿Dónde encontrar paquetes externos interesantes?
Existen varias fuentes, pero una recomendada es [awesome-go.com](https://awesome-go.com), que agrupa numerosos proyectos y librerías por categoría, como:

- Audio y música.
- CSS.
- Bases de datos.
- Machine learning, entre otros.

Adicionalmente, librerías como **Echo** son muy valoradas para el desarrollo web con Go. Estos recursos incluyen enlaces a sus repositorios en GitHub, proporcionando un acceso fácil al código y documentación.

## Reto práctico
Se invita a seguir explorando [awesome-go.com](https://awesome-go.com), elegir una librería de interés y crear un **"Hello World"** con ella. Esto no solo te ayudará a practicar lo aprendido, sino también a familiarizarte más con el ecosistema de Go. Comparte tus avances y descubrimientos en la sección de comentarios, ¡se aprecia mucho el entusiasmo por compartir conocimiento! Te esperamos en la próxima clase para descubrir el uso de **Go Mods** con **Echo**.


============================


# Resumen

## ¿Qué son los problemas comunes al programar en Go y cómo GoModules los soluciona?
Durante mucho tiempo, los desarrolladores de Go enfrentaron dos grandes desafíos al programar. En primer lugar, había que dividir el código dentro de la ruta del `GoPath`: el home, el nombre de usuario y la carpeta `Go`. En segundo lugar, al llevar el código a producción, especialmente cuando se requerían librerías de terceros, muchas plataformas no permitían insertar el código completo a nivel de carpeta, sino solo el binario o código de ejecución. Esto limitaba el uso de frameworks externos para la programación web en Go.

La solución llegó con la introducción de **GoModules** alrededor de 2018-2019. **GoModules** cambió las reglas del juego al simplificar la instalación de librerías de terceros y resolver los problemas iniciales de los desarrolladores de Go.

## ¿Cómo instalar el framework Echo para desarrollo web en Go?
El framework **Echo** es una excelente herramienta para quienes deseen desarrollar aplicaciones web con Go. A continuación, te explico cómo puedes comenzar a trabajar con Echo:

### Acceso al framework Echo:
- Realiza una búsqueda de **"Echo Golang"** en Google.
- Accede a la página del framework y revisa la documentación en **GitHub** o la guía de usuario disponible en **GetStarted**.

### Instalación de Echo:
Desde la terminal, utiliza el comando `go get` para instalar el paquete necesario, asegurando siempre trabajar con la versión más estable del repositorio:
```sh
go get -u github.com/labstack/echo/v4
```

## ¿Cómo dar inicio a un módulo con GoModules y crear un Hello World?
Trabajar con **GoModules** para iniciar un módulo es sencillo y viene con la ventaja de gestionar las dependencias de forma eficiente.

### Iniciar un módulo:
En la terminal, dentro de la carpeta de tu proyecto, ejecuta:
```sh
go mod init github.com/tu-usuario/nombre-del-repositorio
```
Esto genera el archivo `go.mod`, que lista las dependencias del proyecto.

### Crear un servidor básico con Echo:

1. **Instanciar Echo:**
```go
e := echo.New()
```

2. **Configurar una ruta base y definir una función para responder "Hello World":**
```go
e.GET("/", func(c echo.Context) error {
    return c.String(http.StatusOK, "Hello World")
})
```

3. **Iniciar el servidor en el puerto 1323:**
```go
e.Logger.Fatal(e.Start(":1323"))
```

4. **Ejecutar el código con `go run` y verificar que el servidor responde en `localhost:1323`.**

## ¿Cómo modificar un framework o librerías en Go e instalar GoModules?
Modificar librerías en Go no es comúnmente recomendado, pero puede ser útil para fines educativos o de debugging. Aquí los pasos básicos:

### Localizar la librería:
- Accede al `GoPath` y busca la carpeta `pkg`, donde se almacenan los módulos descargados.

### Editar archivos:
- Localiza el archivo que deseas modificar, realiza un respaldo y edita el archivo necesario.

### Comprobar verificación del módulo:
Después de modificar, utiliza `go mod verify` para asegurarte de que las modificaciones no han dañado el módulo:
```sh
go mod verify
```

Con estos pasos, ahora puedes disfrutar de la flexibilidad y eficacia que ofrecen los **GoModules** mientras aprendes y experimentas con el framework **Echo**. ¿Te has subido ya al carro de Go? Sigue explorando y experimentando para descubrir la increíble versatilidad del lenguaje Go.

