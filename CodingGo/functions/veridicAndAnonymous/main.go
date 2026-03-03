package main
import "fmt"


func main() {
	// ANONYMOUS FUNCTION..
	func (a int)int{

	fmt.Println("anonymous fuction is running & they are hided one time use function",a)
	return 0
	}(5)

	//VARIDIC FUNCTIONS...

 fmt.Println("sum =", add(5,6,7 ,777))
 
}

func add(b ...int) int{
	sum := 0 
	for i ,val := range b {
	fmt.Println("i =",i , "val", val) 
	sum = sum + val  

		
}
return sum 
}



