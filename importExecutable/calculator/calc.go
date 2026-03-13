package calculator
import "fmt"

type User struct{
	Name string
	Contno int
	id string
}

type student struct{
	Name string
	Rollno int
	class int
	marks int
}

func Add(a ,b int) int{
	fmt.Println("add func call hua")
	return a+b
}
 func Sub(a ,b int) int{
	 return a-b
 }
func Multiply(a ,b int) int{
	return a*b
}
func Divide(a ,b int) int{
	return a/b
}
func add(a ,b int) int{    //  this func is worked & used only this package in because there in func is start small letter (sub) 
                           // if the func is start letter is capital then they use every package in (main package ,library package)
	return a-b
}