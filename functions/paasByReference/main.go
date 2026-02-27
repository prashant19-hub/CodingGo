package main
import "fmt"


func main() {
	
	
	f := 20
	g := 10
	addbyone(&f,g)
	fmt.Println("f =",f , "g =", g)

}
//paas by reference '*'(pointer) & '&'(address of operator) & b int is pass by value
func addbyone(a *int, b int){
	// db connection breaks
	*a = *a + 1  //reconnecting to db
	b = b + 1
	

}