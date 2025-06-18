package main

import "fmt"

func main(){
	var a1 string="happy"
	if a1 == "happy"{
		fmt.Print("I am happy")
	}else if a1 == "angry"{
		fmt.Print("I am angry")
	}else{
		fmt.Print("I am not happy")
	}

	//switch 
	var i int = 10
	switch i {
	case 10:
		fmt.Println("i is 10")
	case 100,200:
		fmt.Println("i is 100 or 200")
	default:
		fmt.Println("none matched")
	}
	//fallthrough  case a+b == 30:
	
	switch i {
	case -5:
		fmt.Println("-5")
	case 10:
		fmt.Println("i is 10")
		fallthrough
	case 100,200:
		fmt.Println("i is 100 or 200")
		fallthrough
	default:
		fmt.Println("none matched")
	}

	 var a, b = 100, 5
        switch {
        case a/b == 10:
                fmt.Println("10")
        case a/b == 20:
                fmt.Println("20")
        case a/b == 10:
                fmt.Println("30")
        default:
                fmt.Println("default")
        }
		day := "wednesday"
        switch day {
        case "monday":
                fmt.Println("monday")
        case "tuesday":
                fmt.Println("tuesday")
        case "wednesday":
                fmt.Println("wednesday")
                fallthrough
        case "thursday":
                fmt.Println("thursday")
                fallthrough
        case "friday":
                fmt.Println("friday")
        case "saturday", "sunday":
                fmt.Println("weekend")
        default:
                fmt.Println("default")
        }

		//loop with for
		for no :=1 ; no <=5 ; no++{
			fmt.Println(no*no)
		}
		var no=1
		for no<=5{
			if no == 3{
				break
			}
			fmt.Println(no*no)
			no+=1
		}
		no=1
		for no<=5{
			if no == 3{
				continue
			}
			fmt.Println(no*no)
			no+=1
		}
		//break continue 
}