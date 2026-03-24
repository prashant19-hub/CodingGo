package main

import "fmt"

func main() {
	for i := 1; i < 11; i++ {
		defer fmt.Println(i)
	}
	defer fmt.Println(" singh")
	defer fmt.Println("I am Prashant")


	for a := 2034; a>1; a-=2 {
		defer fmt.Println(a)
	}

	fmt.Println("no using Defer then package execute the normal function main rule  & fmt.  print the first     options by to  print of code , then defer used , so start last point of main function and print the (defer fmt.)code down to upper length code")
}