package main

import (
	"fmt"

	"github.com/prajna/orderapp/orders"
)

func main() {
	o := orders.NewOrder("ORD-101", 1000)
	if err := o.Validate(); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Final Order Price:", o.FinalPrice())
}
