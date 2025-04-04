package main 
import(
	"fmt"
	"time"
)

func sendData(msg string, ch chan<- string){
	time.Sleep(time.Second*1)
	fmt.Println("Done")
	ch <- msg
}

func main(){
	ch := make(chan string)
	message:= "estamos happy happy bro"
	go sendData(message, ch)
	msg := <- ch
	fmt.Println("msg")
	fmt.Println(msg)
	fmt.Println("===========================")


	c := make(chan int,3) // canal con capacidad para 3 valores 
	c <- 1
	c <- 2
	c <- 3
	c <- 3
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

}

