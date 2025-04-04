package main 
import (
	"fmt"
	"time"
)

func producer(c chan<- int){
	for x:=0; x<=4; x++{
		time.Sleep(time.Second)
		c <- x
	}
	close(c) //importante para que el range termine
}

func main(){

	c := make(chan int)
	go producer(c)
	//	como regla general cuando se utiliza un range sobre un channel hay que cerrar el chanel por parte del producer 
	for v := range c{
		fmt.Println("Received:",v)
	}


}
