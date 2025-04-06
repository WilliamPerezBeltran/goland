package myLibrery

type Persona struct{
		Id     int
		Name     string
		Address  string
		Apellido string
		Account  int
}

func (this *Persona) SetId(id int){
	this.Id = id
}
func (this *Persona) SetName(name string){
	this.Name = name
}

func (this *Persona) SetAddress(ad string){
	this.Address = ad
}
func (this *Persona) SetApellido(ad string){
	this.Apellido = ad
}
func (this *Persona) SetAccount(number int){
	this.Account = number 
}

