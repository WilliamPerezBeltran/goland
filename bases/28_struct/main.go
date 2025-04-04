
package main 
import (
	"fmt"
	pk "main/mypackage"
)
//el pk es un alias 
func main(){
	//var cat = mypackage.Cat{"william","perez"}   // este sirva quitandole el pk del alias 
	var cat = pk.Cat{"william","perez"}
	fmt.Println(cat)
}
