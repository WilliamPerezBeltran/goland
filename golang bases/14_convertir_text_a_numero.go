package main 

/*
   
   import "fmt"
   import "strconv"
   import "log"

*/
import (
	"fmt"
	"strconv"
	"log"
)

func main(){
	//convertir text a numero 
	value, err := strconv.Atoi("234")
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(value)
	fmt.Println("===============")

	value2, erro2 := strconv.Atoi("asdf")
	if erro2 != nil {
		log.Fatal(erro2)
	}
	fmt.Println(value2)

}
