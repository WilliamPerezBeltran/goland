package myLibrery

type Employee struct{
	 	Id     int
    	Name     string
    	Address  string
    	Apellido string
    	Account  int
}

func (e *Employee) SetId(id int){
	e.Id = id
}
func (e *Employee) SetName(name string){
	e.Name = name
}

func (e *Employee) SetAddress(ad string){
	e.Address = ad
}
func (e *Employee) SetApellido(ad string){
	e.Apellido = ad
}
func (e *Employee) SetAccount(number int){
	e.Account = number 
}
