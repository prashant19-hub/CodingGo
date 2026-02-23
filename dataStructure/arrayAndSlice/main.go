package main
import "fmt"


func main() {
	// Array of integers
	var list [15] int
	list[0] = 1
	list[1] = 2
	list[2] = 3
	list[3] = 4
	list[4] = 5
	list[14] = 16
	list[12] = 15

	fmt.Println("Array:", list)

	list[3] = 10
	fmt.Println("Array after modification:", list)
  // Array of strings
	var name [12] string
	name[0] = "Hello"
	name[1] = "World"
	name[4] = "G"
	name[2] = "!"
	name[3] = "Welcome"
	name[5] = "to"
	name[6] = "Golang"
	name[11] = "Programming"
	fmt.Println("String Array:", name)

	var list2 [5] interface{}
	type arrayofintface interface {
		int | float64 | string
	}
	list2[0] = 1
	list2[1] = 2.5
	list2[2] = "Hello"
	list2[3] = 3
	list2[4] = 4.5
	fmt.Println("Array of interface", list2)
	//slice
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice:", slice)

	var students []string
	students = append(students, "thakur")
	students = append(students, "prashant")
	students = append(students, "singh")
	fmt.Println("Students Slice:", students)

	// Modifying an element in the slice

     students[1] = "David"
	fmt.Println("Modified Students Slice:", students)
	
}