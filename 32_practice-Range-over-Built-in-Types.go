package main 
import "fmt"

func main(){
	// ======== array ===========
	nums := []int{1,2,3,4,5,6}
	sum := 0
	for index,value := range nums{
		fmt.Printf("i: %d  =>  value: %d \n",index,value)
		sum += value

	}

	fmt.Printf("suma: %d  \n",sum)
	// ========  ===========
	asdf := map[string]string{"a": "apple", "b": "banana"}
	for k,v := range asdf{

		fmt.Printf("key: %s  =>  value: %s \n",k,v)
	}

}

