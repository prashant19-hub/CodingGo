package main

import "fmt"

func add(a, b int) int {
    c := a + b
    return c
}

func sub(a, b int) int {
    return a - b
}

func mul(f, g int) int {
    return f * g
}

func divWithError(h, i int) (int, error) {
    if i != 0 {
        return h / i, nil
    } else {
        return 0, fmt.Errorf("division by zero")
    }
}

func main() {
	//pass by value pass by reference
    x := 60
    y := 13

	//data type
	const number2 = 34

	var number int
	number = -12
	number = 54454

	var addrsOfNumber *int
	addrsOfNumber = &number
	number3 := 5454
	fmt.Println(number3)

	//num4 := number

	var decNum float32
	decNum = 45442.3534
	var addrsOfFloat *float32
	addrsOfFloat = &decNum

	var flags bool
	flags = true
	var addrsOfFlag *bool
	addrsOfFlag = &flags

	var name string
	name = "programming concept"
	var addrsOfName *string
	addrsOfName = &name

// data type
    fmt.Printf("number value = %d\n", number)
    fmt.Printf("decimal value = %f\n", decNum)
    fmt.Printf("flags value = %v\n", flags)
    fmt.Printf("name value = %s\n", name)
    fmt.Printf("Address of number value = %v\n", *addrsOfNumber)
    fmt.Printf("Address of decimal value = %v\n", *addrsOfFloat)
    fmt.Printf("Address of flags value = %v\n", *addrsOfFlag)
    fmt.Printf("Address of name value = %v\n", *addrsOfName)
    fmt.Printf("Address of number = %p\n", addrsOfNumber)
    fmt.Printf("decimal = %p\n", addrsOfFloat)
    fmt.Printf("name = %p\n", addrsOfName)

    //data type in use the defer
    defer fmt.Printf("number value = %d\n", number)
    defer fmt.Printf("decimal value = %f\n", decNum)
    defer fmt.Printf("flags value = %v\n", flags)
    defer fmt.Printf("name value = %s\n", name)
    defer fmt.Printf("Address of number value = %v\n", *addrsOfNumber)
    defer fmt.Printf("Address of decimal value = %v\n", *addrsOfFloat)
    defer fmt.Printf("Address of flags value = %v\n", *addrsOfFlag)
    defer fmt.Printf("Address of name value = %v\n", *addrsOfName)
    defer fmt.Printf("Address of number = %p\n", addrsOfNumber)
    defer fmt.Printf("decimal = %p\n", addrsOfFloat)
    defer fmt.Printf("name = %p\n", addrsOfName)

    //pass by value pass by reference in using defer
    defer fmt.Println("defer by addition of x and y is", add(x, y))
    defer fmt.Println("defer by subtraction of x and y is", sub(x, y))
    defer fmt.Println("defer by multiplication of x and y is", mul(x, y))
    result, err := divWithError(x, y)
    if err != nil {
        defer fmt.Println("defer by division error:", err)
    } else {
        defer fmt.Println("defer by division of x and y is", result)
    }
    // pass by value in without using defer
    fmt.Println("addition of x and y is", add(x, y))
    fmt.Println("subtraction of x and y is", sub(x, y))
    fmt.Println("multiplication of x and y is", mul(x, y))
    result2, err2 := divWithError(x, y)
    if err2 != nil {
        fmt.Println("division error:", err2)
    } else {
        fmt.Println("division of x and y is", result2)
    }


	defer fmt.Println("******************->defer start-<*************")
}
