package main 
import "fmt"
// source: 
//    https://pkg.go.dev/github.com/ross-oreto/go-list-map
func main(){
	m := make(map[string]int)
	fmt.Println(m)
	m["jose ignacio latorre"] = 1959
	m["giovanni alfonso borelli"] = 1608

	fmt.Println(m)
	for i, value := range m{
		fmt.Println(i, value)
	}
	fmt.Println("====")
	fmt.Println(m["jose ignacio latorre"])
	fmt.Println(m["jsdf"])

	value, ok := m["jose ignacio latorre"]
	fmt.Println(value, ok )
	val, ok := m["asdf"]
	fmt.Println(val, ok)
}
