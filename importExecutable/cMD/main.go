package main

import (
	"fmt"

	"github.com/prashant19-hub/CodingGo/importExecutable/calculator"
	"github.com/prashant19-hub/CodingGo/importExecutable/store"
)

func main() {

	user := calculator.User{}
	user.Name = "black panther"

	fmt.Println("sum=", calculator.Add(344, 4455))
	fmt.Println("sub=", calculator.Sub(344, 4455))
	fmt.Println("multiply=", calculator.Multiply(344, 4455))
	fmt.Println("Divide=", calculator.Divide(344, 4455))
	// fmt.Println("Divide=", calculator.divide(344, 4455))  ..? this method only woek library packag in because there in func start small letter (divide)
	// if divide is a start capital letter  (Divide) then its work in main package in
	// there in main package only use library package in only Capital letter by start func/method

	// import using interface
	var dboprs store.DBOperations
	dboprs = store.Store{}
	dboprs.SaveRecord("there in print any things .....")

}
