//Pointers

// var store address of another var

package main

import (
	"fmt"
	"strings"
)
func main(){
	//pointer is a var holds address of another variable
	i :=10
	fmt.Printf("%T %v \n", &i , &i)
	fmt.Printf("%T %v \n", *(&i), *(&i))

	//var pointername *datatype
	var ptr *int
	fmt.Println(ptr)

	//var ptr *dtype = &var
	ptr = &i
	fmt.Println(i,ptr)
	//var ptr = &var

	// & address   * deaddress 
	newptr := &i
	fmt.Println(newptr)

	fmt.Println(*newptr," ",&newptr)

	//pass by value -- param is copied to another loc original is never modified
	a := "hello"
	fmt.Println(a)
	modify(a)
	fmt.Println(a)

	//pass by refernce
	var astr *string = &a
	modifypointer(astr)
	fmt.Print(astr, *astr , a)
}
func modify(a string){
	a= strings.ToUpper(a);
}

func modifypointer(a *string){
	*a= strings.ToUpper(*a);
}