package main 
import "fmt"

func main(){
	a := func(){
		fmt.Println("funcion anonima")
	}
	a()

	b := func(a int, b int){
		fmt.Println(a+b)
	}
	b(10,20)

	c := func(a, b int)int{
		return a*b
	}

	fmt.Println(c(10,20))

}

