package main

import "fmt"
var Name string = "My name"
//this is a comment  
func main() {

    name:= `Prajna`
    name ="A S Prajna"
    var q,w ,e,r float32=1.99,2,3,4
    fmt.Print(q,w,e,r)
    age:= 24
    var s string = "I am learning go"
    var i int = 100
    var f float64 = 100.5
    fmt.Println("Hello, World!",name)
    fmt.Println("Age is",age)
    fmt.Println("Hello \n ",name, " - >" , s)
    fmt.Print(i,f)
    fmt.Printf("\n Template %v",name)
    fmt.Println("Marks : %d",i)
    fmt.Printf("Marks %d Name %v",i,name)
   
    var s2,s3 string ="Poo","rna"
    var (
        s4 string="food"
        i4 int=5
    )

    //block inner block - outerblock
    fmt.Println(s4,i4,s2,s3)
    {
        country:="India"
        fmt.Print(country,name)
    }
    //local , global ( top of program, access anywhere )
    fmt.Println(Name)

    var f1 float64
    fmt.Printf("%.2f",f1)
}
