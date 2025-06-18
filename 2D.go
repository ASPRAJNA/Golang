package main

import "fmt"

func main() {
    // Declare a 3×4 fixed-size array of ints
    var matrix [3][4]int

    // Initialize values
    matrix[1][2] = 7
    matrix[2][3] = 11

    // Loop through rows and columns
    for i := range matrix {
        for j := range matrix[i] {
            fmt.Printf("matrix[%d][%d] = %d\n", i, j, matrix[i][j])
        }
    }

    // Alternatively, initialize inline:
    preset := [2][3]int{{1, 2, 3}, {4, 5, 6}}
    fmt.Println("preset:", preset)
}
