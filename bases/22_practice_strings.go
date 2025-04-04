package main
import(
	"fmt"
	"strings"
)
func main(){
	fmt.Println("===========")
	text := "mi texto"

	fmt.Println(text)

	text = strings.ToUpper(text)
	fmt.Println(text)
	
	text = strings.ToLower(text)
	fmt.Println(text)
	
	var myStr []string
	myStr = strings.Split(text,",")
	fmt.Println(myStr)

	myNewStr := "hola!, mundo"
	myNewStr = strings.Replace(myNewStr,"hola","RASPUTIN",-1)
	fmt.Println(myNewStr)
	fmt.Println("===========")

	var str_ string
	str_ = "hola mundo"
	str_ = strings.ReplaceAll(str_,"hola","HOLA")
	fmt.Println(str_)


	otroStr := "la vida loca de bla"
	for i,value := range otroStr{
		fmt.Printf("i:%d   =>   value:%c \n",i,value)
	}










}


























