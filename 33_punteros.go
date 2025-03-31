package main 
import "fmt"

func main(){

	/*
	IMPORTANTE:

		& => Representqa la DIRECCION de memoria de la variable 
		* => Representqa el VALOR que esta alojado en la dirección de memoria

		son signos opuestos 

	*/
	v := 19

	var p1 *int
	var p2 = new(int)
	p3 := &v
	// %T nos permite imprimir el tipo de datos de la variable 

	fmt.Printf("p1: %T \n", p1)
	fmt.Printf("p2: %T \n", p2)
	fmt.Printf("p3: %T \n", p3)
	
	fmt.Println("p1: ", p1)
	fmt.Println("p2: ", p2)
	fmt.Println("p3: ", p3)

	var vv int = 19

	fmt.Println("La direción de memoria de v es: ", &vv)
	fmt.Printf("El tipo de variable es %T \n", &vv)

	// hacemos que el puntero p , referencia la dirección de 
	// memoria de la variable vvv.

	fmt.Println("====================")
	var vvv int = 19
	var p *int
	p = &vvv
	fmt.Printf("La variable vvv es: %d \n",vvv)
	fmt.Printf("La direción de memoria de la variable vvv es: %v \n",&vvv)
	fmt.Printf("El puntero p referencia a la dirección de memoria: %v \n",p)
	fmt.Println("p:",p)
	fmt.Println("p:",*p)
	
	fmt.Println("====================")
	var x int = 20
	fmt.Printf("x: %d \n",x)
	fmt.Printf("x: %v \n",&x)
	var puntero *int
	puntero = &x

	fmt.Printf("puntero: %v \n",puntero)
	fmt.Printf("Desreferencia a puntero para que me de el valor de la variable puntero: %d \n",*puntero)
}
