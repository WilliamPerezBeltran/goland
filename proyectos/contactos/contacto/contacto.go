package contacto

import "fmt"

var lastId int

type contacto struct {
	Id    int
	Name  string
	Phone int
	Email string
}

func NuevoContacto(name string, phone int, email string) *contacto {
	lastId++
	return &contacto{
		Id:    lastId,
		Name:  name,
		Phone: phone,
		Email: email,
	}
}

func (c *contacto) String() string {
	return fmt.Sprintf("Name: %s\nPhone: %d\nEmail: %s\n", c.Name, c.Phone, c.Email)
}

type Agenda struct {
	ContactosArray  [5]*contacto
	ContactosSlices []*contacto
}

func NuevaAgenda(array [5]*contacto, slice []*contacto) Agenda {
	return Agenda{
		ContactosArray:  array,
		ContactosSlices: slice,
	}
}


func NuevaAgenda(array [5]*contacto, slice []*contacto) Agenda {
	return Agenda{
		ContactosArray:  array,
		ContactosSlices: slice,
	}
}
