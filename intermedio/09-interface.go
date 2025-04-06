package main

import "fmt"

type Animal interface{
	Sonido()
}

type Perro struct{}
func (p Perro) Sonido(){
	fmt.Println("guaufff")
}

type Gato struct{}
func (g Gato) Sonido(){
	fmt.Println("miau")
}


func activarSonido(instancia Animal){
	instancia.Sonido()
}
func main(){
	p := Perro{}
	g := Gato{}
	activarSonido(p)
	activarSonido(g)

}
