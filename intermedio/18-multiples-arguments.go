package main 
import "fmt"


func sum(values ...int ) (ponderado, first,  last int){
	first = values[0]
	last = values[len(values)-1]
	for _, x := range values {
		ponderado += x 

	}
	fmt.Printf("%v %v %v \n", first, last, ponderado)
	return ponderado, first, last
}
func main(){
	sum(1,2,3,4,5)
	
}
