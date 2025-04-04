package main
import(
	"fmt"
	"strings"
)


var data string = "juan,pedro,camilo"

func createString(newItem string){
	data += ","+newItem
}

func readString(){
	fmt.Println(data)
}

func deleteString(itemDelete string){
	items := strings.Split(data,",")
	var str []string
	for _,item := range items{
		if itemDelete != item {
			str = append(str,item)
		}
	}
	data = strings.Join(str,",")
}

func updateString(oldStr string,newStr string){
	items := strings.Split(data,",")
	for i,value := range items{
		if value == oldStr{
			items[i] = newStr
			break
		}
	}
	data = strings.Join(items,",")


}
func main(){
	fmt.Println(data)
	createString("william")
	readString()
	deleteString("pedro")
	readString()
	updateString("juan","asprilla")
	readString()
}


























