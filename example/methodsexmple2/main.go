package main

import "fmt"

type Person struct {
    name string
    age  int
}

func (p Person) Display() {  // Value: no change
    fmt.Printf("Name: %s, Age: %d\n", p.name, p.age)
}

func (p *Person) Birthday() {  // Pointer: mutates
    p.age++
}

func main() {
    p := Person{name: "Alice", age: 25}
    p.Display()       // Name: Alice, Age: 25
    p.Birthday()
    p.Display()       // Name: Alice, Age: 26
}
