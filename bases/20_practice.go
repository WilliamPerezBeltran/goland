package main
import(
	"fmt"
)

var data string = "juan,pedro,camilo"
func readString()string{

}

func main(){
	 var arr[4]int
	//var arr[4]int
	arr[0] = 1
	arr[1] = 10
	arr[2] = 100
	arr[3] = 1000
	fmt.Println(arr)
	for x:= range arr{
		fmt.Println(x)
	}

	arr_ := []int{1,2,3}
	fmt.Println(arr_)

	var myArr[3]int
	fmt.Println(myArr)
	fmt.Println("todas las formas de declarar un array ")

	// declaracion sin inicializacion (valores por defecto)	
	var arr1 [3]int
	fmt.Println(arr1)
 
	// declaracion con inicializacion explicita 
	var arr2 [3]int = [3]int{1,2,3}
	fmt.Println(arr2)

	// declaracion con inferencia de tipo 
	arr3 := [3]int{1,2,3}
	fmt.Println(arr3)
	fmt.Println("======")

	var data string = "juan,pedro,camilo"
	createString("william")
	readString()

}
