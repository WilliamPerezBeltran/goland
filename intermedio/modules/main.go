package main 
// import "github.com/donvito/hellomod"
import (
	"github.com/donvito/hellomod"
	helloVersion2 "github.com/donvito/hellomod/v2"
)

func main(){
	//fmt.Println("hello wordk]")
	hellomod.SayHello()
	helloVersion2.SayHello("curso parametro")
}
