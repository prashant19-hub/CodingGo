package main

import "fmt"

func main() {
	x := 60
	y := 12

	//  exmple 1: ->basic division
	// x := 60
	// y := 12
	// fmt.Println("x/y =" , x/y)

	// example 2: -> using functions for basic arithmetic operations
	
	fmt.Println("addition of x and y is " , add(x,y))
	fmt.Println("subtraction of x and y is " , sub(x,y))
	fmt.Println("multiplication of x and y is " , mul(x,y))
	fmt.Println("division of x and y is " , )
	
}

// calculator.go
//for addition 
func add(a , b int) int {
	c := a + b
	return c
}

//for subtraction
func sub(a , b int) int {
	return a - b
}

//for multiplication

func mul(f , g int) int {
	return f * g
}

//for division
//for division ....code....

// func div(e , d int) int {
// 	if d != 0 {
// 	return e / d
// 	} else {
// 		return 0 
// }
 // Return 0 or handle division by zero as needed

// for division by zero,

func divWithError(h , i int) (int , error){
	if i != 0 {
		return h / i , nil
}else{
		return 0, fmt.Errorf("division by zero")
}
}
