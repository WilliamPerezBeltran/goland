package main 
import "fmt"

func hi(text string, c chan string){
	c <- text
}


func main (){
	c := make(chan string,1)
	fmt.Println("hello")

	go hi("william", c)
	fmt.Println(<-c)
}

