package main
import "fmt"


func main() {
	
// 	// ANONYMOUS FUNCTION..
// 	func (){

// 	fmt.Println("anonymous fuction is running & they are hided one time use function")
// 	}()
//  // veridic function..
//  fmt.Println("sum =", add(5,6))


// }
// // simple exmple
// func add(a,b int ) int{
// 	return a+b
// }

// func add(b ...int) int{
// 	sum := 0 
// 	for _ ,val := range b {
// 		sum = sum + val  // <- 1st method , 2nd method -> sum += val (this is shortcut form)

// 		//1st case-> 0 + 5 , 2nd case ->5 + 6 ,3rd case ->11 + 7
// 	}
// 	return sum // sum = 0 + 5 + 6 + 7 .. system is start sum value sum = o
// }





//coding method  .......



// ANONYMOUS FUNCTION..
	func (a int)int{

	fmt.Println("anonymous fuction is running & they are hided one time use function",a)
	return 0
	}(5)

	//VARIDIC FUNCTIONS...

 fmt.Println("sum =", add(5,6,7 ,777))
 add(222,876676,5757,565)//add() isko hum kitni baar v le skte hai aur value rakh skte hai

}

func add(b ...int) int{
	sum := 0 
	for i ,val := range b {
		fmt.Println("i =",i , "val", val) // key & value address show by range method
		sum = sum + val  // <- 1st method , 2nd method -> sum += val (this is shortcut form)

		//1st case-> 0 + 5 , 2nd case ->5 + 6 ,3rd case ->11 + 7
	}
	return sum // sum = 0 + 5 + 6 + 7 .. system is start sum value sum = o
}



