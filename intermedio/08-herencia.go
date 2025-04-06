package main 
import "fmt"

type Persona struct{
	Id int
	Name string
}

func (e Persona) saludar() string{	
	return "hi persona"
}

type Trabajador struct{
	Persona
	Empresa string
}


func main(){
	p := Trabajador{
		Persona: Persona{Id:1, Name:"Daniel Bernoulli"},
		Empresa: "octopus",
	}
	fmt.Println(p)
	
	fmt.Println(p.saludar())
	/*
		la composicion en golang se desarrolla con el concepto de composicion y embeder los structs uno encima de otro 
		es decir 

		type xxx struct{
			a int
			b int 
		}
		func (this xxx) doSomething(){
			return "something"
		}
		
		type aaa struct {
			xxx	=> con esto genera el concpto de herencia le incresta o agrega el comportamiento de la clase xxx a aaa
			otroAtributo string
		}

		uso:

		- se instancia y se una la funcion 
		s := aaa{xxx:xxx{a: 2, b: 3}, otroAtributo:"bla",}
	*/
}
