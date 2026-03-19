package main

import (
	"fmt"
	"sync"
	
)

func a(sg *sync.WaitGroup){                                  //*(poiter)  they are using for main value change and address in all data check and runned &(addresss called the wg.wait and call stop the wait)
	defer sg.Done()                                         //1st method of done
	 fmt.Println("A is running... go routine in")
	// sg.Done()                                            //2nd method of done
}

func main(){
	wg := sync.WaitGroup{}                                 //  <-- 1st method of wait grp
                                                          //2nd -->    var wg sync.WaitGroup
    wg.Add(1)
    go a(&wg)                                             //& - address of wg 
    // go a(&wg)                                          if you are useing this (2nd go routine ) after then  you change the Add() in go routine  value

  	fmt.Println("go routine by run********")
    fmt.Println("go routine by 32run-----------")
	// wg.Wait()

	fmt.Println("we are after complete the go routine then run ------")
     wg.Wait()

}