package main

import (
	"fmt"
	"strconv"
)

func main() {
	// typecasting int variable

	var a int
	a = 10

	var b int16
	b=int16(a)
	fmt.Println("b=",b)


	// typecasting string variable(there in int to string typecasting)
	var c string
	c=fmt.Sprintf("%d",b)
	fmt.Println("c=",c)

    // typecasting string to int
	num :="1231234"
	numint, err := strconv.Atoi(num)
	fmt.Println("numint=",numint ,"err=",err)

	// typecasting string to int with error
	//change to string which is not a number to see the error 
	num ="thakur"
	numint, err = strconv.Atoi(num)
	fmt.Println("numint=",numint ,"err=",err)




}
