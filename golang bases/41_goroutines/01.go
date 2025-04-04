package main
import (
	"fmt"
	"time"
)

func he(name string){
	fmt.Println(name)
}

func main(){
	he("string")
	go he("world")

	time.Sleep(time.Second *1)

}

