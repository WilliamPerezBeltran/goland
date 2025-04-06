package libreria

type Persona struct{
	Id int
	Name string
}

func (this *Persona) SetId(id int){
	this.Id = id
}
func (this *Persona) GetId()int{
	return this.Id
}
func (this *Persona) SetName(name string){
	this.Name = name
}
func (this *Persona) GetName()string{
	return this.Name
}
