package main 
import (
	"fmt"
	"time"
)


func  reciver1(x int, c chan int ){
	time.Sleep(time.Second*1)
	c <- x
}
func  reciver2(y int, c chan int){
	time.Sleep(time.Second*1)
	c <- y
}
func main(){
	c1 := make(chan int)
	c2 := make(chan int)

  go reciver1(1, c1)
  go reciver2(2, c2)

//El select se utiliza para saber que canal responde primero 

  select{
  case msm1 := <- c1:
	  fmt.Println("c1",msm1)
  case msm2 := <- c2:
	  fmt.Println("c2",msm2)

  }

}
