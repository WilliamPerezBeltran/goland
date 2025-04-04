package main 
import "fmt"

type Pc struct{
	ram int 
	disk int
	brand string
}

func (myPc Pc) String() string{
	return fmt.Sprintf("%d GB RAM, %d GB DISK, y es %s ",myPc.ram,myPc.disk,myPc.brand)
}
	
func main(){
	myPc := Pc{ram:10, disk:10, brand:"ibm"}
	fmt.Println(myPc)
}


