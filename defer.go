//defer statement delays the execution of a function until the surrounding function returns

package main

import "fmt"

func main() {
    fmt.Println("Start of main")

    defer printMessage("Deferred message 1")
    defer printMessage("Deferred message 2")

    fmt.Println("End of main")
}

func printMessage(msg string) {
    fmt.Println("Executing:", msg)
}
