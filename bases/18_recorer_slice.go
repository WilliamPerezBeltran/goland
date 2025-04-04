package main 
import (
	"fmt"
	"strings"
)


/*
	el go los caracteres se almacen en bytes y no como un tipo char como en C o java 
	Cuando accede a word[i] el resuldo es un byte (uint8) y no un string 
*/
func isPalindrome(word string){
	word = strings.ToLower(word) 

	var  textReverse string 
	for i:= len(word)-1; i>=0; i--{
		textReverse += string(word[i])
	}	
	fmt.Println(textReverse)

	if word == textReverse{
		fmt.Println("Es palindrome:", word)
	}else{
		fmt.Println("No es palindrome:", word)
	}

}

func main(){
	slice := []string{"hola","que","hace"}

	for x:= range slice{
		fmt.Println(x)
	}
	isPalindrome("word")
	isPalindrome("amor a roma")
	isPalindrome("ama")
	isPalindrome("AmA")
}

