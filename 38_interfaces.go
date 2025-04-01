package main 
import "fmt"


	//fmt.Println("============ primer paso (se crean los structs) ============ ")
type cuadrado struct{
	base float64
}
type rectangulo struct{
	base float64
	altura float64
}


	//fmt.Println("============ segundo paso (se crean los metodos para los structs) ============ ")
func (cu cuadrado) area() float64{
	return cu.base * cu.base 
}

func (re rectangulo) area() float64{
	return re.base * re.altura
}

	//fmt.Println("============ Tercer paso se crea la interface (para utilizar la funcion area())  ============ ")
type figura2D interface{
	area() float64
}

	//fmt.Println("============ Cuarto paso se aplica inyeccion de dependencias para manipular la interface  ============ ")
func area(f figura2D){
	fmt.Println(f.area())
}

func main(){

	fmt.Println("============ no use interface ============ ")
	myCuadrado := cuadrado{base:4}
	myRectangulo := rectangulo{base:4, altura:6}
	fmt.Println(myCuadrado.area())
	fmt.Println(myRectangulo.area())
	
	fmt.Println("=== Apply interface ============ ")
	area(myCuadrado)
	area(myRectangulo)
	
	fmt.Println("=== lista de interface  ============ ")
	myInterfaces := []interface{}{"hola",23, 4.45}

	fmt.Println(myInterfaces...)
	
}
