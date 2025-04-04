package main 
import "fmt"

func main(){
	var myArray []int
	fmt.Println(myArray)
	fmt.Println("================ ")

	 myArray_ := []int{1,2,3}
	fmt.Println(myArray_)
	fmt.Println("================ ")
	
	fmt.Println("ARRAYS: son INMUTABLES (no se pueden cambiar )")
	fmt.Println("slices: son MUTABLES (se pueden cambiar )")
	var x[5]int
	x[4] = 100
	fmt.Println(x)
	fmt.Println("Example: ")

	var y[5] float64
	y[0] = 34
	y[1] = 23
	y[2] = 12
	fmt.Println(y,len(y),cap(y))
//	cuantos elementos tiene el array 
//	y cual es la capacidad maxima del arr

	fmt.Println("===1====") 
	slice := []int{1,2,3,4,5,6,7,8,9}
	fmt.Println(slice,len(slice),cap(slice))
	fmt.Println(slice[0]) // en la posicion 0
	fmt.Println(slice[:3]) // imprimeme hasta la posicion 3 : 0,1,2
	fmt.Println(slice[2:4]) //imprime desde la posicion 2 hasta la 4
	fmt.Println(slice[4:])
	//		en los slices : que son los array dinamicos de golang 
	//     IMPORTANTE ==> [ <inclusivo> : <exclusivo> ]
	//     IMPORTANTE ==> [ lo incluye  : no lo incluye > ]


	fmt.Println("=======") 
	slice = append(slice,100)
	slice = append(slice,200)
	fmt.Println(slice) 

	fmt.Println("=======") 
	extraSlice := []int{10,20,30,40}
	slice = append(slice,extraSlice...)
	fmt.Println(slice) 
}
