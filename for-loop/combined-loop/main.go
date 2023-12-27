package main

import "fmt"

func main() {
    for i, j := 1, 1; i <= 3 && j <= 2; {
        fmt.Printf("Outer loop iteration %d\n", i)
        fmt.Printf("  Inner loop iteration %d\n", j)

        j++
        if j > 2 {
            j = 1
            i++
        }
    }
}
