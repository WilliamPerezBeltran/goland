package main 
import "fmt"

func mapWalo(nums []int,f func(int)int)[]int{
	newSlice := make([]int, len(nums))
	for i,x := range nums{
		newSlice[i] = f(x)
	}
	return newSlice

}
func main(){
	fmt.Println("======== ejemplo ========") 
	array := []int{1,2,3,4,5,6,7,8,9}
	map1 := mapWalo(array,func(x int)int{
		return x*2 
	})
	fmt.Println(map1)

	fmt.Println("======== ejemplo ========") 
	b := mapWalo(array,func(x int)int{
		return x+2 
	})
	fmt.Println(b)

	fmt.Println("======== ejemplo ========") 
	c := mapWalo(array,func(x int)int{
		return x-2 
	})
	fmt.Println(c)

}

