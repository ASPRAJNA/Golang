package main

import "fmt"

func main(){
	var city,city2="Kumta","kumta"
	fmt.Println(city ==city2)
	var city3="kumta"
	fmt.Println(city3 == city2)
	fmt.Println(city != city2)
	// <   >  >=   <=  == !=   comparison op

	// Arithmetic op + - * / %   Unary ops ++ --
	fmt.Print(3+5)
	fmt.Println("hello"+"2")

	fmt.Printf("%.2f\n", float64(40 - 5))
	fmt.Println(12 * 2)
	fmt.Println(12/2)

	fmt.Println(27 % 7)

	i :=1
	i++
	fmt.Println(i)
	i--
	fmt.Println(i)
	// logical ops   && || 
	fmt.Println((10 <100) && ( 10 <200))
	fmt.Println((10 <300) && ( 10 <0))

	fmt.Println((10 <0) || ( 10 <200))
	fmt.Println((10 <0) || ( 10 >200))

	fmt.Print(!(true), !(false))

	// assignment   =   +=  -=  /=  *=  %=
	var x int
	x=i
	fmt.Println(x)
	x+=10
	fmt.Println(x)
	x-=1
	fmt.Println(x)
	x*=10
	fmt.Println(x)
	x/=10
	fmt.Println(x)
	x%=10
	fmt.Println(x)

	//bitwise on every bit  & |  ^ (both are opp)   >>   << moves n bits mentioned 
	var z, w int= 12 , 25
	bitand := z & w
	fmt.Println(bitand)

	bitor := z | w
	fmt.Println(bitor)
	
	bitxor := z ^ w
	fmt.Println(bitxor)

	z = z << 1
	fmt.Println(z)

	w = w >> 1
	fmt.Println(w)
}