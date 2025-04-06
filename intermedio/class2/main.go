package main

import(
	"fmt"
	"myModulo2/libreria"
)

func main(){
	p := libreria.Persona{}
	fmt.Println(p)
	p.SetId(1)
	fmt.Println(p.GetId())
	p.SetName("ramiro")
	fmt.Println(p.GetName())
	fmt.Println(p)
}	
