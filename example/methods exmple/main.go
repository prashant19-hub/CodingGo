package main
import "fmt"

type UserA struct {
	Name string
	Address string 
	contactNo int     //structure1

}
type UserB struct {
	Name string
	Address string 
	contactNo int    //structure2

}

func main(){

var jonh UserA
jonh.Name = "Jonh wick"
jonh.Address ="Kaleen nagar"
jonh.contactNo = 12345
                              //sort form of wick
wick := UserB{
Name : "kentol",
Address : "Kashi",
contactNo : +987654321,
}

wick.addUser(5)
jonh.addUserDub(123456)
jonh.addUser(5)
add(1234)

}

func add(a int) {
	fmt.Println("Add A methods are working , add A =",a)
}                                                   //functions

func (user UserA)addUser(a int){
	fmt.Println("user A methods are working , user =",user)
}                                                  //methods
                                                   //(user UserA&B) is structure
func (user UserB)addUser(a int){
	fmt.Println("user B methods are working , user =",user)
}
func (u UserA)addUserDub(a int){
	fmt.Println("user A Dub methods are working , user =",u)
	// u.Name = "user A dub methods change/isme method ya value change kr skte hai "
}