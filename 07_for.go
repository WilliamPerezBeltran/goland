package main 
import "fmt"

func main(){

	var a int 
	var b float64 
	var c bool 
	var p string
	fmt.Println("a int por default: ",a)
	fmt.Println("b gloat64 por default: ",b)
	fmt.Println("c bool por default: ",c)
	fmt.Println("p string por default: ",p)

	for x:= 0;x<10;x++{
		fmt.Println(x)
	}
	
	u := 0
	for u<4{
		fmt.Println(u)
		u++
	}

	fmt.Println("=============")
	i := 1
	for i<=10{
		fmt.Println(i)
		i++
	}

	fmt.Println("=============")
	length := 10
	for p:= 0;p<length;p=p+1{
		fmt.Println(p)
	}


}
