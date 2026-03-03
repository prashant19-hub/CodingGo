package main 

import (
	"fmt"
	"strconv"
)
func main() {
	var a int 
	a = 10

	var b int16

	b = int16(a)
	println("Value of b:",b)

	var c float64

	c = float64(a)
	println("Value of c:", c)


var x int 
	x = 12

	var d float64

	d= float64(x)
	println("Value of d:", d)

	var e int32
	e = int32(x)
	println("Value of e:", e)

	var f string
	f = fmt.Sprintf("%d", x)
	println("Value of f:", f)
     

	var y int
	y = 123
	 
	var i string
	i = fmt.Sprintf("%d", y)
	println("Value of i:", i)
	var g float64
	g = 1233455

	var h string
	h = fmt.Sprintf("%f", g)
	println("Value of h:", h)

	//...
	num := "1234523233"
	numint, err := strconv.Atoi(num)
	fmt.Println("numint=", numint, "err=", err)
	num = "thakur"
	numint, err = strconv.Atoi(num)
	fmt.Println("numint=", numint, "err=", err)

}