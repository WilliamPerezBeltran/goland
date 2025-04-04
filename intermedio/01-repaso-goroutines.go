package main 
import(
	"fmt"
	"time"
)


func add(number int, c chan<- int) {
	time.Sleep(1*time.Second )
	fmt.Printf("the number is: %d \n",number)
	c<-number
}
func ponderado(d,e,f int, cc chan<- int){
	time.Sleep(time.Second*1)

	arr := []int{}
	arr = append(arr,d)
	arr = append(arr,e)
	arr = append(arr,f)

	var sum int
	lengthArr := len(arr)
	for _,value := range arr{
		sum+=value 

	}
	fmt.Printf("d: %d, e: %d, f: %d \n",d,e,f)
	ponderado := sum/lengthArr
	fmt.Printf("sum/lengthArr => %d/%d  => %d \n",sum,lengthArr,ponderado)
	cc <- ponderado


}
func main(){
	c := make(chan int)
	go add(2,c)
	go add(2,c)
	go add(2,c)
	go add(2,c)
	d := <-c
	e := <-c 
	f := <-c 
	fmt.Printf("d: %d, e: %d, f: %d \n",d,e,f)

	cc := make(chan int)

	go ponderado(d,e,f,cc)
	respuesta := <-cc
	fmt.Printf("respuesta: %d \n",respuesta)


}
