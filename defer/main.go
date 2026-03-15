package main

import "fmt"

func main() {
	for i := 1; i < 11; i++ {
		defer fmt.Println(i)
	}

	// 2nd basic methods


	defer fmt.Println(" singh")
	defer fmt.Println("I am Prashant")


   // 3rd methods (table of 2)


	for q := 20; q>1; q-=2 {
		defer fmt.Println(q)
	}

	fmt.Println("no using Defer then package execute the normal function main rule  & fmt.  print the first     options by to  print of code , then defer used , so start last point of main function and print the (defer fmt.)code down to upper length code")
}