package main
import (
	"fmt"
	"os"
)
func main (){
 //defer recoverpanic()           //   fatal error in not recovery process db close 
	dbconnection := true
	    if !dbconnection{
		  fatalError(("database not connected "))
	    }

 //normal case
 a := []string{
		"name -> A value save",
    }
  saveToDB(a)

//panic case 
	      b := []string{	
	    }
          saveToDB(b)


//normal case
 c := []string{ 
	"name -> C value save",
     }
  saveToDB(c)

}

func recoverpanic(){
	r :=recover()
	if r!=nil{
		fmt.Println("any problem create by , panic code in , massage = ",r)
	}
}

func saveToDB(record []string){
	//  defer recoverpanic()
	if len(record)<1{
		panicError("no record found to save")
		// panicError("no record found to save")
	}else{
		fmt.Println("record=",record)
	}

}

func fatalError(anything interface{}){
	fmt.Printf("fatal error accured , massage =%v",anything)
	os.Exit(1)
}

func panicError(anything interface{}){
	defer recoverpanic()
	fmt.Println("panic error accured ")
	panic(anything)
}
