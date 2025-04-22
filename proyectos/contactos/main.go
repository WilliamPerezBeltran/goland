package main

import (
	"fmt"
	as "contactos/contacto"
)

func main() {
	contacto1 := as.NuevoContacto("walo1", 3002565, "walo1@gmail.com")
	contacto2 := as.NuevoContacto("walo2", 242354, "walo2@gmail.com")
	contacto3 := as.NuevoContacto("walo3", 334534534, "walo3@gmail.com")
	contacto4 := as.NuevoContacto("walo4", 300343, "walo4@gmail.com")
	contacto5 := as.NuevoContacto("walo5", 300566123, "walo5@gmail.com")

	array := [5]*as.contacto{contacto1, contacto2, contacto3, contacto4, contacto5}
	slice := []*as.contacto{contacto1, contacto2, contacto3}

	agenda := as.NuevaAgenda(array, slice)

	for _, value := range agenda.ContactosArray {
		fmt.Println(value)
	}
}


package main

import (
	"fmt"
	as "contactos/contacto"
)

func main() {
	contacto1 := as.NuevoContacto("walo1", 3002565, "walo1@gmail.com")
	contacto2 := as.NuevoContacto("walo2", 242354, "walo2@gmail.com")
	contacto3 := as.NuevoContacto("walo3", 334534534, "walo3@gmail.com")
	contacto4 := as.NuevoContacto("walo4", 300343, "walo4@gmail.com")
	contacto5 := as.NuevoContacto("walo5", 300566123, "walo5@gmail.com")

	array := [5]*as.contacto{contacto1, contacto2, contacto3, contacto4, contacto5}
	slice := []*as.contacto{contacto1, contacto2, contacto3}

	agenda := as.NuevaAgenda(array, slice)

	for _, value := range agenda.ContactosArray {
		fmt.Println(value)
	}
}
package main

import (
	"fmt"
	as "contactos/contacto"
)

func main() {
	contacto1 := as.NuevoContacto("walo1", 3002565, "walo1@gmail.com")
	contacto2 := as.NuevoContacto("walo2", 242354, "walo2@gmail.com")
	contacto3 := as.NuevoContacto("walo3", 334534534, "walo3@gmail.com")
	contacto4 := as.NuevoContacto("walo4", 300343, "walo4@gmail.com")
	contacto5 := as.NuevoContacto("walo5", 300566123, "walo5@gmail.com")

	array := [5]*as.contacto{contacto1, contacto2, contacto3, contacto4, contacto5}
	slice := []*as.contacto{contacto1, contacto2, contacto3}

	agenda := as.NuevaAgenda(array, slice)

	for _, value := range agenda.ContactosArray {
		fmt.Println(value)
	}
}


package main

import (
	"fmt"
	as "contactos/contacto"
)

func main() {
	contacto1 := as.NuevoContacto("walo1", 3002565, "walo1@gmail.com")
	contacto2 := as.NuevoContacto("walo2", 242354, "walo2@gmail.com")
	contacto3 := as.NuevoContacto("walo3", 334534534, "walo3@gmail.com")
	contacto4 := as.NuevoContacto("walo4", 300343, "walo4@gmail.com")
	contacto5 := as.NuevoContacto("walo5", 300566123, "walo5@gmail.com")

	array := [5]*as.contacto{contacto1, contacto2, contacto3, contacto4, contacto5}
	slice := []*as.contacto{contacto1, contacto2, contacto3}

	agenda := as.NuevaAgenda(array, slice)

	for _, value := range agenda.ContactosArray {
		fmt.Println(value)
	}
}
package main

import (
	"fmt"
	as "contactos/contacto"
)

func main() {
	contacto1 := as.NuevoContacto("walo1", 3002565, "walo1@gmail.com")
	contacto2 := as.NuevoContacto("walo2", 242354, "walo2@gmail.com")
	contacto3 := as.NuevoContacto("walo3", 334534534, "walo3@gmail.com")
	contacto4 := as.NuevoContacto("walo4", 300343, "walo4@gmail.com")
	contacto5 := as.NuevoContacto("walo5", 300566123, "walo5@gmail.com")

	array := [5]*as.contacto{contacto1, contacto2, contacto3, contacto4, contacto5}
	slice := []*as.contacto{contacto1, contacto2, contacto3}

	agenda := as.NuevaAgenda(array, slice)

	for _, value := range agenda.ContactosArray {
		fmt.Println(value)
	}
}


package main

import (
	"fmt"
	as "contactos/contacto"
)

func main() {
	contacto1 := as.NuevoContacto("walo1", 3002565, "walo1@gmail.com")
	contacto2 := as.NuevoContacto("walo2", 242354, "walo2@gmail.com")
	contacto3 := as.NuevoContacto("walo3", 334534534, "walo3@gmail.com")
	contacto4 := as.NuevoContacto("walo4", 300343, "walo4@gmail.com")
	contacto5 := as.NuevoContacto("walo5", 300566123, "walo5@gmail.com")

	array := [5]*as.contacto{contacto1, contacto2, contacto3, contacto4, contacto5}
	slice := []*as.contacto{contacto1, contacto2, contacto3}

	agenda := as.NuevaAgenda(array, slice)

	for _, value := range agenda.ContactosArray {
		fmt.Println(value)
	}
}
