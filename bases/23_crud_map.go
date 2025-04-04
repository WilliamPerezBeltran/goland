package main
import(
	"fmt"

)

//personas := map[string]int
var personas = map[string]int{
	"maria":23,
	"juan":10,
	"jose":98,
}
func read(){
	fmt.Println(personas)
}
func create(name string, age int){
	personas[name] = age
}

func update(name string, newAge int){
	if _, ok := personas[name];ok {
		personas[name] = newAge
	}else{
		fmt.Println("no existe el item")
	}
}

func remove(name string){
	_, ok := personas[name]
	if ok{
		delete(personas,name)
		fmt.Println(" Item deleted") 
	}else{
		fmt.Println("El item no existe")
	}
}

//func update(name string, newAge int){
//	_, ok := personas[name] 
//	if ok {
//		personas[name] = newAge
//	}
//}
func main(){
	read()
	create("wallo",78)
	create("irene",80)
	create("jose",56)
	read()
	update("maria", 100)
	read()
	remove("wallo")
	read()
}
