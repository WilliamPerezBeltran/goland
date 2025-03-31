package main 
import "fmt"


func increase(v int){
	v++
}

func increaseConPunteros(v *int){
	*v++
}

func main(){
	var v int = 19
	increase(v)
	fmt.Println(v)
	
	fmt.Println("========")

	increaseConPunteros(&v)
	fmt.Println(v)
}
