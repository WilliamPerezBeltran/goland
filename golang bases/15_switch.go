package main 
import "fmt"

func main(){
	
	x := 234
	switch x%2{
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar ")
	}
	
	fmt.Println("============================= ")
	
	switch modulo := 4 % 2; modulo {
    	case 0:
        	fmt.Println("El número es par.")
    	default:
        	fmt.Println("El número es impar.")
    }
}
