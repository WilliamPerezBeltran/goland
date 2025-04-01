package main 
import (
	"fmt"
//	"sync"
)

func say(text string){
	fmt.Println(text)
}
func main(){
	say("hello")
	go say("world")
}



/*
	corre dentro de una propia goruotine y una vez termine muere.  Y esta linea de codigo (go say("world")) no queda en el mismo hilo
	de ejecución por lo tanto cuando termine de ejecutar lo que tiene main no va salir la ejecucipón de (go say("world"))
*/
