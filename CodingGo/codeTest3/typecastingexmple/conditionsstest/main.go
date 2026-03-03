package main

func main() {

	flag := true

	if flag {
		println("Flag is true")
	} else  {
		println("Flag is false")
	}
// 2nd example
num := 10

	if num == 10 {
		println("Number is 10")
	} else if num == 20 {
		println("Number is 20")
	} else {
		println("Number is neither 10 nor 20")
	}

	// 3rd example ..>>

	name := "Alice"
	if name == "Alice" {
		println("Name is Alice")
	} else {
		println("Name is not Alice")
	}

	// 4th example >> we used "&&(AND)" for the they are used to check if both the condition is true or not

	age := 25
	if age < 18 {
		println("You are a minor")
	} else if age >= 18 && age < 65 {
		println("You are an adult")
	} else {
		println("You are a senior")
	}


	// 5th example

	// ->> we used ||<->"(OR)"for the  they are used to check if the value  and condition is true or not and &&(AND) they are used to check if both the condition is true or not
	value := 15

	if value == 10 || value > 10 && value < 30 {
		println("Value is 10 or greater than 10 and less than 30")
	} else if value < 10 {
		println("Value is less than 10//", value)
	} else {
		println("Value is neither 10 nor 20 , Value ",value)
	}

	// 6th example ..> 

	limit := 20
     if(limit == 10 || limit > 10) &&(limit % 2 == 0) {
        println("Value is 10 or greater than 10 and is even")
     }else if (limit == 10 || limit > 10) && (limit % 2 != 0) {
        println("Value is 10 or greater than 10 and is odd")
     }else {
        println("Value is not 10 or greater than 10 and is odd", limit)
     }
}

