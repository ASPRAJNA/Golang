package main

import "fmt"

func main() {
    x := make([]int, 1000000) // large slice
    fmt.Println("Slice created")

    // x goes out of scope after main ends → GC frees memory automatically
}
