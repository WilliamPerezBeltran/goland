package main
import "fmt"

func main() {

  var arr_ []int 
  explore := []float64{}
  
  fmt.Println(arr_)
  fmt.Println(explore)
  var x [5]float64
  x[0] = 98
  x[1] = 93
  x[2] = 77
  x[3] = 82
  x[4] = 83
	
  var total float64 = 0 
  for _, value := range x{
	  total += value

  }
  fmt.Println("Total: ",total/float64(len(x)))

  /*

	Conceptos nuevos en goland: 
	este es un lenguaje altamente tipado por lo tanto las divisiones deben ser del mismo tipo 
	int / float64     => no es permitido 
	float64 / int     => no es permitido 
	float64 / float64 =>  permitido 
	int/ int         =>  permitido 

	Las variables que no se utilicen es necesasrio nombrar con un guion al piso (_), ejemplo:
	
	for _, value := range x {
  		total += value
	}



  */

  var arr[3] int 
  arr[0] = 1  
  arr[1] = 2  
  arr[2] = 3  
  for x:= range arr{
	  fmt.Println(x)
  }
  fmt.Println(arr)

}
