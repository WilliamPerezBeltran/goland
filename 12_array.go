package main 
import "fmt"

func main(){
	var x[5]int
	x[4] = 100
	fmt.Println(x)
	fmt.Println("Example: ")

	var y[5] float64
	y[0] = 34
	y[1] = 23
	y[2] = 12
	y[3] = 89
	y[4] = 100
	
	var total float64 = 0
	i := 0
	for i < 5{
		total +=  y[i]
		i += 1 
	}
	fmt.Printf("Suma: %.2f\n Promedio: %.2f\n len y: %d\n i:%d \n",total,total/5,len(y),i)

	i = 0 
	total  = 0
	fmt.Println("Total: ",total)
	fmt.Println("=================================")
	for u:= 0; u<len(y) ; u++{
		total += (y[u])	

	}
	fmt.Printf("Suma: %.2f\n Promedio: %.2f\n len y: %d\n i:%d \n",total,total/5,len(y),i)

	fmt.Println("=================================")
	
}
