package main

import (
	"fmt"
	"sync"
)

func main() {
	
	//channel buffering 
	ch2 := make(chan int , 5)
	go func(){
		ch2 <- 1
		ch2 <- 2
		ch2 <- 3
		ch2 <- 4
	    ch2 <- 5
		

	}()
	i :=0
	for val := range ch2{
		fmt.Println("this method is buffering method and Output-=",val)
     if i == 4{
		break
	 }else{
		i++
	 }
		                                         // simple break methodd // if val == 5{
		                                        // 	break
		                                        // }
	}


	//partial sync

	var ch chan int = make(chan int)
    
	wg :=sync.WaitGroup{}
	wg.Add(1)
	go SendNumber(ch)
	go ReciveNumber(ch , &wg)
	wg.Wait()


}
func SendNumber(ch chan int) {
	    //func SendNumber(ch <-chan int) {  ....* (ch <- chan) by only send the msg,data,number this method using security perpose

	for i := 0; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}
func ReciveNumber(ph chan int , wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range ph {
		fmt.Println(val)
	}
	
}




//basic method of channel rules


// // unc main() {  

// 	var ch chan string = make(chan string)
// 	var ch1 chan string //=make(chan string)
// 	                    //ch <- "value pass method"
// 						//ch1 <- "value pass method"
// 	go func(){           //anonymous func using and (ch <- "____") by value passed and print time  ( " ch =%v",<-ch) by value observe 
// 		ch <- "hello team"
// 		ch1 <- "value passed but not print because print method ch1 in not observe symbol and  not use func time ch1 in 'make(chan string)'"
// 	} ()					
// 	fmt.Printf("chan=%v ,chan1=%v",<-ch ,ch1 )

// }
