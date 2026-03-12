package main
import "fmt"

type UserA struct {
	Name string
	Address string 
	contactNo int     
}
type UserB struct {
	Name string
	Address string 
	contactNo int    
}

type UserOperations interface{
addUser()
}

func main(){

var jonh UserA
jonh.Name = "Jonh wick"
jonh.Address ="Kaleen nagar"
jonh.contactNo = 12345
                                                          
wick := UserB{
Name : "kentol",
Address : "Kashi",
contactNo : +987654321,
}

// wick.addUser()            
// jonh.addUser()

var UserOperations UserOperations          //interface methods
UserOperations = jonh                     //UserA (types user jonh)
UserOperations.addUser()

UserOperations = wick
UserOperations.addUser()

}


func (user UserA)addUser(){
	fmt.Println("user A methods are working , user =",user)
}                                                         
func (user UserB)addUser(){                         
	user.Name ="NEW NAME -> jonh shnow,"                       
	fmt.Println("user B methods are working , user =",user)
	
}

	                
