en golang las clases privadas y publicas no se definen como en java, pythom o ruby, tiene su propia sintaxis:
	public => Se nombran en MAYUSCULA la primera letra 
	privade => Se nombran en minuscula la primera letra
Las clases protegidas no existen en golang pero si existen los setter y los getters 

Resumen
🔹 Go solo tiene público y privado.
🔹 Todo lo que empieza con mayúscula es público.
🔹 Todo lo que empieza con minúscula es privado.
🔹 No hay protected, pero puedes usar getters/setters para control de acceso.


Resumen
🔹 Go solo tiene público y privado.
🔹 Todo lo que empieza con mayúscula es export osea publico.

ahora tenemos:
	- go mod init <define elnombre del modulo>
	go mod   => crea un archivo que se llama "go.mod" especificando la version que se tiene de golang 
	init     => inicializa el modulo
	<nombre> => Define el nombre del modulo  


si tenemos esta estructura:
.
├── go.mod
├── main.go   <======= aqui se ejecuta "go mod init main"
    
    package main 
    import "fmt"
    import "struct/mypackage"

    func main(){
	var myCar mypackage.CarPublic 
	fmt.Println(myCar)
    }

├── mypackage
│   └── mypackage.go
        
	package mypackage

        // CarPublic es una estructura pública con marca y año.
        type CarPublic struct {
         	Brand string
	        Year  int
        }

└── README.md

