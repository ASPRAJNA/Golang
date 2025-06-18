package main

import (
	"fmt"
	"math"
)

func main() {
	const inflation = 6.5
	var investmentAmount float64
	var years int
	expectedReturnRate := 5.5

	// var investmentAMount , year  =1000 , 10
	// investment , year : = 1000.0, 10.0
	fmt.Print("Enter investment amount : ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter Expected return : ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter Years : ")
	fmt.Scan(&years)
	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, float64(years))
	fmt.Println(futureValue)

	futureRealValue := futureValue / math.Pow(1+inflation/100, float64(years))
	fmt.Println(futureRealValue)
}
