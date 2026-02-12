package main
import "fmt"
//1.
type model struct{
	Name string
	Version int
	Features string
	Company string	
}
//2.
type student struct{
	Name string
	Age int
	Grade string
	School string
	RollNo int
	Address address
}
//3.
type address struct{
	Street string
	City string
	State string
	ZipCode int
}
//4.
type employee struct{
	Name string
	Age int
	Position string
	Department string
	Salary float64

}
//5.
type product struct{
	Name string
	Category string
	Price float64
	Stock int
	SKU string
}
//6.
type book struct{
	Title string
	Author string
	Publisher string
	ISBN string
	Pages int
	Genre string
}
//7.
type car struct{
	Make string
	Model string
	Year int
	Color string
	Mileage int
	VIN string
}
//8.
type laptop struct{
	
	Brand string
	Model string
	Processor string
	RAM int
	Storage int
	Price float64
}
//9.
type smartphone struct{
    BasicInfo
	IMEI string
	Config 
}
//10.
type Config struct{
	RAM int
	Storage int
	Camera string
	Battery int
	OS string
	Proccessor string
}
//11.
type Computer struct{
	BasicInfo
	SerialNumber string
	Configration Config
}
//12.
type BasicInfo struct{
	Brand string
	Model string
}

//example of city management system
type City struct{
	Name string
	Population int
	Area float64
	Country string
	GDP float64
}



//main function
func main() {
	// Struct example 1
	myModel := model{
		Name: "Golang",
		Version: 1,
		Features: "Concurrency, Simplicity, Performance",
		Company: "Google",
		
	}
	mymodel2 := model{
		Name: "Python",
		Version: 3,
		Features: "Easy to learn, Versatile, Large community",
		Company: "Python Software Foundation",
	}
	fmt.Println("mymodel2:", mymodel2)
	fmt.Println("myModel:", myModel)
	// Struct example 2

	var student1 student
	student1.Name = "JP"	
	student1.Age = 20
	student1.Grade = "A"
	student1.School = "ABC School"
	student1.RollNo = 12345
	//example 3.
	student1.Address = address{
		Street: "123 Main Street",
		City: "New York",
		State: "NY",
		ZipCode: 10001,
	}
	fmt.Println("student1:", student1)

	student2 := student{
		Name: "Alic1e",
		Age: 22,
		Grade: "B",
		School: "XYZ University",
		RollNo: 67890,
		//exmple 3.
		Address: address{
			Street: "456 Elm Street",
			City: "Los Angeles",
			State: "CA",
			ZipCode: 90001,
		},
	}
	fmt.Println("student2:", student2)

	// Struct example 4.
	employee1 := employee{
		Name: "PRASHANT SINGH",
		Age: 22,
		Position: "Software Engineer",
		Department: "IT",
		Salary: 75000.00,
	}
	fmt.Println("employee1:", employee1)

	// Struct example 5.
	product1 := product{
		Name: "Laptop",
		Category: "Electronics",
		Price: 1200.00,
		Stock: 50,
		SKU: "LAP12345",
	}
	fmt.Println("product1:", product1)

	// Struct example 6.
	book1 := book{
		Title: "The Go Programming Language",
		Author: "Alan A. A. Donovan",
		Publisher: "Addison-Wesley Professional",
		ISBN: "978-0134190440",
		Pages: 420,
		Genre: "Programming",
	}
	fmt.Println("book1:", book1)

	// Struct example 7.
	car1 := car{
		Make: "Toyota",
		Model: "Camry",
		Year: 2020,
		Color: "Silver",
		Mileage: 15000,
		VIN: "1HGCM82633A123456",
	}
	fmt.Println("car1:", car1)

	// Struct example 8.
	laptop1 := laptop{
		Brand: "Dell",
		Model: "XPS 13",
		Processor: "Intel Core i7",
		RAM: 16,
		Storage: 512,
		Price: 1499.99,
	}
	fmt.Println("laptop1:", laptop1)

	// Struct example 9.

	smartphone1 := smartphone{

		// example 12.
		BasicInfo: BasicInfo{
			Brand: "Apple",
			Model: "iPhone 12",
		},
		IMEI: "123456789012345",

		// example 10.
		Config: Config{
			RAM: 4,
			Storage: 128,
			Camera: "12MP",
			Battery: 2815,
			OS: "iOS 14,",
			Proccessor: "A14 Bionic",
		},
	}
	fmt.Println("smartphone1:", smartphone1)

	// Struct example 11.
	computer1 := Computer{
		// example 12.basicInfo
		BasicInfo: BasicInfo{
			Brand: "HP",
			Model: "Pavilion",
		},
		SerialNumber: "SN123456789",
		// example 10.config
		Configration: Config{
			RAM: 8,
			Storage: 256,
			Camera: "None",
			Battery: 0,
			OS: "Windows 10",
			Proccessor: "Intel Core i5",
		},
	}
	fmt.Println("computer1:", computer1)

	//part 2/ex 10.
	Computer2 := Computer{
		BasicInfo: BasicInfo{
			Brand: "Dell",
			Model: "Inspiron",
		},
		SerialNumber: "SN987654321",
		Configration: Config{
			RAM: 16,
			Storage: 512,
			Camera: "None",
			Battery: 0,
			OS: "Windows 11",
			Proccessor: "Intel Core i7",
		},
	}
	fmt.Println("computer2:", Computer2)

	// Struct example city management system
	city1 := City{
		Name: "Metropolis",
		Population: 5000000,
		Area: 600.5,
		Country: "Fictionland",
		GDP: 250000000000.00,
	}
	fmt.Println("city1:", city1)

	city2 := City{
		Name: "Gotham",
		Population: 3000000,
		Area: 450.3,
		Country: "Fictionland",
		GDP: 150000000000.00,
	}
	fmt.Println("city2:", city2)

	// Arithmetic example datatypes and values
	
	num1 := 42
	num2 := 27
	sum := num1 + num2
	fmt.Println("The sum of", sum)
	subtraction := num1 - num2
	fmt.Println("The subtraction of", subtraction)

	num4 := 10
	num5 := 3
	division := num4 / num5
	fmt.Println("The division of", division)

	multiplication := num4 * num5
	fmt.Println("The multiplication of", multiplication)

	modulus := num4 % num5
	fmt.Println("The modulus of", modulus)

	increment := num4
	increment++
	fmt.Println("The increment of", increment)

	decrement := num5
	decrement--
	fmt.Println("The decrement of", decrement)


	// primitive Pointer example datatypes and values

	const number2 = 34

	var number int
	number = -12
	number = 12345
	
	var addOfNumber *int
	addOfNumber = &number
	number3 := 123567
	fmt.Println(number3)

	var decNum float32
	decNum = 123.45
	
	var addrsOfFloat *float32
	addrsOfFloat = &decNum

	var flags bool
	flags = true//false
	var addrsOfFlag *bool
	addrsOfFlag = &flags
	
	var name string
	name = "Golang"
	var addrsOfName *string
	addrsOfName = &name



	



	fmt.Printf("number of value= %d, decimal value= %f, flag value= %v, name value= %s\n", number, decNum, flags, name)
	fmt.Printf("address of number value= %v, address of decimal value= %v, address of flag value= %v, address of name value= %v\n", addOfNumber, addrsOfFloat, addrsOfFlag, addrsOfName)
    fmt.Printf("Address of number value= %p , decimal value= %p, flag value= %p, name value= %p\n", addOfNumber, addrsOfFloat, addrsOfFlag, addrsOfName)
	
}
