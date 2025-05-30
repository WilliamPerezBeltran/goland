package main 
import "fmt"

type Persona struct{
	Id int
	Name string
}

func NewPersona(id int, name string)*Persona{
	return &Persona{
		Id: id, 
		Name: name,
	}

}

func main(){
	p := NewPersona(1,"pollo2")
	fmt.Println(p)
	fmt.Println(p.Id)
	fmt.Println(p.Name)
	fmt.Println("========== apuntadores repaso ==========")
	a := 12 // a => 12
	fmt.Println(&a)
	b := &a // b => le da la direccion de memoria de a (un puntero) = 0xc000010190
	fmt.Println(b)
	c := *b // le asigna el valor de lo que tenga aqui (0xc000010190)
	/*
	lo mismo sucede con la forma de generar un contructor en los struct le retorna la direccion de memoria del 
	struct y la funcion constructura que es una funcion normal le retorna la desreferenciacion de Persona 
	*/
	fmt.Println(c)
	fmt.Println("========== otra forma de crear constructores del objecto ==========")

	x := Persona{ Id: 2 , Name: "daniel Bernoulli"}
	fmt.Println(x)
	fmt.Println("========== otra forma de crear constructores del objecto ==========")
	i := new(Persona)
	fmt.Println(*i)
	i.Id = 2
	i.Name = "bernoulli"
	fmt.Println(*i)
	fmt.Println("========== otra forma de crear constructores del objecto ==========")

}
