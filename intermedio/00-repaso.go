package main
import (
	"fmt"
	"strconv"
	"time"
)

func repaso1(){
	var x int
	x = 10
	y := 4
	fmt.Printf("x: %d \n",x)
	fmt.Printf("y: %d \n" ,y)
 	myValue,err := strconv.ParseInt("23",0,64)
 	//myValue,err := strconv.ParseInt("akdsf",0,64)
	if err != nil{
		fmt.Printf("Error: %v \n",err)
	}else{
		fmt.Printf("myValue: %d \n",myValue)
	}

	m := make(map[string]int)
	m["a"] = 23
	fmt.Printf("m['a']: %d \n",m["a"])

	s := []int{1,2,3,4,5}
	for index, value := range s {
		fmt.Printf(" %d => %d   ",index,value)
	}
	fmt.Printf("\n")
	s = append(s,50)

	for index, value := range s {
		fmt.Printf(" %d => %d   ",index,value)
	}
	fmt.Printf(" \n")
}


func doSomething(c chan<- int){
	time.Sleep(3*time.Second)
	fmt.Println("Done")
	c <- 1
}

func main(){
	repaso1()

	//repaso de goutines 
	c := make(chan int) // definicion del canal 
	go doSomething(c) // go func  (se define un goroutine)
	d := <-c
	fmt.Printf("%d \n",d)
	

}


