package main

import (
	"fmt"
	"os"
)

func main (){
	//*******system generated error*******

	// b := []int{
	// 	1,2,3,4,
	// }
	// fmt.Println(b)
	// fmt.Println(len(b))
	// fmt.Println(b[4])

	//*****user generated error******

	// pass := "12356"
	// fmt.Println(len(pass))
	// if len(pass)<=9{
	//  1st method of print --> fmt.Println(fmt.Errorf("password dose not match criteria,password%v",pass))	
	// 	2nd method of print	--> // err:=(fmt.Errorf("password dose not match criteria,password%v",pass))	
    //     	                    // fmt.Println(err)....2nd method
	// 	}else{
	// 	 savetodb(pass)    
	//                                                3rd method basic print-->  //   fmt.Println("passwoed match successfully")
	
	dbconnection := true
	if !dbconnection {
		fatalerror("can't connect db")
	}	
	pass := "123567696"
	fmt.Println(len(pass))
	if len(pass)<=9{
	err := fmt.Errorf("password dose not match criteria,password%v",pass)
	                                                            // fmt.Println(err) fatal error
	panic (err)
	} else {
		     savetodb(pass)
}
}
	// Fatal error (its stops execution)
func fatalerror(massage string){
	fmt.Printf("some error occure which is causing Fatal error , massage =%v\n",massage)
	os.Exit(1)
}
func savetodb(data string){
	fmt.Println("Record saved successfully ,record =",data)
}

