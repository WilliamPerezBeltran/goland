package main
import "fmt"

type Person struct{
	Id int
}
type Employee struct{
	Name string
	Age int
}


// simulacion de herencia en golang 

type Emprendedor struct{
	Person // campo anonimo 
	Employee // campo anonimo	
}
func main(){
	emp := Emprendedor{}
	emp.Id = 1
	emp.Name = "euler"
	emp.Age = 39
	fmt.Println(emp)
}
