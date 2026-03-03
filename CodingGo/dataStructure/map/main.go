package main
import "fmt"

func main(){

    //int to int map
	o := make(map[int]int)
	o[1] = 10
	o[2] = 20
	o[3] = 30
	fmt.Println(o)

	p := make(map[int]int)
	p[1] = 100
	p[2] = 200
	p[3] = 300
	val1 , ok := p[2]
	fmt.Println("Value:", val1, "Ok:", ok)

	val1 , ok = p[4]
	if ok {
		fmt.Println("Value:", val1, "Ok:", ok)
	}else{
		fmt.Println("Key not found-----")
	}
	// map declaration and initialization
	m := make(map[string]int)
	// Adding key-value pairs to the map
	// save data in map
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
	// read data from map
	val , ok := m["two"]
	 
	fmt.Println("Value:", val, "Ok:", ok)

	val , ok = m["three"]
	if ok {
		fmt.Println("Value:", val, "Ok:", ok)
	}else{
		fmt.Println("Key not found")
	}
	
}