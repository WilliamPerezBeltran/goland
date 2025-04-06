package main 
import "fmt"
import "modulo/myLibrery"

func libreria(){

	employee := myLibrery.Employee{Id:2,Name:"willima"} 
	fmt.Println(employee)
	employee.Id = 12
	employee.Name = "aasdf" 
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

func libreria2(){

	employee := myLibrery.Persona{Id:2,Name:"willima"} 
	fmt.Println(employee)
	employee.Id = 12
	employee.Name = "aasdf" 
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

func main(){
 libreria()
 libreria2()
}
