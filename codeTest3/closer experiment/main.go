package main
import "fmt"

func add() func() int{
	  q := 5 
	 return func() int {
		q+=5 ; 
		return q
	 }
}

func main(){
	result := add()
	 
	fmt.Println("reuslt=",result())
	fmt.Println("reuslt=",result())
	fmt.Println("reuslt=,",result())
	fmt.Println("reuslt=",result())
fmt.Println("reuslt=",result())
	fmt.Println("reuslt=",result())
	fmt.Println("reuslt=,",result())
	fmt.Println("reuslt=",result())
fmt.Println("reuslt=",result())
	fmt.Println("reuslt=",result())
	fmt.Println("reuslt=,",result())
	fmt.Println("reuslt=",result())

}