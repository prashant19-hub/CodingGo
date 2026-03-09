package main

import "fmt"

func swap(x *int, y *int) {
    temp := *x      // Dereference to get value at address
    *x = *y         // Change value at x's address
    *y = temp       // Change value at y's address
}

func main() {
    a := 10
    b := 20
    fmt.Println("Before:", a, b)  // Before: 10 20
    
    swap(&a, &b)    // Pass addresses with &
    
    fmt.Println("After:", a, b)   // After: 20 10 [web:1][web:3]
}
