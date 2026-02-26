package main

import "fmt"

func main() {
{
	// for loop
	for i := -20; i < 100; i += 1 {
		fmt.Println(i)
	}
}
// for loop with multiple variables
var d int

for d = 0; d < 10; d++ {
	fmt.Println("second (d) for loop iteration:", d)
}
// infinite loop
d = 10
for {
	d++
	if d == 100 {
		break
     }else {
	  fmt.Println("third for loop iteration:", d)
     }
    }

	// 3 examples of for loop
	d =11 
	condition := true
	for condition {
		if d <= 101 {
			condition = false
			break
		}else {
			fmt.Println("fourth for loop iteration:", d)
		}
	}


	//>>...RANGE LOOP...<<
	// range loop with array
	var list [10]int
	list[0] = 11
	list[1] = 21
	list[2] = 13
	list[3] = 41
	list[4] = 15

	for i := 0; i < len(list); i++ {
		fmt.Println("first range loop iteration:", list[i])
}
for i , j := range list {
	fmt.Println("at", i, "index value is", j)
}


// range loop with slice
s := []int{11, 21, 13, 41, 15}
for i, j := range s {
	fmt.Println("at", i, "index value is", j)
}



// range loop with map
m := make(map[string]int)
m["one"] = 1
m["two"] = 2
m["three"] = 3

for k, v := range m {
	fmt.Println("key:", k, "value:", v)
}
}