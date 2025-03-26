package main 
import "fmt"

var arr[]string

func read(){
	fmt.Println(arr)
}

func create(name string){
	arr = append(arr,name)
}

func remove(name string){
	var temp []string
	for _,value := range arr{
		if value != name{
			temp = append(temp,value)
		}
	}
	arr = temp
	fmt.Println("finished process deleting")
}

func remove_(name string){
	 var index int 
	 for i,item := range arr{
		 if name == item{
			 index = i
		 }
	 }
	 fmt.Println(index)
	 array := []int{0,1,2,3,4,5,6}
	 fmt.Println(array)
	 fmt.Println("array[:1] ==>", array[:1])
	 fmt.Println("array[1:] ==>", array[1:])
	 fmt.Println("array[1:5] ==>", array[1:5])
	 //array[:b] union array[index+1:]
	 var s []int
	 var u []int
	 s = array[:index] 
	 u = array[index+1:]
	 fmt.Println("==")
	 fmt.Println(s)
	 fmt.Println(u)
	 arr = append(arr[:index],arr[index+1:]...)

}
func main(){
	fmt.Println(arr)
	create("jose")	
	create("maria")	
	create("william")	
	create("ignocio")	

	read()	
	remove("william")	
	read()	
	remove_("maria")	
	read()	




}


