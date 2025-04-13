package main 
import "fmt"


// 1. definimos la interfaz comun 

	type Animal interface{
		Speak() string
	}

// 2. Creamos dos structs que la implementen 

	type Dog struct{}
	func (d Dog) Speak()string{
		return "Woof!"
	}

	type Cat struct{}
	func (Cat)Speak()string{
		return "miau"
	}

// 3. creamos la funcion factory 
	
	func AnimalFactory(animalType string)Animal{
		if animalType == "dog"{
			return Dog{}
		}else if animalType == "cat"{
			return Cat{}
		}
		return nil
	}

func main(){

// 4. implementamos la funcion factory (AnimalFactory) 
	animal1 := AnimalFactory("dog")
	fmt.Println(animal1.Speak())

	animal2 := AnimalFactory("cat")
	fmt.Println(animal2.Speak())

}

