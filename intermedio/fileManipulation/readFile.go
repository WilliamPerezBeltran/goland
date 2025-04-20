package main
import (
	"fmt"
	"os"
)

func main(){
	data, err := os.ReadFile("archivo.txt")
	if err != nil {
		fmt.Println("Error en la lectura del archivo", err)
		return 
	}

	fmt.Println("Contenido :", string(data))

}
