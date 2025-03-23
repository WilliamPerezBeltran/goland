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

	fmt.Println("=============")
	fmt.Println("If we list all the natural numbers beloe $10$ that are multiples of 3 or 5, we get 3, 5, 6$ and 9. The sum of these multiples is 23. ")
	fmt.Println("Find the sum of all the multiples of 3 or 5 below 1000.")
	number := 1000
	x := 1
	acco := 0
	for (x < number){
		if(x%3 == 0 ){
			acco += x 
		}else if (x%5 ==0){
			acco += x 
		}
		x++
	}
	fmt.Println("respuesta: ", acco)
	
}
