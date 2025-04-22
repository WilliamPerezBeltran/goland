 Ejercicio

Crear un programa que maneje una agenda de contactos. La agenda debe contener la siguiente información de cada contacto:

	Nombre
	Teléfono
	Correo electrónico

Requisitos:

1) Array: Crear un array de 5 contactos precargados.
2) Slice: Crear un slice que permita agregar nuevos contactos a la agenda.
3) Map: Crear un mapa que permita buscar un contacto por su número de teléfono.
4) Estructura privada: Crear una estructura privada contacto que contenga los datos de cada contacto.
5) Estructura pública: Crear una estructura pública Agenda que contenga el array y el slice de contactos.
6) Función para modificar elemento de estructura por puntero: Crear una función que permita modificar el teléfono de un contacto en la agenda, utilizando punteros.
7) Stringer: Implementar la interfaz fmt.Stringer en la estructura contacto para que se pueda imprimir la información de un contacto de forma amigable.

Pasos a seguir:

1) Crear la estructura contacto con los campos nombre, teléfono y correo electrónico.
2) Crear la estructura Agenda que contenga el array y el slice de contactos.
3) Implementar la interfaz fmt.Stringer en la estructura contacto.
4) Crear funciones para:
	a)Agregar un nuevo contacto al slice.
	b)Buscar un contacto por teléfono en el mapa.
	c)Modificar el teléfono de un contacto en la agenda.
5) Crear el programa principal que demuestre el uso de todos los elementos solicitados.
