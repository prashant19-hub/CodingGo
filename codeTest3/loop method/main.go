package main
import "fmt"

func main(){
	//1st exmple
	for a :=2 ; a<=20 ; a+=2{
		fmt.Println("a",a)
	}
   //2nd exmple
	var d int
  for 	d =10; d<20 ; d++ {
		fmt.Println("2nd for loop=",d)
	}

	d = 10
	for {
		d++
		if d == 100 {
			break
		}else{
			fmt.Println("3rd exmple =",d)
		}
	}
}