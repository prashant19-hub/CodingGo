package main

import "fmt"

//non primitive data types

type student struct {
	Name       string
	Class      int
	RollNumber int
	Address    address
}
type address struct {
	Lane1   string
	Lane2   string
	Post    string
	Dist    string
	Village string
}
type Phone struct{
		BasicInfo
		IMEI  string
	    Config
	}

	 type Config struct{
		RAM int
		Storage int
		Battery int
		OS string
		Processor string
	 }
	 type Laptop struct{
		BasicInfo
		SerialNumber string
		Configration Config
	}
	type BasicInfo struct{
		Brand string
		Model string
	}

func main() {
	// ->> // primitive approch

	// var name_student1 string
	// var class_student1 int
	// var rollnumber_student1 int

	// var name_student2 string
	// var class_student2 int
	// var rollnumber_student2 int.

	//-> non primitive approch

	var student1 student

	student1.Class = 11
	student1.Name = "Pransh"
	student1.RollNumber = 23
	student1.Address.Dist = "Pune,"
	student1.Address.Village = "kothrud,"

	student2 := student{
		Name:       "Ankit",
		Class:      12,
		RollNumber: 24,
		Address: address{
			Dist:    "Pune,",
			Village: "Ganjay society",
			Lane1:   "karve road,",
			Lane2: "deccan gymkhana,",},
	}

	student4 := student{
		Name:       "Kantala",
		Class:      11,
		RollNumber: 26,
		Address: address{
			Dist:    "Pune,",
			Village: "Kharadi",
			Lane1:   "Kharadi road,",
			Lane2: "Kharadi main road,",
			Post: "Kharadi post office,",
		},
	}

	student3 := student{
		Name:       "Rohan",
		Class:      10,
		RollNumber: 25,
		Address: address{
			Dist:    "Pune,",
			Village: "Shivaji nagar",
			Lane1:   "FC road,",
			Lane2: "pound road,",
			Post: "gandhi bhavan,",
		},
	
	}



	phone1 := Phone{
		BasicInfo: BasicInfo{
			Brand: "Apple",
			Model: "iPhone 12",
		},
		IMEI: "123456789012345",
		Config: Config{
			RAM:       4,
			Storage:   64,
			Battery:   2815,
			OS:        "iOS",
			Processor: "A14 Bionic",
		},
	}

	laptop1 := Laptop{
		BasicInfo: BasicInfo{
			Brand: "Dell",
			Model: "Inspiron 15",
		},
		SerialNumber: "SN1234567890",
		Configration: Config{
			RAM:       8,
			Storage:   256,
			Battery:   4000,
			OS:        "Windows 10",
			Processor: "Intel Core i5",
		},
	}
	Val := 122
	Val2 := "1221"

	var interfaceVarexmple interface{}
	interfaceVarexmple = Val
    fmt.Println("Interface Value Example with int:", interfaceVarexmple)
	interfaceVarexmple = Val2
    fmt.Println("Interface Value Example with int:", interfaceVarexmple)
	interfaceVarexmple = false
	fmt.Println("Interface Value Example with int:", interfaceVarexmple)

	fmt.Println("Hello ++++1 Team..this is PRANSH's struct", student1)
	fmt.Println("Hello ++++ Team..this is ANKIT's struct", student2)
	fmt.Println("Hello ++++ Team..this is ROHAN's struct", student3)
	fmt.Println("Hello ++++ Team..this is KANTALA's struct", student4)
	fmt.Println("Phone Details:", phone1)
	fmt.Println("Laptop Details:", laptop1)
}
