package main 
import "fmt"


func add(dd *int){
	*dd++
}
func main(){

	arr := []int{}
	fmt.Println(arr)
	arr = append(arr,1)
	arr = append(arr,12)
	arr = append(arr,103)
	arr = append(arr,1004)
	arr = append(arr,10005)
	arr = append(arr,100006)
	var xx = []int{12,23,34,45,56}
	fmt.Println(xx)
	var myArray = []int{}
	myArray = append(myArray,xx...)
	fmt.Println(myArray)
	fmt.Println(arr)

	for index, value := range arr{
		fmt.Printf("index: %d => value: %d \n",index, value)
	}



	fmt.Printf("====================")

	nums := []int{1,2,3,4,5}
	for index, value := range nums{
		fmt.Printf("index: %d => value: %d \n",index, value)
	}


	for index, value := range nums{
		fmt.Printf("index: %d => value: %d \n",index, value)
	}


	fmt.Printf("====================")
	for x:= 0;x<len(nums);x++{
		fmt.Printf("index: %d => value: %d \n",x, nums[x])
	} 

	fmt.Printf("====================")
	hash := map[string]string{}
	hash["a"] = "value1"
	hash["b"] = "value2"
	hash["c"] = "value3"
	hash["d"] = "value4"
	hash["e"] = "value5"

	for key, value := range hash{
		fmt.Printf("key: %s => value: %s \n",key, value)
	}

	var dd int = 23

	add(&dd)
	fmt.Println(dd)










}
