package main 
import "fmt"

func operar(a, b int, f func(int,int)int)int{
	return f(a,b)

}
func main(){
	a := func(){
		fmt.Println("funcion anonima")
	}
	a()
	//map(iterable,)

	suma := operar(1,2,func(x, y int)int{
		return x + y
	})
		fmt.Println("Resultado:", suma)


}

