# generar modulo
- Para generar un modulo es necesario iniciar con este comando:
```bash
go mod init <nombre_del_modulo>
```

- posteriormente se genera un carpeta preferiblemente con nombre "libreria"
mkdir libreria 
```bash
cd libreria 
```

- se genera un archivo con nombre igual que la carpeta que lo contiene en este caso:
```bash
touch libreria.go 
```

- ./libreria.go:

```bash
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

```

