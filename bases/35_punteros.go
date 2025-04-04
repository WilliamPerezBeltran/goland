package main 
import "fmt"

type pc struct{
	ram int
	disk int 
	brand string
}

	//(myPc pc) => es un reciver de 
func (myPc pc) ping(){
	fmt.Println(myPc.brand, "Pong")

}

func (myPc pc) pip(){
	fmt.Println(myPc.brand, "pip")
}


func (myPc *pc)duplicateRam(){
	myPc.ram = myPc.ram * 2
}

func main(){

	a := 50
	b := &a

	fmt.Println(b)
	fmt.Println(*b)
	*b = 100

	fmt.Println(*b)
	fmt.Println(a)

	myPc := pc{ram:16, disk:200, brand:"msi"}
	fmt.Println(myPc)
	myPc.ping()
	myPc.pip()
	fmt.Println("======")
	fmt.Println(myPc)
	myPc.duplicateRam()
	fmt.Println(myPc)
	myPc.duplicateRam()
	fmt.Println(myPc)
	myPc.duplicateRam()
}
