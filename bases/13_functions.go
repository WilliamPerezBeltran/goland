package main 
import "fmt"

func print(message string){
	fmt.Printf("he %s \n", message)
}

func tripleArgument(x int , y int , t string){
	fmt.Println(x,y,t)
}

func tripleArgument_(x,y int, t string){
	fmt.Println(x,y,t)
}

func sum(a, b int) int{
	return a+b 
}

func sum_(a int) (b,c int){
	return a,a+a 
}



func main(){
	print("ai")
  	tripleArgument(1,2,"string")
	tripleArgument_(1,2,"string")
	fmt.Println(sum(1,2))
	value1, value2 := sum_(2)
	fmt.Println(value1, value2)
	_, value3 := sum_(2)
	fmt.Println(value3)


}
