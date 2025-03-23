package main

import "fmt"

func main (){
	fmt.Println("Write a program that prints the numbers from 1 to 100.") 
	fmt.Println("But for multiples of three print Fizz instead of the number and for the multiples of five print Buzz. ") 
	fmt.Println("For numbers which are multiples of both three and five print 'FizzBuzz'.") 
	
	x := 1
	for (x <= 100) {
		if(x%3==0 && x%5==0){
			fmt.Println("FizzBuzz")	
		}else if(x%3 == 0){
			fmt.Println("Fizz")	
		}else if(x%5 == 0){
			fmt.Println("Buzz")	
		}
		x++
	}
}
