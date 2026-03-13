package main

import "fmt"

func main() {
    message := "Kanpur Go dev"  // Short variable declaration
    for i := 0; i < 3; i++ {
        fmt.Printf("%s - Iteration %d\n", message, i)
    }
}
