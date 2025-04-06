package main 
import "fmt"


type Employee struct{
	id int
	name string 
	address string
	apellido string
	account int
}

func (e *Employee) SetId(id int){
	e.id = id
}
func (e *Employee) SetName(name string){
	e.name = name
}

func (e *Employee) SetAddress(ad string){
	e.address = ad
}
func (e *Employee) SetApellido(ad string){
	e.apellido = ad
}
func (e *Employee) SetAccount(number int){
	e.account = number 
}
func main(){
	employee := Employee{id:2,name:"willima"} 
	fmt.Println(employee)
	employee.id = 12
	employee.name = "aasdf" 
	fmt.Println(employee)
	employee.SetId(85)
	fmt.Println("=========")
	fmt.Println(employee)
	employee.SetName("wally")
	employee.SetAddress("car 12 # 21-07")
	employee.SetApellido("perez")
	employee.SetAccount(234)
	fmt.Println(employee)
}
