package main
import "fmt" 

func main(){
	//Array   length & capacity
	var grades [5] int 
	fmt.Println(grades)
	
	//initialization

	// var grades [3] int = [3]int{10,2,5}
	//grades :=[3]int{10,20,30}
	//grades:=[...]int{10,20}  //empty will be 0 , extra will give error

	fruits := [...]string{"apple","orange","kiwi","banana"}
	fmt.Println(fruits,len(fruits))
	fmt.Println(fruits[1],fruits[len(fruits)-1])
	fruits[0]="apples"

	for i :=0; i< len(fruits); i++{
		fmt.Println(fruits[i])
	}
	for index, ele := range fruits{
		fmt.Println(index, " => ", ele)
	}

	// 2d array
	arr := [3][2]int{{2,3},{4,6}}
	fmt.Println(arr)
}