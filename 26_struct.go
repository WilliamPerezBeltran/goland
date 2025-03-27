package main 
import "fmt"


type car struct{
	brand string
	year int
}

type circle struct{
	name string 
	apellido string
}
func main(){
	//instanciar car 
	myCar := car{brand: "ford", year:2024}
	fmt.Println(myCar)

	//otra manera de instanciarlo
	var otherCar car
	otherCar.brand = "ferrari"
	fmt.Println(otherCar)
	fmt.Println("=============")

	//instancia cirlce 
	myCircle := circle{name:"wil", apellido:"per"}

	fmt.Println(myCircle)

	var myCircle_1 circle
	 myCircle_1.name = "ramiwo" 
	 myCircle_1.apellido = "beltran" 

	fmt.Println(myCircle_1)

}
