package main 
import "fmt"

type Persona struct {
	Id int
	Name string 
}

func NewPersona(id int, name string) *Persona{
	return &Persona{
		Id: id,
		Name: name,
	}
}


func main(){
	p := NewPersona(1,"jose ignario latorre")	
	fmt.Println(p)
}

/* 


Esto:
	return &Persona{
		Id int
		Name string 
	}

Esto crea un valor de tipo Persona y devuelve su dirección de memoria.
Es decir, un puntero a la estructura 




Esto:
	func NewPersona(id int, name string) *Persona{}
Tiene como tipo de retorno *Persona lo que quire decir es que devuelve una referencia (puntero) a una estructura Persona 




*/
