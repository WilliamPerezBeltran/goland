package main 
import "fmt"

func funcion(){
	for x := 1; x<=5 ; x++{ 
		fmt.Println("cat",x)
	}
}
func continue_(){
	fmt.Println("=======")
	for x := 1; x<=5 ; x++{ 
		if x == 2{
			continue
		}
		fmt.Println("cat",x)
	}
}
func break_(){
fmt.Println("=======")
	for x := 1; x<=5 ; x++{ 
		if x == 2{
			break	
		}
		fmt.Println("cat",x)
	}
}

func main(){
	/*

	defer:
	- Es una palabra clave para decirle al programa 
	que es la ultima linea de codigo que se va a ejecutar
	antes de que todo muera.

	- Va ejecutar la ultima funcion antes de que todo muera.
	en que casos se utiliza defer:
	por ejemplo:
	Cuando se abre una conexion a una db y le decimos que cierre la base de datos (lo ulitmo que haga es cerrar la db)
 	en el caso de abrir un archivo le decimos que cuando dejemos de utilizarlo lo cierre automaticamente 
 	 ¿y para que se hace? Para garantizar que nuestro programa no consuma mas recursos de los necesario (cuando deje
	 de utilizar un recurso, cierrelo ). Garantizar una tarea en este caso en cerrar el archivo.
	*/


	defer fmt.Println("word")
	fmt.Println("cat")
	funcion()
	continue_()
	break_()
}
