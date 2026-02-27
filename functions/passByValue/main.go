package main

import "fmt"

func main() {
	x := 60
	y := 12
	fmt.Println("Addition:", add(x, y))
	fmt.Println("Subtraction:", sub(x, y))
	fmt.Println("Multiplication:", mul(x, y))

	//...NORMAL CASE OF DIVISION
	div , err := divWithError(x, y)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Division:", div)
	}

    //...> FAILING CASE OF DIVISION
	// Testing division by zero
	div , err = divWithError(x, 0)
	if err == nil {
		fmt.Println("division:", div)
	} else {
		fmt.Println("Division is not defined:", err)
	}

	

	//....> PASSING CASE OF DIVISION
	// testing division by other number
	div , err = divWithError(x, 2)
	if err == nil {
		fmt.Println("division:", div)
	} else {
		fmt.Println("Division is not defined:", err)
	}

	fmt.Println("x =",x , "y =", y)

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

func div(h , i int) int {
	return h / i
}


// for division by zero, 

func divWithError(h , i int) (int , error){
	if i != 0 {
		return h / i , nil
}else{
		return 0, fmt.Errorf("division by zero")
}
}
