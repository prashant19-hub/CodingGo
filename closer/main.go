package main 

import "fmt"

func add () func() int{
	i := 5                    
	return func () int {
		i++
		return i
	}
                                               // return i ,they are only for using >func add() int {} < function in
}

func main (){

	result := add()                           //closer / as a resource /as pointer-> they are change the address value and system worked new  input based
                                              // there in i value is not a dead .they are update the input value on based output value

	fmt.Println("result",result()) 
	fmt.Println("result",result())            //they are based on output value through run because output is   func add () in pointer-> (fun()int{}). by crate and change the address copy value and again run 
	fmt.Println("result",result())
	fmt.Println("result",result())
	

}