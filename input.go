package main
import (
	"fmt"
	"reflect"
	"strconv"
)

//%d - int %s string %t bool
func main() {
	var name string
	var age int
	fmt.Print("Enter your name and age :")
	fmt.Scanf("%s %d", &name, &age)
	fmt.Println("Hey there, ", name, "Age is ",age)
	fmt.Printf("variable Age %v is of type %T \n",age,age)
	fmt.Printf("variable name %v is of type %T \n",name,name)
	fmt.Printf("Type %v \n",reflect.TypeOf(10000))

	// type casting 
	var i int = 90
	var f float64 = float64(i)
	fmt.Printf("%.2f\n",f)

	var s string = strconv.Itoa(i)
	fmt.Printf("%q",s)

	const pi =3.142
	const pi1,pi2=3.142,"3.1"
	fmt.Println(pi,pi1,pi2)

	const name1 string ="harry Potter"
	fmt.Println(name1)
	// const name no short hand const name :=   -- error 

}