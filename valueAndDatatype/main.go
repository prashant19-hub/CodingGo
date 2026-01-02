package main

import "fmt"

func main() {
	// 12  ->int
	// 12.22-> float
	// true / false -> bool
	// c -> byte /char /string
	// "hello" / coding concept -> string
	// 'a' -> rune
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

	fmt.Printf("number value = %d , decimal value = %f ,flags value = %v , name value = %s \n", number, decNum, flags, name)

	fmt.Printf("Address of number value = %v, Address of decimal value = %v ,Address of flags value = %v \n, Address of name value = %v \n", *addrsOfNumber, *addrsOfFloat, *addrsOfFlag, *addrsOfName)
	fmt.Printf("Address of number = %p, decimal = %p, flags = %p, name = %p\n", addrsOfNumber, addrsOfFloat, addrsOfFlag, addrsOfName)

}
