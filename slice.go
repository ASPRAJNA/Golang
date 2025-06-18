package main
import "fmt"

func main(){
	//Slice len()  cap() is refernce for array 
	
	grades := []int{10,20,30,40,50,60,70,80}
	fmt.Println(grades)
	slice1 := grades[1:8]
	fmt.Println(slice1)

	sub_slice := slice1[:2]
	fmt.Println(sub_slice,len(sub_slice),cap(sub_slice))

	slice := make([]int, 5,10)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice1[2]=7868
	fmt.Println(grades)

	newlslice := append(slice1,sub_slice...)
	fmt.Println(newlslice,len(newlslice),cap(newlslice))

	i := 2
	sl1 := grades[:i]
	sl2:= grades[i+1:]
	newsl := append(sl1,sl2...)
	fmt.Println(newsl)

	//copy
	dest_slice := make([]int,3)
	num := copy(dest_slice, grades)
	fmt.Println(dest_slice,num)

	//loop
	for index,ele:=range grades{
		fmt.Println(index, " => " , ele)
	}
}