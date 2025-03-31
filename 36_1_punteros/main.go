package main 
import (
	"libreria/libreria"
	"fmt"
)

func main(){

	myPc := libreria.Pc{Ram: 10, Disk: 20, Brand: "window"}
	fmt.Println(myPc)
	myPc.DuplicateRam()
	fmt.Println(myPc)
	myPc.DuplicateRam()
	fmt.Println(myPc)
}
