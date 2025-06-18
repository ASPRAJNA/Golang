package main

import "fmt"

// A function that adds two integers and returns the result
func add(a int, b int) int {
    return a + b
}

func ops(a int, b int)(int,int){
	sum := a+b
	diff:= a-b
	return sum,diff
}

// func ops(a int, b int)(sum int, diff int){
// 	sum = a+b
// 	diff= a-b
// 	return 
// }

//variadic function

func sumNumbers(i int,numbers ...int)int{
	sum :=0
	for _,value := range numbers{
		sum += value
	}
	return sum
}
func f()(int,int){
	return 34,56
}
func calcSquare(numbers []int) []int {
        squares := []int{}
        for _, v := range numbers {
                squares = append(squares, v*v)
        }
        return squares

}

//recursive function
func fact(n int) int{
	if n ==1{
		return 1
	}
	return n*fact(n-1)
}
//anonymous func

func main() {
    result := add(3, 5)
    fmt.Println("Sum is:", result)
	sum,diff:=ops(20,10)
	fmt.Println(sum,diff)
	fmt.Println(sumNumbers(10,20,30))
	fmt.Println(sumNumbers(10,20,30,89,67,-8))

	v,_ :=f()
	
	fmt.Println(v)

	nums := [3]int{10, 20, 15}
    fmt.Println(calcSquare(nums[:]))
	factorial := fact(5)
	fmt.Println("Factorial of 5 is ", factorial)

	anonymous2 := func (l int, b int)int{
		return l*b
	}(5,7)
	fmt.Printf("%T \n",anonymous2)
	fmt.Println(anonymous2)

		anonymous := func (l int, b int)int{
		return l*b
	}
	fmt.Printf("%T \n",anonymous)
	fmt.Println(anonymous(20,30))
	//var (my_func = func(s string) {fmt.Println("Hey there,", s)})
	/*
	 var (
                my_func = func(s string) { fmt.Println("Hey there,", s) }
        )
        my_func("Joe")
	*/
}

//high order function : either function as argumnet or returns a function 