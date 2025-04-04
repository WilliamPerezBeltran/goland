package main 
import "fmt"


func main(){



	index := 4
	 array := []int{0,1,2,3,4,5,6}
	 fmt.Println(array)
	 fmt.Println("array[:1] ==>", array[:1])
	 fmt.Println("array[1:] ==>", array[1:])
	 fmt.Println("array[1:5] ==>", array[1:5])
	 //array[:b] union array[index+1:]
	 var s []int
	 var u []int
	 s = array[:index] 
	 u = array[index+1:]
	 fmt.Println("==")
	 fmt.Println(s)
	 fmt.Println(u)
	 array = append(array[:index],array[index+1:]...)
	 fmt.Println(array)


}


